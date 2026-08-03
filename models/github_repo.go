package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GithubStats struct {
	Stars    int    `bson:"stars" json:"stars"`
	Forks    int    `bson:"forks" json:"forks"`
	Language string `bson:"language" json:"language"`
	ImageURL string `bson:"image_url" json:"image_url"`
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
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	Stats       *GithubStats       `bson:"stats,omitempty" json:"stats,omitempty"`
}