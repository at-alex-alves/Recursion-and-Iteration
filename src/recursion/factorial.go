package recursion

// Calculates the factorial of the passed number using recursion.
func recursiveFactorial(number int) int {
	if number == 1 {
		return 1
	}

	return number * recursiveFactorial(number-1)
}

// Calculates the factorial of the passed number using iteration.
func iterativeFactorial(number int) int {
	result := 1

	for i := number; i >= 1; i-- {
		result = result * i
	}

	return result
}
