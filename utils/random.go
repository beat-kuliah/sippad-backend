package utils

import "math/rand"

var aplhabets string = "abcdefghijklmnopqrstuvwxyz"

func RandomString(r int) string {
	bits := []rune{}
	k := len(aplhabets)

	for i := 0; i < r; i++ {
		index := rand.Intn(k)
		bits = append(bits, rune(aplhabets[index]))
	}

	return string(bits)
}

func RandomUsername() string {
	return RandomString(8)
}
