package raindrops

import (
	"strconv"
)

// Convert a number to a string, the contents of which depend on the number's factors.
func Convert(number int) string {
	var factors string

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			if i == 3 {
				factors += "Pling"
			} else if i == 5 {
				factors += "Plang"
			} else if i == 7 {
				factors += "Plong"
			}
		}
	}

	if factors == "" {
		factors = strconv.Itoa(number)
	}

	return factors
}
