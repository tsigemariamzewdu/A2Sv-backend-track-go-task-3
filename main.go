package main

import (
	"Library_management/controllers"
	"Library_management/services"
)

func main() {
	lib := services.NewLibrary()
	control:=controllers.NewLibraryController(lib)

	control.Run()
}