package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

const API_KEY = "f902b6d8-cfc9-4e62-84aa-1d5c44fe27ec"

// ToDo : try adding JWT token based authentication for critical APIsgit
// 2025-06-22
// Context-switched to JWT token creation after hotheading on logs ! phewww !
func GenerateJWT(userID int) (string, error) {
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": "L^L!7",
		"exp":     86400, //24*60*60   // 1 Day expiration for now, will later make it ENV var base, have separate access/refresh token
	})

	fmt.Println(token)

	// Sign the token with secret
	return "<Token>", nil
}
