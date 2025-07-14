package controllers

import (
	"Library_management/models"
	"Library_management/services"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LibraryController struct{
	lib *services.Library
}

func NewLibraryController (lib *services.Library)*LibraryController{
	return &LibraryController{lib:lib}
}

func (lc *LibraryController) addBook(){
	var id int
	fmt.Println("Enter book ID: ")
	fmt.Scan(&id)

	var title string
	fmt.Println("Enter book title: ")
	fmt.Scan(&title)

	var author string
	fmt.Println("Enter book author: ")
	fmt.Scan(&author)

	book := models.Book{
		ID:id,
		Title: title,
		Author: author,
		Status: models.Available,
	}
	lc.lib.AddBook(book)
	fmt.Println("book added successfully")
}

func (lc *LibraryController) addMember(){
	var id int
	fmt.Println("Enter Member ID: ")
	fmt.Scan(&id)

	var name string
	fmt.Println("Enter Member Name: ")
	fmt.Scan(&name)

	
	member := models.Member{
		ID:id,
		Name: name,
		BorrowedBooks: []models.Book{},

	}
	lc.lib.AddMember(member)
	fmt.Println("Member  added successfully")
}
func (lc *LibraryController) borrowBook(){
	var memberid int
	fmt.Println("Enter Member ID: ")
	fmt.Scan(&memberid)

	var bookid int
	fmt.Println("Enter book ID: ")
	fmt.Scan(&bookid)

	
	
	err:= lc.lib.BorrowBook(bookid,memberid)

	if err !=nil{
		fmt.Println("error: ", err)
	}
	fmt.Println("book borrowed successfully")
}

func (lc *LibraryController) returnBook(){
	var memberid int
	fmt.Println("Enter Member ID: ")
	fmt.Scan(&memberid)

	var bookid int
	fmt.Println("Enter book ID: ")
	fmt.Scan(&bookid)

	
	
	err:= lc.lib.ReturnBook(bookid,memberid)

	if err !=nil{
		fmt.Println("error: ", err)
	}
	fmt.Println("book returned successfully")
}

func (lc *LibraryController) listAvailableBooks(){
	books := lc.lib.ListAvailableBooks()
	if len(books)==0{
		fmt.Println("No available books")
		return
	}
	for _,b := range books{
		fmt.Println(b.ID,b.Title,b.Author)
	}
}
func (lc *LibraryController) listBorrowedBooks(){
	var id int
	fmt.Println("Enter Member ID: ")
	fmt.Scan(&id)

	books:=lc.lib.ListBorrowedBooks(id)
	if len(books)==0{
		fmt.Println("This member has not borrowed a book yet")
		return
	}

	for _,b := range books{
		fmt.Println(b.ID,b.Title,b.Author)
	}

	
	
}



func (lc *LibraryController)Run(){
	reader :=bufio.NewReader(os.Stdin)

	for{
		fmt.Println("\n Welcome to Library Mangement : ")
		fmt.Println("1. Add Member")
		fmt.Println("2. Add Book ")
		fmt.Println("3. Borrow book")
		fmt.Println("4. Return book ")
		fmt.Println("5. List available books")
		fmt.Println("6. List borrowed books")
		fmt.Println("7. exit")

		fmt.Println("choose from those options and to exti press 7")

		input,_:=reader.ReadString('\n')
		choice:= strings.TrimSpace(input)

		switch choice{
		case "1":
			lc.addMember()
		case "2":
			lc.addBook()
		case "3":
			lc.borrowBook()
		case "4":
			lc.returnBook()
		case "5":
			lc.listAvailableBooks()
		case "6":
			lc.listBorrowedBooks()
		case "7":
			fmt.Println("Thank you for using Library management see you next time !")
			return
		default:
			fmt.Println("please press between 1 and 7")
		}
		
	}
}