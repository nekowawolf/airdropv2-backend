package utils

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GetJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("Warning: JWT_SECRET is not set!")
	}
	return []byte(secret)
}

type JWTClaims struct {
	AdminID string `json:"admin_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(adminID string) (string, error) {
	claims := JWTClaims{
		AdminID: adminID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), 
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetJWTSecret())
}

func ValidateJWT(tokenString string) (string, error) {
	log.Println("Validating Token:", tokenString) 

	tokenString = strings.TrimSpace(strings.Replace(tokenString, "Bearer", "", 1))
	log.Println("Cleaned Token:", tokenString) 

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("Unexpected signing method: %v\n", token.Header["alg"])
			return nil, errors.New("unexpected signing method")
		}
		return GetJWTSecret(), nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v\n", err)
		return "", err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		log.Println("Token claims could not be cast to JWTClaims")
		return "", errors.New("invalid token claims")
	}
	if !token.Valid {
		log.Println("Token is not valid")
		return "", errors.New("invalid token")
	}

	log.Printf("Parsed token claims: AdminID: %s\n", claims.AdminID)

	return claims.AdminID, nil
}
