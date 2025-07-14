package controllers

import (

    "context"
	"github.com/gofiber/fiber/v2"
	"github.com/tuusuario/hospital-m/config"
	"github.com/tuusuario/hospital-m/models"
	"github.com/tuusuario/hospital-m/utils"
	"golang.org/x/crypto/bcrypt"


)
func RegisterUsuario(c *fiber.Ctx) error {
    var input models.Usuario
    if err := c.BodyParser(&input); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
    }

    // 1. Crear usuario (rol paciente)
    input.Rol = "paciente"

	// Validar longitud y seguridad de contraseña
	if len(input.Password) < 12 {
		return c.Status(400).JSON(fiber.Map{"error": "La contraseña debe tener al menos 12 caracteres"})
	}

    secret, qrURL, err := utils.GenerarMFA(input.Correo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo generar MFA"})
	}

hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo hashear contraseña"})
	}

	_, err = config.DB.Exec(context.Background(),
		"INSERT INTO usuarios (nombre, rol, correo, password, mfa_secret) VALUES ($1, $2, $3, $4, $5)",
		input.Nombre, input.Rol, input.Correo, string(hashedPassword), secret)

		if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo registrar usuario"})
	}
    // 4. Retornar QR URL para escanear
    return c.JSON(fiber.Map{
		"message": "Registro exitoso",
		"qr_url":  qrURL, // esto lo usarás en el frontend para mostrar el QR
	})
}



