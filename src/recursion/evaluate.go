package recursion

import (
	"fmt"
	"time"
)

var (
	recursiveFactorialTime int
	iterativeFactorialTime int

	recursiveFibonacciTime int
	iterativeFibonacciTime int
)

/*
	Calculates the processing time for each algorithm based on the passed number.

	Params:
		number (int): The number that will be passed fo each algorithm execution.
*/
func run(number int) {
	start := time.Now()
	recursiveFactorial(number)
	end := time.Since(start).Nanoseconds()

	recursiveFactorialTime += int(end)

	start = time.Now()
	iterativeFactorial(number)
	end = time.Since(start).Nanoseconds()

	iterativeFactorialTime += int(end)

	start = time.Now()
	recursiveFibonacci(number)
	end = time.Since(start).Nanoseconds()

	recursiveFibonacciTime += int(end)

	start = time.Now()
	iterativeFibonacci(number)
	end = time.Since(start).Nanoseconds()

	iterativeFibonacciTime += int(end)
}

/*
	Evaluates the performance by getting the mean value of the execution of all algorithms with the passed number.

	Args:
		loops (int): The number of executions that the algorithms will have with the number incrementing with each execution.
*/
func EvaluatePerformace(loops int) {
	recursiveFactorialTime = 0
	iterativeFactorialTime = 0

	recursiveFibonacciTime = 0
	iterativeFibonacciTime = 0

	for i := 1; i <= loops; i++ {
		run(i)
	}

	fmt.Printf("\nRecursion Algorithms Average in Nanoseconds:\n")
	
	fmt.Printf(
		"\tFactorial:\n\t\tRecursive: %v\n\t\tIterative: %v\n\n",
		recursiveFactorialTime/loops,
		iterativeFactorialTime/loops,
	)

	fmt.Printf(
		"\tFibonacci:\n\t\tRecursive: %v\n\t\tIterative: %v\n",
		recursiveFibonacciTime/loops,
		iterativeFibonacciTime/loops,
	)
}