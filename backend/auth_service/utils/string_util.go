package utils

import (
	"fmt"
	"strings"
	"unicode"
)

// Fungsi generic untuk tipe yang memiliki metode String()
func NormalizeString[T fmt.Stringer](input T) string {
	str := input.String()
	return removeSpacesAndLower(str)
}

// Fungsi generic untuk tipe string atau ~string
func NormalizeAnyString[T ~string](input T) string {
	return removeSpacesAndLower(string(input))
}

// Fungsi helper untuk menghapus spasi dan mengubah ke lowercase
func removeSpacesAndLower(s string) string {
	var result strings.Builder

	for _, r := range s {
		if !unicode.IsSpace(r) {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	return result.String()
}
