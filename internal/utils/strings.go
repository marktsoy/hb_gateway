package utils

import "math/rand"

var (
	symbols = "abcdefghijklmnopqrtsuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func StrRandom(length int) string {
	str := ""

	for i := 0; i < length; i++ {
		n := rand.Intn(len(symbols) - 1)
		str += string(symbols[n])
	}

	return str
}
