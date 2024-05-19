package google

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/MaxRazen/tutor/pkg/oauth"
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

func (p *Provider) BeginAuth(state string) string {
	return p.config.AuthCodeURL(state, p.authCodeOptions...)
}

func (p *Provider) CompleteAuth(callbackQuery map[string]string) (oauth.Token, oauth.Profile, error) {
	code, ok := callbackQuery["code"]

	if !ok {
		return nil, nil, errors.New("oauth/google: missing required parameter code")
	}

	authToken, err := p.config.Exchange(context.Background(), code)

	if err != nil {
		return nil, nil, err
	}

	if !authToken.Valid() {
		return nil, nil, errors.New("oauth/google: invalid token received from 'google'")
	}

	u, err := p.fetchUser(authToken.AccessToken)

	if err != nil {
		return nil, nil, err
	}

	token := Token{
		accessToken:  authToken.AccessToken,
		refreshToken: authToken.RefreshToken,
		expiresIn:    authToken.Expiry.Unix(),
	}

	profile := Profile{
		id:     u.ID,
		name:   u.Name,
		email:  u.Email,
		avatar: u.Picture,
	}

	return &token, &profile, nil
}

/*
	{
		"id": "100000000000000000001",
		"email": "john.doe@gmail.com",
		"verified_email": true,
		"name": "John Doe",
		"given_name": "John",
		"family_name": "Doe",
		"picture": "https://lh3.googleusercontent.com/a/...=s96-c",
		"locale": "en"
	}
*/
type googleUser struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
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
