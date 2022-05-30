package recursion

import (
	"fmt"
	"time"
)

var executionTimes []int

// Calculates the processing time for each algorithm based on the passed number.
func run(number int) {
	var functions []func(int) int

	functions = append(functions, recursiveFactorial)
	functions = append(functions, iterativeFactorial)
	functions = append(functions, recursiveFibonacci)
	functions = append(functions, iterativeFibonacci)

	for functionIndex, function := range functions {
		calculateExecutionTime(function, number, functionIndex)
	}
}

func calculateExecutionTime(passedFunc func(int) int, functionParameter int, functionIndex int) {
	start := time.Now()
	passedFunc(functionParameter)
	end := time.Since(start).Nanoseconds()

	executionTimes[functionIndex] += int(end)
}

// Evaluates the performance by getting the mean value of the execution of all algorithms with the passed number.
func EvaluatePerformance(loops int) {
	executionTimes = []int{0, 0, 0, 0}

	for i := 1; i <= loops; i++ {
		run(i)
	}

	fmt.Printf("\nRecursion Algorithms Average in Nanoseconds:\n")

	fmt.Printf(
		"\tFactorial:\n\t\tRecursive: %v\n\t\tIterative: %v\n\n",
		executionTimes[0]/loops,
		executionTimes[1]/loops,
	)

	fmt.Printf(
		"\tFibonacci:\n\t\tRecursive: %v\n\t\tIterative: %v\n",
		executionTimes[2]/loops,
		executionTimes[3]/loops,
	)
}
