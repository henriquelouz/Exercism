package hamming

import "errors"

//Distance calculates the Hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {

	var distance int

	if len(a) != len(b) {
		return -1, errors.New("Error: Strings length do not match")
	}

	for pos, char := range a {
		if string(b[pos]) != string(char) {
			distance++
		}
	}

	return distance, nil
}
