package helpers

type AlertStatus string

const (
	SUCCESS AlertStatus = "success"
	ERROR   AlertStatus = "danger"
	INFO    AlertStatus = "info"
)

type Alert struct {
	Message string
	Type    string
}

func NewAlert(message string, status AlertStatus) Alert {
	return Alert{
		Message: message,
		Type:    string(status),
	}
}
