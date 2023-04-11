package main

import (
	"chapter2-projek/database"
	"chapter2-projek/models"
	"chapter2-projek/routers"

)

func main()  {
	db := database.StartDB()
	db.AutoMigrate(&models.Book{})

	g := routers.StartServer(db)

	g.Run(":8000")
}