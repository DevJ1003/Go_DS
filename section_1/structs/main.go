package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
	Role  string
	Age   int
}

func (u User) Salary() int {
	switch u.Role {
	case "Developer":
		return 100
	case "Architect":
		return 200
	default:
		return 0
	}
}

func (u *User) UpdateEmail(email string) {
	u.Email = email
}

func main() {

	zoey := User{Name: "Zoey", Role: "Developer", Age: 8}
	// deo := User{Name: "Deo", Role: "Architect", Age: 6}

	// fmt.Println(zoey.Name, zoey.Salary())
	// fmt.Println(deo.Name, deo.Salary())
	zoey.UpdateEmail("zoey@gmail.com")
	fmt.Println(zoey)

}
