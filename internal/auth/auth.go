package auth

import (
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/MaxRazen/tutor/internal/config"
	"github.com/MaxRazen/tutor/internal/db"
	"github.com/MaxRazen/tutor/internal/utils"
	"github.com/MaxRazen/tutor/pkg/oauth"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID           int
	Name         string
	Email        string
	Avatar       string
	SocialID     string
	LastLoggedAt time.Time
}

var accessTokenTTL time.Duration = time.Hour * 24

var secretKey []byte

func IssueAccessToken(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": user.ID,
		"exp": time.Now().Add(accessTokenTTL).Unix(),
	})
	return token.SignedString(secretKey)
}

func ValidateToken(accessToken string) (bool, jwt.MapClaims) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		log.Printf("auth: %v", err)
		return false, nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Printf("auth: claims decoding error: %v", err)
	}

	return token.Valid, claims
}

func CreateAccessTokenCookie(accessToken string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:    "jwt",
		Value:   accessToken,
		Path:    "/",
		Expires: time.Now().Add(accessTokenTTL),
	}
}

func ExpireAccessTokenCookie() *fiber.Cookie {
	return &fiber.Cookie{
		Name:    "jwt",
		Value:   "",
		Path:    "",
		Expires: time.Now().Add(-accessTokenTTL),
	}
}

func SetSecretKey() {
	key := config.GetEnv(config.APP_KEY, "")
	b, err := base64.StdEncoding.DecodeString(key)

	if err != nil {
		log.Fatalln(err)
	}

	if len(b) == 0 {
		log.Fatalln("auth: `APP_KEY=` env variable must be set with a valid base64 string")
	}

	secretKey = b
}

func FindUser(userId int) (*User, error) {
	var id int
	var name string
	var email string
	var avatar string
	var socialId string
	var lastLoggedAt string
	sql := `SELECT id, name, email, avatar, social_id, last_logged_at FROM users WHERE id = ? LIMIT 1`
	err := db.GetConnection().First(sql, userId).Scan(&id, &name, &email, &avatar, &socialId, &lastLoggedAt)

	if err != nil {
		return nil, err
	}

	u := User{
		ID:           id,
		Name:         name,
		Email:        email,
		Avatar:       avatar,
		SocialID:     socialId,
		LastLoggedAt: utils.ParseTimestamp(lastLoggedAt),
	}
	return &u, nil
}

func FindOrCreateUser(p oauth.Profile) (*User, error) {
	sql := `SELECT COUNT(*) > 0 FROM users WHERE email = ?`

	var exists bool
	db.GetConnection().First(sql, p.Email()).Scan(&exists)

	if !exists {
		if err := createUser(p); err != nil {
			return nil, err
		}
	}

	var id int
	var name string
	var email string
	var avatar string
	var socialId string
	var lastLoggedAt string
	sql = `SELECT id, name, email, avatar, social_id, last_logged_at FROM users WHERE email = ? LIMIT 1`
	db.GetConnection().First(sql, p.Email()).Scan(&id, &name, &email, &avatar, &socialId, &lastLoggedAt)

	u := User{
		ID:           id,
		Name:         name,
		Email:        email,
		Avatar:       avatar,
		SocialID:     socialId,
		LastLoggedAt: utils.ParseTimestamp(lastLoggedAt),
	}

	return &u, nil
}

func createUser(p oauth.Profile) error {
	sql := `INSERT INTO users (email, name, social_id, avatar) VALUES (?, ?, ?, ?)`
	return db.GetConnection().Exec(sql, p.Email(), p.Name(), p.ID(), p.Avatar())
}
