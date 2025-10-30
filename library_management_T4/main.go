package main

import (
	"library_management/controllers"
	"library_management/services"
)

func main() {
	lib := services.NewLibrary()
	ctrl := controllers.NewController(lib)
	ctrl.Run()
}
