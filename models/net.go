package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type NetMedia struct {
	ScreenshotURLs []string `bson:"screenshot_urls,omitempty" json:"screenshot_urls,omitempty"`
	VideoURL       string   `bson:"video_url,omitempty" json:"video_url,omitempty"`
}

type NetSocials struct {
	Twitter   string `bson:"twitter,omitempty" json:"twitter,omitempty"`
	Instagram string `bson:"instagram,omitempty" json:"instagram,omitempty"`
	Discord   string `bson:"discord,omitempty" json:"discord,omitempty"`
	Youtube   string `bson:"youtube,omitempty" json:"youtube,omitempty"`
}

type Net struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	ImageURL    string             `bson:"image_url,omitempty" json:"image_url,omitempty"`
	Website     string             `bson:"website,omitempty" json:"website,omitempty"`
	Categories  []string           `bson:"categories,omitempty" json:"categories,omitempty"`
	Media       NetMedia           `bson:"media,omitempty" json:"media,omitempty"`
	Socials     NetSocials         `bson:"socials,omitempty" json:"socials,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}