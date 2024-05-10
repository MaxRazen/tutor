package ui

type alertType string

const (
	AlertSuccess alertType = "success"
	AlertError   alertType = "error"
)

type Alert struct {
	Type    alertType `json:"type"`
	Message string    `json:"message"`
}

func NewAlert(t alertType, m string) Alert {
	return Alert{
		Type:    t,
		Message: m,
	}
}
