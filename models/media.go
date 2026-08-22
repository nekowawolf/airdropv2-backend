package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Media struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Filename    string             `bson:"filename,omitempty" json:"filename,omitempty"`
	URL         string             `bson:"url,omitempty" json:"url,omitempty"`
	Size        int64              `bson:"size,omitempty" json:"size,omitempty"`
	ContentType string             `bson:"content_type,omitempty" json:"content_type,omitempty"`
	MediaType   string             `bson:"media_type,omitempty" json:"media_type,omitempty"`
	R2Key       string             `bson:"r2_key,omitempty" json:"r2_key,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}