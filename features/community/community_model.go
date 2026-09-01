package community

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CryptoCommunity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Platforms string             `bson:"platforms,omitempty" json:"platforms,omitempty"`
	Category  string             `bson:"category,omitempty" json:"category,omitempty"`
	ImageURL  string             `bson:"image_url,omitempty" json:"image_url,omitempty"`
	Link      string             `bson:"link,omitempty" json:"link,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}