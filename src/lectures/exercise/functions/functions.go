//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Println("Hello", name)
}

//* Write a function that returns any message, and call it from within
func message() string {
	return "Hello from message()"
}

//* Write a function to addThree 3 numbers together, supplied as arguments, and
//  return the answer
func addThree(a, b, c int) int {
	return a + b + c
}

//* Write a function that returns any number
func returnNumber() int {
	return 8
}

//* Write a function that returns any two numbers
func returnTwoNumbers() (int, int) {
	return 5, 6
}

func main() {
	greet("Duane")
	fmt.Println(message())

	//* Add three numbers together using any combination of the existing functions.
	var a, b = returnTwoNumbers()
	var sum = addThree(returnNumber(), a, b)
	fmt.Println("The sum of the numbers is", sum)

}
