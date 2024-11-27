package exercise

import (
	"bufio"
	"os"
)

var input string = ""

func ReadingString() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
}

func PrintString() string {
	return input
}
