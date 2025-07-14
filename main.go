package main

import (
	"github.com/tuusuario/hospital-m/config"
	"github.com/tuusuario/hospital-m/routes"
	"github.com/tuusuario/hospital-m/middleware"
	"github.com/gofiber/fiber/v2/middleware/cors"

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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:4200", // Permite Angular
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

    app.Use(middleware.RateLimitMiddleware())
	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
