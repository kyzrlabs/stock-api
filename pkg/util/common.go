package util

import "strings"

func ToPtr[T any](v T) *T {
	return &v
}

func UpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
