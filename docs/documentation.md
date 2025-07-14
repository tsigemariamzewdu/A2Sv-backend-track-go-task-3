# 📚 Library Management System Documentation

## 🧾 Overview

This is a console-based library management system built using the Go programming language. It demonstrates key Go concepts such as:

- Structs and slices
- Maps and interfaces
- Methods and pointer receivers
- Modular folder structure
- Console interaction

The system allows a librarian or user to manage books and members by adding, borrowing, returning, and listing books via a simple command-line interface.

## 📁 Folder Structure

library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   └── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod



## 🕹 Console Interface

The console interface is implemented in `controllers/library_controller.go` and provides a menu-driven interaction allowing users to:

* Add members and books
* Borrow and return books
* List available books
* List borrowed books by member

Users interact via numeric menu choices entered into the command line.



## 🔧 Technologies and Concepts Used

* Structs and slices in Go
* Maps for efficient data storage and retrieval
* Interfaces for abstraction and flexibility
* Methods with pointer receivers
* Modular package structure
* Console I/O with `fmt` and `bufio`



## 👨‍💻 How to Run

1. Make sure Go is installed on your system (`go version`).
2. Navigate to the project root directory.
3. Run the program using:


go run main.go


4. Follow the on-screen menu to manage library operations.



## ✅ Summary

This project demonstrates a practical application of Go fundamentals by creating a library management system that supports:

* Struct and interface usage
* Map and slice data structures
* Modular code organization
* Interactive console-based user interface




