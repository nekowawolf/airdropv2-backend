package portfolio

import "go.mongodb.org/mongo-driver/bson/primitive"

type Certificate struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title    string             `bson:"title" json:"title"`
	ImageURL string             `bson:"image_url" json:"image_url"`
}