package ui

type alertType string

const (
	AlertSuccess alertType = "success"
	AlertError   alertType = "error"

	AlertMessageAuthenticationFailed = "User cannot be logged in. Please try again later or contact support"
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
