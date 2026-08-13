package module

import (
	"log"
	"time"

	"github.com/nekowawolf/airdropv2/utils"
)

func SyncAllGithubRepoStats() {
	log.Println("Starting Github Repo Stats Sync...")
	snapshotTime := time.Now().UTC()
	
	repos, err := GetAllGithubRepos()
	if err != nil {
		log.Printf("Failed to get repos for sync: %v\n", err)
		return
	}

	successCount := 0
	failCount := 0

	for _, repo := range repos {
		if repo.Owner == "" || repo.RepoName == "" {
			continue
		}

		stats, err := utils.FetchGithubRepoStats(repo.Owner, repo.RepoName)
		if err != nil {
			log.Printf("Failed to fetch stats for %s/%s: %v\n", repo.Owner, repo.RepoName, err)
			failCount++
			continue
		}

		err = UpdateGithubRepoStatsByID(repo.ID, stats)
		if err != nil {
			log.Printf("Failed to update stats in DB for %s/%s: %v\n", repo.Owner, repo.RepoName, err)
			failCount++
			continue
		}

		err = InsertGithubRepoStatsHistory(repo.ID, stats.Stars, stats.Forks, snapshotTime)
		if err != nil {
			log.Printf("Failed to insert history for %s/%s: %v\n", repo.Owner, repo.RepoName, err)
		}

		successCount++
	}

	log.Printf("Github Repo Stats Sync completed. Success: %d, Failed: %d\n", successCount, failCount)

	utils.InvalidateCache("githubrepo", "githubrepo_stats")
	utils.InvalidateCachePrefix("githubrepo_history_")
}