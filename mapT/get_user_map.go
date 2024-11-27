package MapT

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("names and phone numbers must be the same length")
	}
	users := make(map[string]user, len(names))
	for i := range names {
		users[names[i]] = user{names[i], phoneNumbers[i]}
	}
	return users, nil
}

type user struct {
	name        string
	phoneNumber int
}

func main() {
	names := []string{"Alice", "Bob", "Charlie", "David"}
	phoneNumbers := []int{1234567890, 9876543210, 5551234567, 4321098765}
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		panic(err)
	}
	for name, user := range users {
		fmt.Printf("%s: %d \n", name, user.phoneNumber)
	}
}
