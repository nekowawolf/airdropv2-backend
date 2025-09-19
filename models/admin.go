package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Admin struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
  	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

type RefreshToken struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Token     string             `bson:"token" json:"token"`
	AdminID   string             `bson:"admin_id" json:"admin_id"`
	ExpiresAt time.Time          `bson:"expires_at" json:"expires_at"`
}