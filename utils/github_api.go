package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/nekowawolf/airdropv2/models"
)

var useBackupTokenUntil int64 = 0

type githubRepoResponse struct {
	StargazersCount int       `json:"stargazers_count"`
	ForksCount      int       `json:"forks_count"`
	Language        string    `json:"language"`
	PushedAt        time.Time `json:"pushed_at"`
	Owner           struct {
		AvatarURL string `json:"avatar_url"`
	} `json:"owner"`
}

func FetchGithubRepoStats(owner, repoName string) (*models.GithubStats, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repoName)
	
	tokens := []string{os.Getenv("GITHUB_TOKEN"), os.Getenv("GITHUB_TOKEN2")}
	var validTokens []string
	for _, t := range tokens {
		if t != "" {
			validTokens = append(validTokens, t)
		}
	}

	currentTokenIndex := 0
	if len(validTokens) > 1 && time.Now().UnixMilli() < useBackupTokenUntil {
		currentTokenIndex = 1
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if len(validTokens) > 0 {
		req.Header.Set("Authorization", "Bearer "+validTokens[currentTokenIndex])
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 403 || resp.StatusCode == 401 {
		if len(validTokens) > 1 {
			if currentTokenIndex == 0 {
				useBackupTokenUntil = time.Now().UnixMilli() + 5*60*60*1000 // Switch for 5 hours
				fmt.Println("Github API rate limit hit on Token 1. Switching to Token 2 for 5 hours.")
			} else {
				useBackupTokenUntil = 0
				fmt.Println("Github API rate limit hit on Token 2. Reverting to Token 1.")
			}
			
			req.Header.Set("Authorization", "Bearer "+validTokens[(currentTokenIndex+1)%2])
			resp2, err2 := client.Do(req)
			if err2 != nil {
				return nil, err2
			}
			defer resp2.Body.Close()
			resp = resp2
		}
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("github api returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data githubRepoResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	stats := &models.GithubStats{
		Stars:      data.StargazersCount,
		Forks:      data.ForksCount,
		Language:   data.Language,
		ImageURL:   data.Owner.AvatarURL,
		LastUpdate: data.PushedAt,
	}

	return stats, nil
}