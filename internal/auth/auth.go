package auth

import (
	"fmt"
	"log"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID        string
	Name      string
	FirstName string
	LastName  string
	Email     string
	AvatarURL string
	Role      string
	SocialID  string
}

type Token struct {
}

var accessTokenTTL time.Duration = time.Hour * 24

var secretKey []byte = []byte("some_secret_key")

func SignAccessToken(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": user.SocialID,
		"exp": time.Now().Add(accessTokenTTL).Unix(),
	})
	return token.SignedString(secretKey)
}

func ValidateToken(accessToken string) bool {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		log.Printf("auth: %v", err)
		return false
	}

	return token.Valid
}

func CreateAccessTokenCookie(accessToken string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:    "jwt",
		Value:   accessToken,
		Path:    "/",
		Expires: time.Now().Add(accessTokenTTL),
	}
}
