package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RepoSubmissionRequest struct {
	RepoURL        string `json:"repo_url"`
	Name           string `json:"name"`
	Link           string `json:"link,omitempty"`
	TurnstileToken string `json:"turnstile_token"`
}

type RepoSubmission struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	RepoURL   string             `bson:"repo_url,omitempty" json:"repo_url,omitempty"`
	AddedBy   *AddedByInfo       `bson:"added_by,omitempty" json:"added_by,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}