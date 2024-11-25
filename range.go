package main

import "fmt"

func main() {
	// 1. Ranging over slices/arrays
	fruits := []string{"apple", "banana", "orange"}
	for index, value := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// 2. Ranging over maps
	scores := map[string]int{"Alice": 90, "Bob": 85}
	for key, value := range scores {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// 3. Ranging over strings (iterates over Unicode code points)
	for index, char := range "Hello" {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	// 4. Using only index or value
	// Skip value with _
	for i, _ := range fruits {
		fmt.Printf("Index: %d\n", i)
	}

	// Skip index by omitting it
	for _, value := range fruits {
		fmt.Printf("Value: %s\n", value)
	}

	// 5. Ranging over channels
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()
	for value := range ch {
		fmt.Printf("Value: %d\n", value)
	}
}

/*
Những điều cần lưu ý khi sử dụng range:

Range trả về 2 giá trị:

Với slice/array: index và value
Với map: key và value
Với string: index và rune (ký tự Unicode)
Với channel: chỉ trả về value


Bạn có thể bỏ qua giá trị không cần thiết bằng dấu gạch dưới (_)
Range tạo ra một bản copy của các phần tử, nếu bạn muốn thay đổi giá trị gốc, cần sử dụng chỉ số index
Với map, thứ tự lặp là ngẫu nhiên
Với channel, range sẽ lặp cho đến khi channel được đóng
*/
