package auth

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
