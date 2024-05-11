package google

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/MaxRazen/tutor/internal/auth"
	"golang.org/x/oauth2"
)

const endpointProfile string = "https://www.googleapis.com/oauth2/v2/userinfo"

type Provider struct {
	clientKey       string
	secret          string
	callbackURL     string
	httpClient      *http.Client
	config          *oauth2.Config
	authCodeOptions []oauth2.AuthCodeOption
	providerName    string
}

func (p *Provider) Name() string {
	return p.providerName
}

func (p *Provider) Client() *http.Client {
	return p.httpClient
}

// TODO: beter to call it GetAuthUrl
func (p *Provider) BeginAuth(state string) string {
	return p.config.AuthCodeURL(state, p.authCodeOptions...)
}

func (p *Provider) CompleteAuth(callbackQuery map[string]string) (*auth.User, error) {
	code, ok := callbackQuery["code"]

	if !ok {
		return nil, errors.New("oauth/google: missing required parameter code")
	}

	token, err := p.config.Exchange(context.Background(), code)

	if err != nil {
		return nil, err
	}

	if !token.Valid() {
		return nil, errors.New("oauth/google: invalid token received from 'google'")
	}

	u, err := p.fetchUser(token.AccessToken)

	if err != nil {
		return nil, err
	}

	return &auth.User{
		Name:     u.Name,
		Email:    u.Email,
		SocialID: u.ID,
	}, nil
}

func (p *Provider) fetchUser(accessToken string) (*googleUser, error) {
	endpoint := endpointProfile + "?access_token=" + url.QueryEscape(accessToken)

	response, err := p.Client().Get(endpoint)

	if err != nil {
		return nil, fmt.Errorf("oauth/google: error on fetching user data: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("oauth/google: error on fetching user data: %v", err)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("oauth/google: body could not be read: %v", err)
	}

	log.Println("oauth/google: raw user data", string(body))

	var u googleUser
	if err = json.Unmarshal(body, &u); err != nil {
		return nil, fmt.Errorf("oauth/google: body could not be deserialized: %v", err)
	}

	return &u, nil
}

func New(clientKey, secret, callbackURL string, scopes []string) *Provider {
	p := &Provider{
		clientKey:    clientKey,
		secret:       secret,
		callbackURL:  callbackURL,
		providerName: "google",
		httpClient:   http.DefaultClient,
		authCodeOptions: []oauth2.AuthCodeOption{
			oauth2.AccessTypeOffline,
		},
	}
	p.config = newConfig(p, scopes)
	return p
}

func newConfig(provider *Provider, scopes []string) *oauth2.Config {
	c := &oauth2.Config{
		ClientID:     provider.clientKey,
		ClientSecret: provider.secret,
		RedirectURL:  provider.callbackURL,
		Endpoint:     Endpoint,
		Scopes:       []string{},
	}

	if len(scopes) > 0 {
		c.Scopes = append(c.Scopes, scopes...)
	} else {
		c.Scopes = []string{"email", "profile"}
	}
	return c
}

type googleUser struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Link      string `json:"link"`
	Picture   string `json:"picture"`
}
