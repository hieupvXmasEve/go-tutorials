package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0

	for i := 1; i <= 5; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				count++
				fmt.Println(count)
			}
		}()
	}
	time.Sleep(time.Second * 1)
}
