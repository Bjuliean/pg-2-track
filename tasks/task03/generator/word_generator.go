package generator

import "math/rand"

const(
	minAscii = 97
	maxAscii = 122
)

func generateWord(wordLen int) string {
	resRaw := make([]rune, 0, wordLen)

	for range wordLen {
		resRaw = append(resRaw, rune(rand.Intn(maxAscii - minAscii + 1) + minAscii))
	}

	return string(resRaw)
}