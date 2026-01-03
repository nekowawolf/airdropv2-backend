package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AirdropFree struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Task        string             `json:"task,omitempty" bson:"task,omitempty"`
	Link        string             `json:"link,omitempty" bson:"link,omitempty"`
	Level       string             `json:"level,omitempty" bson:"level,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	Backed      string             `json:"backed,omitempty" bson:"backed,omitempty"`
	Funds       string             `json:"funds,omitempty" bson:"funds,omitempty"`
	Supply      string             `json:"supply,omitempty" bson:"supply,omitempty"`
	Fdv         string             `json:"fdv,omitempty" bson:"fdv,omitempty"`
	MarketCap   string             `json:"market_cap,omitempty" bson:"market_cap,omitempty"`
	Vesting     string             `json:"vesting,omitempty" bson:"vesting,omitempty"`
	LinkClaim   string             `json:"link_claim,omitempty" bson:"link_claim,omitempty"`
	LinkDiscord string             `json:"link_discord,omitempty" bson:"link_discord,omitempty"`
	LinkTwitter string             `json:"link_twitter,omitempty" bson:"link_twitter,omitempty"`
	ImageURL    string             `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	LinkGuide   string             `json:"link_guide,omitempty" bson:"link_guide,omitempty"`
	Price       float64            `json:"price,omitempty" bson:"price,omitempty"`
	USDIncome   int                `json:"usd_income,omitempty" bson:"usd_income,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	EndedAt     *time.Time         `json:"ended_at,omitempty" bson:"ended_at,omitempty"`
}

type AirdropPaid struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Task        string             `json:"task,omitempty" bson:"task,omitempty"`
	Link        string             `json:"link,omitempty" bson:"link,omitempty"`
	Level       string             `json:"level,omitempty" bson:"level,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	Backed      string             `json:"backed,omitempty" bson:"backed,omitempty"`
	Funds       string             `json:"funds,omitempty" bson:"funds,omitempty"`
	Supply      string             `json:"supply,omitempty" bson:"supply,omitempty"`
	Fdv         string             `json:"fdv,omitempty" bson:"fdv,omitempty"`
	MarketCap   string             `json:"market_cap,omitempty" bson:"market_cap,omitempty"`
	Vesting     string             `json:"vesting,omitempty" bson:"vesting,omitempty"`
	LinkClaim   string             `json:"link_claim,omitempty" bson:"link_claim,omitempty"`
	LinkDiscord string             `json:"link_discord,omitempty" bson:"link_discord,omitempty"`
	LinkTwitter string             `json:"link_twitter,omitempty" bson:"link_twitter,omitempty"`
	ImageURL    string             `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	LinkGuide   string             `json:"link_guide,omitempty" bson:"link_guide,omitempty"`
	Price       float64            `json:"price,omitempty" bson:"price,omitempty"`
	USDIncome   int                `json:"usd_income,omitempty" bson:"usd_income,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	EndedAt     *time.Time         `json:"ended_at,omitempty" bson:"ended_at,omitempty"`
}