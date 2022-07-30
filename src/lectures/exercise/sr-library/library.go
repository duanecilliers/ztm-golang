//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Member struct {
	name string
}
type Check struct {
	time time.Time
	by Member
}
type Book struct {
	checkout []Check
	checkin []Check
}

type Library struct {
	members map[string]Member
	books map[string]Book
}

func checkin(bookName string, member Member, library *Library) {
	book := library.books[bookName]
	book.checkin = append(library.books[bookName].checkin, Check{time: time.Now(), by: member})
	library.books[bookName] = book
}

func checkout(bookName string, member Member, library *Library) {
	book := library.books[bookName]
	book.checkout = append(library.books[bookName].checkout, Check{time: time.Now(), by: member})
	library.books[bookName] = book
}

func printLibrary(library *Library) {
	fmt.Println("Library:", *library)
}

func main() {
	books := make(map[string]Book)
	members := make(map[string]Member)

	books["Absalom, Absalom!"] = Book{}
	books["A Time to Kill"] = Book{}
	books["The House of Mirth"] = Book{}
	books["East of Eden"] = Book{}

	members["Anne"] = Member{name: "Anne"}
	members["Jake"] = Member{name: "Jake"}
	members["Dina"] = Member{name: "Dina"}

	library := Library{books: books, members: members}
	printLibrary(&library)

	checkout("Absalom, Absalom!", members["Anne"], &library)
	fmt.Println("After checkout:")
	printLibrary(&library)

	checkin("Absalom, Absalom!", members["Anne"], &library)
	fmt.Println("After checkin:")
	printLibrary(&library)
}
