package google

type Profile struct {
	id     string
	name   string
	email  string
	avatar string
}

func (p *Profile) ID() string {
	return p.id
}

func (p *Profile) Name() string {
	return p.name
}

func (p *Profile) Email() string {
	return p.email
}

func (p *Profile) Avatar() string {
	return p.avatar
}
