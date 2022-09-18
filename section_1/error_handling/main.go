package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name  string
	Email string
	Role  string
	Age   int
}

func (u User) Salary() (int, error) {

	if u.Role == "" {
		return 0, errors.New(
			fmt.Sprintf("Not abel to handle empty role!"),
		)
	}

	switch u.Role {
	case "Developer":
		return 100, nil
	case "Architect":
		return 200, nil
	}
	return 0, errors.New(
		fmt.Sprintf("Not able to handle %s role!", u.Role),
	)

}

func main() {

	user := User{Name: "Zoey", Role: "", Age: 8}
	salary, err := user.Salary()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Salary:", salary)
	}

}
