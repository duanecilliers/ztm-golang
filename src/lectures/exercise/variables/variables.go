//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var myName = "Duane"
	fmt.Println("My name is", myName)

	var name string = "Kathy"
	fmt.Println("My name = ", name)

	userName := "admin"
	fmt.Println("User name = ", userName)

	var sum int
	fmt.Println("Sum = ", sum)

	part1, other := 1, 5
	fmt.Println("Part = ", part1, "Other = ", other)

	part2, other := 2, 0
	fmt.Println("Part = ", part2, "Other = ", other)

	sum = part1 + part2
	fmt.Println("Sum = ", sum)

	var(
		lessonName = "Variables"
		lessonType = "Demo"
	)
	
	fmt.Println("lessonName=", lessonName)
	fmt.Println("lessonType=", lessonType)

	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1, word2)
}
