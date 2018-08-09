package luhn

import (
	"strconv"
	"strings"
)

// Valid ...
func Valid(number string) bool {

	number = strings.Replace(number, " ", "", -1)

	// Check if contains non-numbers || small length
	if _, err := strconv.Atoi(number); err != nil || len(number) <= 1 {
		return false
	}

	arrayNumber := make([]int, len(number))
	var sum int

	for i, j := len(number)-1, 1; i >= 0; i, j = i-1, j+1 {
		arrayNumber[i], _ = strconv.Atoi(string(number[i]))

		if j%2 == 0 {
			arrayNumber[i] = arrayNumber[i] * 2
			if arrayNumber[i] > 9 {
				arrayNumber[i] = arrayNumber[i] - 9
			}
		}

		sum += arrayNumber[i]
	}

	if sum%10 == 0 {
		return true
	}
	return false
}
