package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuusuario/hospital-m/handlers"
	"github.com/tuusuario/hospital-m/middleware"
	"github.com/tuusuario/hospital-m/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/login", controllers.Login)
	api.Post("/refresh", controllers.RefreshToken)
	api.Get("/logs", middleware.Protected(), middleware.RoleRequired("admin"), handlers.GetLogs)
	api.Post("/register", controllers.RegisterUsuario)


	api.Get("/usuarios", middleware.Protected(), middleware.PermisoRequerido("ver_usuarios"), handlers.GetUsuarios)
    api.Post("/usuarios", middleware.Protected(), middleware.PermisoRequerido("crear_usuario"), handlers.CreateUsuario)
    api.Put("/usuarios/:id", middleware.Protected(), middleware.PermisoRequerido("editar_usuario"), handlers.UpdateUsuario)
    api.Delete("/usuarios/:id", middleware.Protected(), middleware.PermisoRequerido("eliminar_usuario"), handlers.DeleteUsuario)



	api.Get("/consultorios", middleware.Protected(), middleware.RolesAllowed("admin", "medico"), controllers.GetConsultorios)
	api.Post("/consultorios", middleware.Protected(), middleware.RoleRequired("admin"), controllers.CreateConsultorio)
    api.Put("/consultorios/:id", middleware.Protected(), middleware.RoleRequired("admin"), controllers.UpdateConsultorio)
    api.Delete("/consultorios/:id", middleware.Protected(), middleware.RoleRequired("admin"), controllers.DeleteConsultorio)

	api.Get("/expedientes", middleware.Protected(), middleware.RolesAllowed("admin", "medico", "enfermera", "paciente"), controllers.GetExpedientes)
	api.Post("/expedientes", middleware.Protected(), middleware.RolesAllowed("medico", "enfermera"), controllers.CreateExpediente)
    api.Put("/expedientes/:id", middleware.Protected(), middleware.RolesAllowed("medico", "enfermera"), controllers.UpdateExpediente)
    api.Delete("/expedientes/:id", middleware.Protected(), middleware.RolesAllowed("medico", "enfermera"), controllers.DeleteExpediente)
              // Expedientes del usuario logueado
    api.Get("/expedientes/todos", controllers.GetExpedientesConUsuario) // Todos con join



	api.Get("/consultas", middleware.Protected(), middleware.PermisoRequerido("ver_consultas"), controllers.GetConsultas)
    api.Post("/consultas", middleware.Protected(), middleware.PermisoRequerido("crear_consulta"), controllers.CreateConsulta)
    api.Put("/consultas/:id", middleware.Protected(), middleware.PermisoRequerido("crear_consulta"), controllers.UpdateConsulta)
    api.Delete("/consultas/:id", middleware.Protected(), middleware.PermisoRequerido("crear_consulta"), controllers.DeleteConsulta)
    api.Get("/consultas/todas", controllers.GetTodasConsultas)
    api.Get("/medico/citas-dia", middleware.AuthRequired, controllers.GetCitasDelDiaMedico)



	api.Get("/recetas", middleware.Protected(), middleware.RolesAllowed("admin", "medico", "enfermera"), controllers.GetRecetas)
	api.Post("/recetas", middleware.Protected(), middleware.RoleRequired("medico"), controllers.CreateReceta)
	api.Put("/recetas/:id", middleware.Protected(), middleware.RoleRequired("medico"), controllers.UpdateReceta)
    api.Delete("/recetas/:id", middleware.Protected(), middleware.RoleRequired("medico"), controllers.DeleteReceta)

	api.Get("/horarios", middleware.Protected(), middleware.PermisoRequerido("ver_horarios"), controllers.GetHorarios)
    api.Post("/horarios", middleware.Protected(), middleware.PermisoRequerido("crear_horario"), controllers.CreateHorario)
    api.Put("/horarios/:id", middleware.Protected(), middleware.PermisoRequerido("editar_horario"), controllers.UpdateHorario)
    api.Delete("/horarios/:id", middleware.Protected(), middleware.PermisoRequerido("eliminar_horario"), controllers.DeleteHorario)
    api.Get("/horarios-disponibles", middleware.Protected(), controllers.GetHorariosDisponibles)

    

	
}
