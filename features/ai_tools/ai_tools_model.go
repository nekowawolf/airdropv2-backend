package ai_tools

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AIToolsMedia struct {
	VideoURL       string   `bson:"video_url,omitempty" json:"video_url,omitempty"`
	ScreenshotURLs []string `bson:"screenshot_urls,omitempty" json:"screenshot_urls,omitempty"`
}

type AIToolsSocials struct {
	Twitter   string `bson:"twitter,omitempty" json:"twitter,omitempty"`
	Instagram string `bson:"instagram,omitempty" json:"instagram,omitempty"`
	Discord   string `bson:"discord,omitempty" json:"discord,omitempty"`
	Youtube   string `bson:"youtube,omitempty" json:"youtube,omitempty"`
}

type AITools struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	ImageURL    string             `bson:"image_url,omitempty" json:"image_url,omitempty"`
	Website     string             `bson:"website,omitempty" json:"website,omitempty"`
	Categories  []string           `bson:"categories,omitempty" json:"categories,omitempty"`
	Media       AIToolsMedia       `bson:"media,omitempty" json:"media,omitempty"`
	Socials     AIToolsSocials     `bson:"socials,omitempty" json:"socials,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}