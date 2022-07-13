package helpers

import "math/rand"

func RandomNumber(min int, max int) int {
	return rand.Intn(max-min) + min
}
