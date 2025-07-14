package services

import (
    "fmt"
    "Library_management/models"
)


type LibraryManager interface{
	
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(boodID int,memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberId int) []models.Book
	AddMember(member models.Member)
}

type Library struct{
	Books map[int]models.Book
	Members map[int]models.Member


}
func NewLibrary()*Library{
	return &Library{
		Books: make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}
func (l *Library)AddMember(member models.Member){
	l.Members[member.ID]=member
}
func (l *Library) AddBook(book models.Book){
	l.Books[book.ID]=book
}
func (l *Library)RemoveBook(bookID int){
	delete(l.Books,bookID)
}
func (l *Library)BorrowBook(bookID int,memberID int)error{
	book,bookexist := l.Books[bookID]
	if !bookexist{
		return fmt.Errorf("book with ID %v doesn't exist",bookID)
	}
	if book.Status==models.Borrowed{
		return fmt.Errorf("book with ID %v is already borrowed",bookID)
		
	}

	member,memberexist:= l.Members[memberID]
	if !memberexist{
		return fmt.Errorf("member with Id %v doesn't exits",memberID)
	}

	member.BorrowedBooks=append(member.BorrowedBooks, book)

	book.Status=models.Borrowed
	l.Books[bookID]=book

	return nil


}

func (l *Library) ReturnBook(bookID int,memberID int) error{
	book,bookexist:= l.Books[bookID]

	if !bookexist{
		return fmt.Errorf("book with ID %v doesn't exist",bookID)
	}

	book.Status=models.Available
	l.Books[bookID]=book
	return nil
}

func (l* Library) ListAvailableBooks()[]models.Book{
	var availablebooks []models.Book

	for _, book := range l.Books{
		if book.Status==models.Available{
			availablebooks=append(availablebooks, book)
		}
	}
	return availablebooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, memberexist := l.Members[memberID]
	if !memberexist {
		return []models.Book{}
	}

	return member.BorrowedBooks
}