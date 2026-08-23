package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type TurnstileResponse struct {
	Success bool `json:"success"`
}

func VerifyTurnstile(token string) bool {
	secret := os.Getenv("TURNSTILE_SECRET_KEY")
	if secret == "" {
		fmt.Println("TURNSTILE_SECRET_KEY is not set in environment")
		return false
	}

	data := url.Values{}
	data.Set("secret", secret)
	data.Set("response", token)

	resp, err := http.PostForm("https://challenges.cloudflare.com/turnstile/v0/siteverify", data)
	if err != nil {
		fmt.Println("Turnstile verification request failed:", err)
		return false
	}
	defer resp.Body.Close()

	var result TurnstileResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Turnstile response decode failed:", err)
		return false
	}

	return result.Success
}