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

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn time.Time
}

type Member struct {
	name Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total int
	lent int
}

type Library struct {
	members map[Name]Member
	books map[Title]BookEntry
}

func checkin() {
}

func checkout() {
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "not yet returned"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMembersAudit(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lent:", book.lent)
	}
	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("book not part of library")
		return false
	}
	if book.lent == book.total {
		fmt.Println("No more books available to lend")
		return false
	}

	book.lent += 1
	library.books[title] = book
	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("book not part of library")
		return false
	}
	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out this book")
		return false
	}

	book.lent -= 1
	library.books[title] = book
	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	library := Library{
		books: make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Absalom, Absalom!"] = BookEntry{
		total: 4,
		lent: 0,
	}
	library.books["A Time to Kill"] = BookEntry{
		total: 3,
		lent: 0,
	}
	library.books["The House of Mirth"] = BookEntry{
		total: 2,
		lent: 0,
	}
	library.books["East of Eden"] = BookEntry{
		total: 1,
		lent: 0,
	}

	library.members["Duane"] = Member{"Duane", make(map[Title]LendAudit)}
	library.members["Anne"] = Member{"Anne", make(map[Title]LendAudit)}
	library.members["Jake"] = Member{"Jake", make(map[Title]LendAudit)}

	fmt.Println("\n Initial:")
	printLibraryBooks(&library)
	printMembersAudit(&library)

	member := library.members["Duane"]
	checkedOut := checkoutBook(&library, "Absalom, Absalom!", &member)
	if checkedOut {
		printLibraryBooks(&library)
		printMembersAudit(&library)
	}

	returned := returnBook(&library, "Absalom, Absalom!", &member)
	if returned {
		printLibraryBooks(&library)
		printMembersAudit(&library)
	}
}
