package helpers

const (
	SUCCESS = "success"
	ERROR   = "danger"
)

type Alert struct {
	Message string
	Type    string
}
