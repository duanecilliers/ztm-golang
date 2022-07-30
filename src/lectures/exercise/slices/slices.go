//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printContents(assembly []Part) {
	fmt.Println()
	for i := 0; i < len(assembly); i++ {
		item := assembly[i]
		fmt.Println(item)
	}
}

func main() {
	assembly := make([]Part, 3)
	assembly[0] = "Wrench"
	assembly[1] = "Screw driver"
	assembly[2] = "Drill"
	printContents(assembly)
	assembly = append(assembly, "Torch", "Door")
	printContents(assembly)
	assembly = assembly[3:]
	printContents(assembly)
}
