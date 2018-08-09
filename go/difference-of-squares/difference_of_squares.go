package diffsquares

// SquareOfSums ...
func SquareOfSums(number int) int {
	var sum int

	for i := 1; i <= number; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares ...
func SumOfSquares(number int) int {
	var square int

	for i := 1; i <= number; i++ {
		square += i * i
	}

	return square
}

// Difference ...
func Difference(number int) int {

	SuOSq := SumOfSquares(number)
	SqOSm := SquareOfSums(number)

	return SqOSm - SuOSq
}
