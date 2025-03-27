package generator

import "math/rand"

func generatePrice(min int, max int) int {
	return rand.Intn(max - min + 1) + min
}