package main

import (
	"fmt"
)

func main() {

	numbers := []int{1, 2, 3, 4, 5, 8, 12}
	// fmt.Println(numbers[len(numbers)-1])
	numbers = append(numbers, 100, 9, 50, 13)
	fmt.Println(numbers)

}
