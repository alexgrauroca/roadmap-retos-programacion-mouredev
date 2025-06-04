package factorial

// Factorial calculates the factorial of a given number n
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * Factorial(n-1)
}
