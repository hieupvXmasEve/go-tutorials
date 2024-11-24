package main

import (
	"fmt"
	"math"
)

// Booleans
// Integers (int, int8, int16, int32, int64),
// rune and byte are aliases for int, byte is an alias for uint8, rune is an alias for int32

// Unsigned integers (uint, uint8, uint16, uint32, uint64)
// Floats (float32, float64)
// Complex numbers (complex64, complex128): a + bi
// Strings (string)
// Pointers
// Structs

func main() {
	str := "initial"
	fmt.Println(str)

	var b, c int = 1, 2
	fmt.Println(b, c)

	bool := true
	fmt.Println(bool)

	var e int // 0
	fmt.Println(e)

	f := "Apple"
	fmt.Println(f)

	// Constants

	const num1 = 500000000
	const num2 = 3e20 / num1
	fmt.Println(num2)
	fmt.Println(int64(num2))
	fmt.Println(math.Sin(num1))
}
