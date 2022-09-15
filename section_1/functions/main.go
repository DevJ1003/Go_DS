package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func sum(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

// ==============
func main() {

	hello()
	fmt.Println()

	fmt.Println(sum(10, 20))
	fmt.Println(sum(1, 6))
	fmt.Println(sum(14, 200))
	fmt.Println()

	fmt.Println(swap(10, 20))
	fmt.Println()

	a, b := swap(10, 20)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println()

	sum := func(x, y int) int { //closures
		return x + y
	}

	fmt.Println(sum(45, 55))

}
