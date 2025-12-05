package encoder

import (
	"strings"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Encode converts a number to base62 string
func Encode(num uint64) string {
	if num == 0 {
		return "0"
	}

	var encoded strings.Builder
	base := uint64(len(base62Chars))

	for num > 0 {
		remainder := num % base
		encoded.WriteByte(base62Chars[remainder]) // Converts remainder into base 62 and appends it
		num = num / base
	}

	// Reverse the string
	result := encoded.String()
	return reverse(result)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
