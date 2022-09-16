package main

import (
	"fmt"
)

func main() {

	user := map[string]string{

		"name":  "Zoey",
		"email": "zoey@deo.com",
		"role":  "Developer",
		"age":   "45",
	}

	// fmt.Println(user)
	// age, ok := user["age"]
	// if ok == true {
	// 	fmt.Println("age:", age)
	// } else {
	// 	fmt.Println("age not found!")
	// }

	for key, value := range user {
		fmt.Println("Key:", key, "Value:", value)
	}

}
