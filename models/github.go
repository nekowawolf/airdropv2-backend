package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

type Image struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Filename string             `bson:"filename" json:"filename"`
	URL      string             `bson:"url" json:"url"`
	Size     int64              `bson:"size" json:"size"`
	Sha      string             `bson:"sha" json:"sha"`
	Path     string             `bson:"path" json:"path"`
}

type GitHubUploadRequest struct {
	Message string `json:"message"`
	Content string `json:"content"`
}

type GitHubUploadResponse struct {
	Content struct {
		Path string `json:"path"`
		Sha  string `json:"sha"`
	} `json:"content"`
}