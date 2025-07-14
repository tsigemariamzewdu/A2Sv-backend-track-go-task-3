package models

type Status string

const(
	Available Status="Available"
	Borrowed Status="Borrowed"

)
 type Book struct{
	ID int
	Title string
	Author string
	Status Status 
}