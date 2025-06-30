package controllers

import (
	"context"
	"fmt"
	"github.com/tuusuario/hospital-m/config"
	"github.com/tuusuario/hospital-m/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsuarios(c *fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), "SELECT id_usuario, nombre, rol FROM usuarios")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener usuarios"})
	}
	defer rows.Close()

	var usuarios []models.Usuario
	for rows.Next() {
		var u models.Usuario
		if err := rows.Scan(&u.IdUsuario, &u.Nombre, &u.Rol); err != nil {
			fmt.Println(err)
			continue
		}
		usuarios = append(usuarios, u)
	}

	return c.JSON(usuarios)
}

func CreateUsuario(c *fiber.Ctx) error {
	var user models.Usuario
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	_, err := config.DB.Exec(
		context.Background(),
		"INSERT INTO usuarios (nombre, rol) VALUES ($1, $2)",
		user.Nombre, user.Rol,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo crear usuario"})
	}

	return c.JSON(fiber.Map{"message": "Usuario creado correctamente"})
}
