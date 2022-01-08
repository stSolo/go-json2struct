package helpers

import "strings"

func capitalize(s string) string {
	return strings.Title(s)
}

func removeSCase(s string) string {
	return strings.Replace(s, "_", "", -1)
}

func Normalize(s string) string {
	return capitalize(removeSCase(strings.ToLower(s)))
}
