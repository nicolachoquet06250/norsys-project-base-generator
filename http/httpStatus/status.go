package httpStatus

import "net/http"

const (
	OK              int = http.StatusOK
	NotFound        int = http.StatusNotFound
	Accepted        int = http.StatusAccepted
	AlreadyReported int = http.StatusAlreadyReported
	BadGateway      int = http.StatusBadGateway
	BadRequest      int = http.StatusBadRequest
	Created         int = http.StatusCreated
	Forbidden       int = http.StatusForbidden
)
