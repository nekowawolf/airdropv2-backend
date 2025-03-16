package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nekowawolf/airdropv2/module"
)

func PriceHandler(c *fiber.Ctx) error {
	coins := map[string]string{
		"btc":   "https://api.alternative.me/v1/ticker/bitcoin/",
		"eth":   "https://api.alternative.me/v1/ticker/ethereum/",
		"sol":   "https://api.alternative.me/v1/ticker/solana/",
		"bnb":   "https://api.alternative.me/v1/ticker/binancecoin/",
		"matic": "https://api.alternative.me/v1/ticker/matic-network/",
		"xrp":   "https://api.alternative.me/v1/ticker/ripple/",
	}

	results := make(map[string]interface{})

	for key, url := range coins {
		data, err := module.GetPrice(url)
		if err != nil {
			log.Println("Error fetching price for", key, ":", err)
			results[key] = "Error"
		} else {
			results[key] = data
		}
	}

	return c.JSON(results)
}