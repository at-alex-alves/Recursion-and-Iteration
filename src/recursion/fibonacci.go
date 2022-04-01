package recursion

/*
	Calculates the fibonacci of the passed number using recursion.

	Args:
		number (int): The number that will have the fibonacci calculated.

	Returns:
		int: The calling of the same funtion using the passed number minus one, plus 
			the calling of the same funtion using the passed number minus two
*/
func recursiveFibonacci(number int) int {
	if number < 2 {
		return number
	} else {
		return recursiveFibonacci(number-1) + recursiveFibonacci(number-2)
	}
}

/*
	Calculates the fibonacci of the passed number using recursion.

	Args:
		number (int): The number that will have the fibonacci calculated.

	Returns:
		int: The fibonacci of the passed number.
*/
func iterativeFibonacci(number int) int {
	a, b := 0, 1

	for i := 0; i < number; i++ {
		a, b = b, a+b
	}

	return a
}
