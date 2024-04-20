package main

import (
	"fmt"

	"syreclabs.com/go/faker"
)

type User struct {
	name         string
	address      string
	emailAddress string
	phoneNumber  string
}

func main() {
	var users [5]User
	matrice := [3][3]int{
		{3, -1, 0},
		{0, -1, 1},
		{1, 7, 2},
	}

	fmt.Println("Printing a multidimensional array")

	for rowIndex, row := range matrice {
		for colIndex, val := range row {
			fmt.Printf("[%d][%d] => (%d)   ", rowIndex, colIndex, val)
		}
		fmt.Println("\n")
	}

	for i := 0; i < len(users); i++ {
		users[i] = User{
			emailAddress: faker.Internet().Email(),
			name:         faker.Name().FirstName() + " " + faker.Name().LastName(),
			address:      faker.Address().City() + ", " + faker.Address().Country(),
			phoneNumber:  faker.PhoneNumber().CellPhone(),
		}
	}

	for i, user := range users {
		if user.address != "" {
			fmt.Printf("%d => %v\n\a", i+1, user.emailAddress)
		} else {
			continue
		}
	}
}
