package recursion

// Calculates the fibonacci of the passed number using recursion.
func recursiveFibonacci(number int) int {
	if number < 2 {
		return number
	} else {
		return recursiveFibonacci(number-1) + recursiveFibonacci(number-2)
	}
}

// Calculates the fibonacci of the passed number using recursion.
func iterativeFibonacci(number int) int {
	a, b := 0, 1

	for i := 0; i < number; i++ {
		a, b = b, a+b
	}

	return a
}
