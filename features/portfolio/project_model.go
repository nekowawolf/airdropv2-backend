package portfolio

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	ImageURL    string             `bson:"image_url" json:"image_url"`
	Link        string             `bson:"link" json:"link"`
	GitHubURL   string             `bson:"github_url,omitempty" json:"github_url,omitempty"`
	Screenshots []string           `bson:"screenshots,omitempty" json:"screenshots,omitempty"`
	SSDesc      string             `bson:"ss_desc,omitempty" json:"ss_desc,omitempty"`
	VideoURL    string             `bson:"video_url,omitempty" json:"video_url,omitempty"`
	UseCase     DiagramItem        `bson:"use_case,omitempty" json:"use_case,omitempty"`
	Activity    DiagramItem        `bson:"activity,omitempty" json:"activity,omitempty"`
	ERD         DiagramItem        `bson:"erd,omitempty" json:"erd,omitempty"`
	Flowchart   DiagramItem        `bson:"flowchart,omitempty" json:"flowchart,omitempty"`
	Stack       []string           `bson:"stack,omitempty" json:"stack,omitempty"`
}

type DiagramItem struct {
	ImageURL    string `bson:"image_url" json:"image_url"`
	Description string `bson:"description" json:"description"`
}