package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Username  string             `bson:"username" json:"username"`
	Bio       string             `bson:"bio" json:"bio"`
	AvatarURL string             `bson:"avatar_url" json:"avatar_url"`
	CoverURL  string             `bson:"cover_url" json:"cover_url"`
	Links     SocialLinks        `bson:"links" json:"links"`
}

type SocialLinks struct {
	Github     string `bson:"github" json:"github"`
	Twitter    string `bson:"twitter" json:"twitter"`
	Tiktok     string `bson:"tiktok" json:"tiktok"`
	Website    string `bson:"website" json:"website"`
	Instagram  string `bson:"instagram" json:"instagram"`
}

type LinkPost struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Username    string             `bson:"username" json:"username"`
	IsVerified  bool               `bson:"is_verified" json:"is_verified"`
	Caption     string             `bson:"caption" json:"caption"`
	URL         string             `bson:"url,omitempty" json:"url,omitempty"`
	Category    string             `bson:"category" json:"category"`
	Views       int                `bson:"views" json:"views"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type ViewStats struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	PostID    primitive.ObjectID `bson:"post_id" json:"post_id"`
	SessionID string             `bson:"session_id" json:"session_id"`
	ViewedAt  time.Time          `bson:"viewed_at" json:"viewed_at"`
}