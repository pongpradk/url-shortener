package encoder

import (
	"strings"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const base uint64 = uint64(len(base62Chars))

// Encode converts a number to base62 string
func Encode(num uint64) string {
	if num == 0 {
		return "0"
	}

	var encoded strings.Builder

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
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}
