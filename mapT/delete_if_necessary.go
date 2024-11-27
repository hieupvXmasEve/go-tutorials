package MapT

import "fmt"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	if users[name].scheduledForDeletion {
		delete(users, name)
		return true, nil
	}
	return false, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func main() {
	users := make(map[string]user)
	users["Alice"] = user{"Alice", 1234567890, false}
	users["Bob"] = user{"Bob", 9876543210, false}
	users["Charlie"] = user{"Charlie", 5551234567, true}
	users["David"] = user{"David", 4321098765, false}

	fmt.Println(users)

	deleted, err := deleteIfNecessary(users, "Charlie")
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted:", deleted)
	fmt.Println("users", users)

	deleted, err = deleteIfNecessary(users, "Bob")
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted:", deleted)
	fmt.Println("users", users)
}
