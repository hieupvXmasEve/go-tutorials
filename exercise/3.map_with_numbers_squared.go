package exercise

import "strconv"

//With a given integral number n, write a program to generate a map that contains (i, i*i) such that is an integral number between 1 and n (both included), and then the program should print the map with representation of the value
//Suppose the following input is supplied to the program: 8
//Then, the output should be: map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]

func MapWithNumbersSquared(n int) map[int]string {
	m := make(map[int]string, n)

	for i := 1; i <= n; i++ {
		m[i] = strconv.Itoa(i * i)
	}
	return m
}
