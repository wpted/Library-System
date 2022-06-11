//  A program to manage lending and returning of library books.
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Print out initial library information
//* There must only ever be one copy of the library in memory at any time
//

package main

import (
	"fmt"
	"time"
)

type member struct {
	id            string
	name          string
	booksBorrowed []book
}

const booksCanBorrow = 3

type book struct {
	title    string
	author   string
	isbn     string
	inStock  bool
	checkOut checkOutInfo
	checkIn  checkInInfo
}

type checkOutInfo struct {
	time     []time.Time
	borrower string
}

type checkInInfo struct {
	time     []time.Time
	borrower string
}

type library struct {
	members map[string]member
	books   map[string]book
}

var myLibrary library

func (l library) libraryStatus() {
	numberOfMembers := len(l.members)
	numberOfBooks := len(l.books)
	numberOfBooksInstock := 0

	for _, collection := range l.books {
		if collection.inStock == true {
			numberOfBooksInstock += 1
		}
	}
	fmt.Println("Library has", numberOfMembers, "members.")
	fmt.Println("Library has", numberOfBooks, "books.")
	fmt.Println("Library has", numberOfBooksInstock, "books in stock.")

}

// checkMemberStatus returns a boolean depends on whether the user exist
func checkMemberExist(m member) bool {
	user, exists := myLibrary.members[m.id]

	// If ID exists and the name is the same as in the database
	if exists && user.name == m.name {
		return true
	} else {
		return false
	}
}

// signUp adds user to library member database
func signUp(m *member) {
	// if member not in the library database
	if checkMemberExist(*m) == false {
		// create a map[member.id]member if the member database is empty
		if myLibrary.members == nil {
			myLibrary.members = make(map[string]member)
		}
		// Add member to the member map
		myLibrary.members[m.id] = *m
		fmt.Println("Member", m.name, "started a new membership.")
	} else {
		fmt.Println("Member", m.name, "exist.")
	}
}

// BookExist returns a boolean depends on whether the book exist
func BookExist(title string) bool {
	_, exists := myLibrary.books[title]

	// If title exists in the database
	if exists {
		return true
	} else {
		return false
	}
}

// checkBookStatus checks whether the book is currently in the library via the book title
func checkBookStatus(title string) {
	// if book in library collection
	if BookExist(title) {
		// if inStock status is true
		if myLibrary.books[title].inStock {
			fmt.Println(title, "currently in stock.")
		} else {
			fmt.Println(title, "is currently borrowed.")
		}
	} else {
		fmt.Println("Sorry, we don't have", title, "in our collection.")
	}
}

// addBook adds a new book to the library books database
func addBook(b *book) {
	// Check if the book is already in the library database
	if BookExist(b.title) == false {
		// create a new map[book.title]book if the library database is empty
		if myLibrary.books == nil {
			myLibrary.books = make(map[string]book)
		}
		// turn the inStock status to true
		b.inStock = true
		// adding the new book
		myLibrary.books[b.title] = *b

		fmt.Println("Book /", b.title, "/ added.")

	} else {
		fmt.Println("Book /", b.title, "/ exist.")
	}
}

// checkOut the borrowed book
// when the borrower have borrowed less than 3 books and the borrowed book in stock
func (m *member) checkOut(b *book) {
	// check if member exists
	if checkMemberExist(*m) {

		if myLibrary.books[b.title].inStock && len(m.booksBorrowed) < booksCanBorrow {

			// Change the status of the book
			b.inStock = false
			b.checkOut.time = append(b.checkOut.time, time.Now())
			b.checkOut.borrower = m.name

			// Update the database
			myLibrary.books[b.title] = *b

			m.booksBorrowed = append(m.booksBorrowed, *b)
			fmt.Println(b.title, "check out successfully.")
		} else if len(m.booksBorrowed) == 3 {
			fmt.Println("Sorry you've exceeded the limit of books you can borrow.(3/3)")
		} else if myLibrary.books[b.title].inStock == false {
			fmt.Println(b.title, "not in stock")
		}
	} else {
		// if member doesn't exist, create a new membership then checkout the book
		signUp(m)
		m.checkOut(b)
	}
}

// checkIn the returned book
func (m *member) checkIn(b *book) {
	// check if member exists

	if myLibrary.books[b.title].inStock == false {

		// Change the status of the book
		b.inStock = true
		b.checkIn.time = append(b.checkIn.time, time.Now())
		b.checkIn.borrower = m.name

		// Update the database
		myLibrary.books[b.title] = *b

		// Update the member status, remove the checkin book from the borrowed list
		for idx, borrowedBook := range m.booksBorrowed {
			if borrowedBook.title == b.title {
				m.booksBorrowed = append(m.booksBorrowed[:idx], m.booksBorrowed[idx+1:]...)
			}
		}

		fmt.Println(b.title, "check in successfully.")
	}

}

func main() {
	edward := member{
		id:   "1",
		name: "Edward",
	}
	victoria := member{
		id:   "2",
		name: "Victoria",
	}

	b1 := book{
		title:  "The Image Of City",
		author: "Kevin Lynch",
		isbn:   "4000241389",
	}
	b2 := book{
		title:  "Introducing Python",
		author: "Bill Lubanovic",
		isbn:   "1492051365",
	}
	b3 := book{
		title:  "Mastering Go",
		author: "Mihalis Tsoukalos",
		isbn:   "1801079315",
	}
	b4 := book{
		title:  "Built: The Hidden Stories Behind our Structures",
		author: "Roma Agrawal",
		isbn:   "1635570220",
	}
	b5 := book{
		title:  "Snow Crash",
		author: "Neal Stephenson",
		isbn:   "0553380958",
	}
	addBook(&b1)
	addBook(&b2)
	addBook(&b3)
	addBook(&b4)
	addBook(&b5)

	signUp(&edward)

	edward.checkOut(&b1)
	edward.checkOut(&b2)
	edward.checkOut(&b3)
	edward.checkOut(&b4)

	fmt.Println(edward.booksBorrowed)
	//fmt.Println(b2.checkOut)
	//fmt.Println(myLibrary.books[b1.title].inStock)
	victoria.checkOut(&b1)
	edward.checkIn(&b1)
	fmt.Println(edward.booksBorrowed)
	victoria.checkOut(&b1)

	myLibrary.libraryStatus()

}
