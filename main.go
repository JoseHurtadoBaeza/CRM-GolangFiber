package main

import (
	"fmt"

	"github.com/JoseHurtadoBaeza/CRM-GolangFiber/database"
	"github.com/JoseHurtadoBaeza/CRM-GolangFiber/lead"
	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {

	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)

}

func initDatabase() {

	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Connection opened to database")

	if err := database.DBConn.AutoMigrate(&lead.Lead{}); err != nil {
		panic("Error while migration: " + err.Error())
	}

	fmt.Println("Database Migrated")

}

func main() {

	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")

}
