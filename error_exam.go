package main

import (
	"fmt"
	"math"
)

// Định nghĩa kiểu ErrNegativeSqrt
type ErrNegativeSqrt float64

// Implement phương thức Error() cho ErrNegativeSqrt
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sửa hàm Sqrt
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	// Gọi Sqrt với giá trị hợp lệ và không hợp lệ
	fmt.Println(Sqrt(2))  // Kết quả: (1.4142135623730951, <nil>)
	fmt.Println(Sqrt(-2)) // Kết quả: (0, cannot Sqrt negative number: -2)
}
