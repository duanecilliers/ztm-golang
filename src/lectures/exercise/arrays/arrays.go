//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name string
	price float64
}

func printInfo(products [4]Product) {
	var sum float64
	var totalProducts int
	for i := 0; i < len(products); i++ {
		item := products[i]
		sum += float64(item.price)
		if item.name != "" {
			totalProducts += 1
		}
	}
	fmt.Println("Last product", products[totalProducts - 1])
	fmt.Println("Total products:", totalProducts)
	fmt.Println("Sum:", sum)
}

func main() {
	products := [4]Product {
		{name: "P1", price: 10.99},
		{name: "P2", price: 15.99},
		{name: "P3", price: 14.99},
	}
	printInfo(products)
	products[3] = Product{name:  "P4", price: 21.99}
	printInfo(products)
}
