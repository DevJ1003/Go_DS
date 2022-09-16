package main

import "fmt"

func price(item string) int {

	switch item {
	case "apple":
		return 10
	case "orange":
		return 20
	case "carrot":
		return 5

	default:
		return 0
	}

}

func main() {

	fmt.Println(price("laptop"))

	// for i := 0; i <= 10; i++ {
	// 	if i < 5 {
	// 		fmt.Println(i)
	// 	} else {
	// 		fmt.Println("I can't print i")
	// 	}
	// }

	// for i := 0; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

}
