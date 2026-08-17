package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GithubStats struct {
	Stars      int       `bson:"stars" json:"stars"`
	Forks      int       `bson:"forks" json:"forks"`
	Language   string    `bson:"language" json:"language"`
	ImageURL   string    `bson:"image_url" json:"image_url"`
	LastUpdate time.Time `bson:"last_update" json:"last_update"`
}

type AddedByInfo struct {
	Name string `bson:"name,omitempty" json:"name,omitempty"`
	URL  string `bson:"url,omitempty" json:"url,omitempty"`
}

type GithubRepo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Category    string             `bson:"category,omitempty" json:"category,omitempty"`
	RepoURL     string             `bson:"repo_url,omitempty" json:"repo_url,omitempty"`
	Owner       string             `bson:"owner,omitempty" json:"owner,omitempty"`
	RepoName    string             `bson:"repo_name,omitempty" json:"repo_name,omitempty"`
	Website     string             `bson:"website,omitempty" json:"website,omitempty"`
	Twitter     string             `bson:"twitter,omitempty" json:"twitter,omitempty"`
	Instagram   string             `bson:"instagram,omitempty" json:"instagram,omitempty"`
	Discord     string             `bson:"discord,omitempty" json:"discord,omitempty"`
	Stats       *GithubStats       `bson:"stats,omitempty" json:"stats,omitempty"`
	AddedBy     *AddedByInfo       `bson:"added_by,omitempty" json:"added_by,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

type GithubRepoStatsHistory struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	RepoID    primitive.ObjectID `bson:"repo_id" json:"repo_id"`
	Stars     int                `bson:"stars" json:"stars"`
	Forks     int                `bson:"forks" json:"forks"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
}

type MdFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type GithubRepoDetailsData struct {
	RepoData map[string]interface{} `json:"repoData"`
	MdFiles  []MdFile               `json:"mdFiles"`
}