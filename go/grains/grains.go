package grains

import (
	"errors"
	"math"
)

// Square ...
func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("invalid number")
	}
	squareVal := math.Pow(2, float64(number-1))

	return uint64(squareVal), nil
}

// Total ...
func Total() uint64 {
	var number int

	for i, j := 0, 1; i < 64; i, j = i+1, j*2 {
		number = number + j
	}

	return uint64(number)
}
