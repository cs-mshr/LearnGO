package helpers

import "math/rand"

func RandonNumber(n int) int {
	value := rand.Intn(n)
	return value
}
