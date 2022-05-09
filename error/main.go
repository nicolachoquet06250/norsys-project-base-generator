package error

type IError interface {
	error

	IsError() bool
}

type errorType struct {
	IError

	Err string
}

func (e errorType) Error() string {
	return e.Err
}

func (e errorType) IsError() bool {
	return e.Err != ""
}

func NewError(err string) IError {
	return errorType{
		Err: err,
	}
}
