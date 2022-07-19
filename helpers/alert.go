package helpers

const (
	SUCCESS = "success"
	ERROR   = "danger"
)

type Alert struct {
	Message string
	Type    string
}

func NewAlert(message string, status string) Alert {
	return Alert{
		Message: message,
		Type:    status,
	}
}
