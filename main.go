package main

import (
	"fmt"

	"github.com/JoseHurtadoBaeza/CRM-GolangFiber/database"
	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {

	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)

}

func initDatabase() {

	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")

}

func main() {

	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()

}
