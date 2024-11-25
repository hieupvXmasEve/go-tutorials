/*
Exercise: Maps

Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Tách chuỗi thành các từ
	words := strings.Fields("hello world")
	fmt.Println(words)

}
