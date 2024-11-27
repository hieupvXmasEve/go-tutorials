package main

import "fmt"

func main() {

	// type Key struct {
	// 	Path, Country string
	// }
	// hits := make(map[Key]int)
	// hits[Key{"/", "vn"}] = 2
	// fmt.Println(hits[Key{"/", "vn1"}])

	// attended := map[string]bool{
	// 	"Ann": true,
	// 	"Joe": true,
	// }
	// person := "Ann1"

	// if attended[person] { // will be false if person is not in the map
	// 	fmt.Println(person, "was at the meeting")
	// }

	fmt.Println([]rune("hieu"))
	names := []string{"hieu", "billy", "billy", "bob", "joe"}
	result := getNameCounts(names)
	for key, innerMap := range result {
		fmt.Printf("  '%c': {\n", key) // Chuyển int thành ký tự (rune)
		for name, count := range innerMap {
			fmt.Printf("    \"%s\": %d,\n", name, count)
		}
		fmt.Println("  },")
	}
}
func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int)
	for _, name := range names {
		// Get the first character as a rune
		firstChar := []rune(name)[0]
		// Initialize the inner map for this first character if it doesn't exist
		if _, exists := nameCounts[firstChar]; !exists {
			nameCounts[firstChar] = make(map[string]int)
		}

		// Increment the count for this name
		nameCounts[firstChar][name]++
	}
	return nameCounts
}
