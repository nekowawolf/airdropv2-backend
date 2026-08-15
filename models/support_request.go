package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SupportRequestRequest struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	Platform       string `json:"platform"`
	TurnstileToken string `json:"turnstile_token"`
}

type SupportRequest struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	URL       string             `bson:"url,omitempty" json:"url,omitempty"`
	Platform  string             `bson:"platform,omitempty" json:"platform,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}