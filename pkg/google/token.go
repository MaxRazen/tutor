package google

type Token struct {
	accessToken  string
	refreshToken string
	expiresIn    int64
}

func (t *Token) AccessToken() string {
	return t.accessToken
}

func (t *Token) RefreshToken() string {
	return t.refreshToken
}

func (t *Token) ExpiresIn() int64 {
	return t.expiresIn
}
