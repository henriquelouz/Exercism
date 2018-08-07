// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate converts a phrase to its acronym
func Abbreviate(s string) string {
	var acronym string

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	arrayStrings := strings.FieldsFunc(strings.ToUpper(s), f)

	for _, word := range arrayStrings {
		acronym += string(word[0])
	}

	return acronym
}
