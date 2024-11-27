package exercise

import (
	"fmt"
	"strconv"
	"strings"
)

// Write a program which can compute the factorial of a given numbers. The results should be printed in a comma-separated sequence on a single line.
func ComputeFactorial(n int) (string, error) {
	if n < 0 {
		return "", fmt.Errorf("Factorial of negative values is not allowed")
	}
	var numbers []string
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i // factorial = factorial * i
		numbers = append(numbers, strconv.Itoa(factorial))
	}
	return strings.Join(numbers, ","), nil
}
