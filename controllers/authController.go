package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type LoginInput struct {
	Nombre string `json:"nombre"`
	Rol    string `json:"rol"`
}

// POST /api/login
func Login(c *fiber.Ctx) error {
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	// ⚠️ Validación simulada: aquí deberías validar contra la BD real
	if input.Nombre == "" || input.Rol == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Credenciales inválidas"})
	}

	// Crear token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nombre": input.Nombre,
		"rol":    input.Rol,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Expira en 24h
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo generar el token"})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}
