package main

import (
	"github.com/tuusuario/hospital-m/config"
	"github.com/tuusuario/hospital-m/routes"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("utils/.env")
	if err != nil {
		log.Fatal("Error cargando .env")
	}

	config.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
