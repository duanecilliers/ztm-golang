//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y int
}
type Rectangle struct {
	a Coordinate
	b Coordinate
}

func width(rec Rectangle) int {
	return rec.b.x - rec.a.x
}

func length(rec Rectangle) int {
	return rec.a.y - rec.b.y
}

func area(rec Rectangle) int {
	return length(rec) * width(rec)
}

func perimeter(rec Rectangle) int {
	return (width(rec) * 2) + (length(rec) * 2)
}

func printInfo(rec Rectangle) {
	fmt.Println("Area is", area(rec))
	fmt.Println("Perimeter is", perimeter(rec))
}

func main() {
	rec := Rectangle{ a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	printInfo(rec)

	rec.a.y *= 2
	rec.b.x *= 2

	printInfo((rec))
}
