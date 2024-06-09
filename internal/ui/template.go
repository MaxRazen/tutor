package ui

import (
	"encoding/json"

	"github.com/MaxRazen/tutor/internal/auth"
	"github.com/MaxRazen/tutor/internal/config"
)

type TemplateData struct {
	BaseUrl  string
	PageData string `json:"pageData"`
}

func NewTemplateData(data any) (TemplateData, error) {
	var td TemplateData
	b, err := json.Marshal(data)

	if err != nil {
		return td, err
	}

	td.PageData = string(b)
	td.BaseUrl = config.GetRuntimeConfig().BaseUrl

	return td, nil
}

type userInfo struct {
	User        auth.PublicUser `json:"user"`
	AccessToken string          `json:"accessToken"`
	Authorized  bool            `json:"authorized"`
}

func WrapWithKey(data any, key string) (TemplateData, error) {
	m := make(map[string]any)
	m[key] = data

	return NewTemplateData(m)
}

func WrapUserInfo(u auth.PublicUser, t string) (TemplateData, error) {
	data := userInfo{
		User:        u,
		AccessToken: t,
		Authorized:  true,
	}

	return NewTemplateData(data)
}
