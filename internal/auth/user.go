package auth

import "time"

type User struct {
	ID           int
	Name         string
	Email        string
	Avatar       string
	SocialID     string
	LastLoggedAt time.Time
}

type PublicUser struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func (u *User) Convert2PublicUser() PublicUser {
	return PublicUser{
		ID:     u.ID,
		Name:   u.Name,
		Avatar: u.Avatar,
	}
}
