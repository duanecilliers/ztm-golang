package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func great() {
	fmt.Println("Hello from great function!")
}

func main() {
	great()

	dozen:= double(6)
	fmt.Println(dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println("A baker's dozen is", bakersDozen)

	anotherBakersDozen := add(double(6), 1)
	fmt.Println("Another baker's dozen is", anotherBakersDozen)
}
