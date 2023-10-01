package main

import (
	"uts/database"
	"uts/route"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()

	database.Connect()

	//Kumpulan Route Route 
	
	app.Post("/insert", route.InsertData )
	app.Get("/getData", route.GetAllData)
	app.Get("/getDataUser/:id_user", route.GetUserByid)

	app.Delete("/delete/:id_user", route.Delete)
	app.Put("/update/:id_user", route.Update) 

	app.Listen(":3000")	
}