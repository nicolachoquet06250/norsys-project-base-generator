package helpers

type MaybeErrorFunc = error
type ErrorFunc[T interface{}] func(err error) *T

func MaybeError[T any](err MaybeErrorFunc, f2 ErrorFunc[T]) *T {
	var r *T
	if err != nil {
		r = f2(err)
	}
	return r
}
