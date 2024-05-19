package oauth

import (
	"fmt"
)

type ProviderConfig struct {
	ClientKey   string
	Secret      string
	CallbackURL string
}

type Provider interface {
	Name() string
	BeginAuth(state string) string
	CompleteAuth(map[string]string) (Token, Profile, error)
}

type Token interface {
	AccessToken() string
	RefreshToken() string
	ExpiresIn() int64
}

type Profile interface {
	ID() string
	Name() string
	Email() string
	Avatar() string
}

type Providers map[string]Provider

var providers = Providers{}

func UseProvider(p Provider) {
	providers[p.Name()] = p
}

func GetProvider(name string) (Provider, error) {
	p, ok := providers[name]

	if !ok || p == nil {
		return nil, fmt.Errorf("provider '%v' is not supported", name)
	}

	return p, nil
}
