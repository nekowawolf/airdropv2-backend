package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Notes struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`        
	Title   string             `bson:"title,omitempty" json:"title,omitempty"`    
	Content string             `bson:"content,omitempty" json:"content,omitempty"` 
	Date    time.Time   	   `bson:"date,omitempty" json:"date,omitempty"` 
}
