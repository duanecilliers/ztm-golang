//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active = true
	Inactive = false
)

type SecurityTag bool
type Tag struct {
	name string
	tag SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(tags []Tag) {
	for i := 0; i < len(tags); i++ {
		deactivate(&tags[i].tag)
	}
}

func main() {
	item1 := Tag{name: "1", tag: Active}
	item2 := Tag{name: "2", tag: Active}
	item3 := Tag{name: "3", tag: Active}
	item4 := Tag{name: "4", tag: Active}
	tags := []Tag{item1, item2, item3, item4}

	fmt.Println("Tags:", tags)
	
	deactivate(&tags[0].tag)
	fmt.Println("Tags:", tags)

	checkout(tags)
	fmt.Println("checked out", tags)
}
