package portfolio

import "go.mongodb.org/mongo-driver/bson/primitive"

type Design struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title        string             `bson:"title" json:"title"`
	Description  string             `bson:"description" json:"description"`
	ImageURL     string             `bson:"image_url" json:"image_url"`
	Link         string             `bson:"link" json:"link"`
	VideoURL     string             `bson:"video_url,omitempty" json:"video_url,omitempty"`
	Category     string             `bson:"category" json:"category"`
	Tools        []string           `bson:"tools,omitempty" json:"tools,omitempty"`
	Screenshots  []string           `bson:"screenshots,omitempty" json:"screenshots,omitempty"`
	SSDesc       string             `bson:"ss_desc,omitempty" json:"ss_desc,omitempty"`
	ColorPalette VisualAsset        `bson:"color_palette,omitempty" json:"color_palette,omitempty"`
	Typography   VisualAsset        `bson:"typography,omitempty" json:"typography,omitempty"`
}

type VisualAsset struct {
	ImageURL    string `bson:"image_url" json:"image_url"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}