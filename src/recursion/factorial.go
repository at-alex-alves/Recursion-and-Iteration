package recursion

/*
	Calculates the factorial of the passed number using recursion.

	Args:
		number (int): The number that will have the factorial calculated.

	Returns:
		int: The passed number times the calling of the same function, with the number minus one as parameter.
*/
func recursiveFactorial(number int) int {
	if number == 1 {
		return 1
	}

	return number * recursiveFactorial(number-1)
}

/*
	Calculates the factorial of the passed number using iteration.

	Args:
		number (int): The number that will have the factorial calculated.

	Returns:
		int: The factorial of the passed number.
*/
func iterativeFactorial(number int) int {
	result := 1

	for i := number; i >= 1; i-- {
		result = result * i
	}

	return result
}
