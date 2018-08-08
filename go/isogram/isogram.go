package isogram

import "strings"

var replacer = strings.NewReplacer("-", "", " ", "")

// IsIsogram ...
func IsIsogram(word string) bool {

	word = replacer.Replace(word)

	for j, letter := range word {
		for i := j + 1; i < len(word); i++ {
			if strings.ToUpper(string(letter)) == strings.ToUpper(string(word[i])) {
				return false
			}
		}
	}

	return true
}
