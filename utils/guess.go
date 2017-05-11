package utils

import "math/rand"

func Result(local string, visitor string) int {
	r := rand.New(rand.NewSource(2))
	return r.Int()
}
