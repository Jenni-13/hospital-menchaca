package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuusuario/hospital-m/handlers"
	"github.com/tuusuario/hospital-m/middleware"
	"github.com/tuusuario/hospital-m/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	// Rutas públicas
	api.Post("/usuarios", handlers.CreateUsuario)
	api.Post("/login", controllers.Login)
	// Rutas protegidas
	api.Get("/usuarios", middleware.Protected(), handlers.GetUsuarios)
	api.Get("/consultorios", middleware.Protected(), controllers.GetConsultorios)
	api.Post("/consultorios", middleware.Protected(), controllers.CreateConsultorio)
	api.Get("/expedientes", middleware.Protected(), controllers.GetExpedientes)
	api.Post("/expedientes", middleware.Protected(), controllers.CreateExpediente)
	api.Get("/consultas", middleware.Protected(), controllers.GetConsultas)
	api.Post("/consultas", middleware.Protected(), controllers.CreateConsulta)
	api.Get("/recetas", middleware.Protected(), controllers.GetRecetas)
	api.Post("/recetas", middleware.Protected(), controllers.CreateReceta)
	
}
