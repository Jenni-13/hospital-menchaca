package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuusuario/hospital-m/controllers"
)

// GET /api/usuarios
func GetUsuarios(c *fiber.Ctx) error {
	return controllers.GetUsuarios(c)
}

// POST /api/usuarios
func CreateUsuario(c *fiber.Ctx) error {
	return controllers.CreateUsuario(c)
}

// DELETE /api/usuarios/:id
func DeleteUsuario(c *fiber.Ctx) error {
	return controllers.DeleteUsuario(c)
}

func UpdateUsuario(c *fiber.Ctx) error {
	return controllers.UpdateUsuario(c)
}


