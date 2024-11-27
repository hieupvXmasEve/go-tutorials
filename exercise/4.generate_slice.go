package exercise

import (
	"strings"
)

// Create a slice from comma-seperated input string
// Write a program which accepts a sequence of comma-separated numbers from console and generate an slice out of them. Return the slice.
// Suppose the following input is supplied to the program: 34, 67, 55, 33, 12, 98.
// Then, the output should be: [34 67 55 33 12 98]

func GenerateSlice(s string) []string {
	if s == "" {
		return nil
	}
	// Split the input string into an array of strings
	var split []string
	for _, num := range strings.Split(s, ",") {
		split = append(split, num)
	}
	return split
}
