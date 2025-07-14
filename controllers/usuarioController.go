package controllers

import (
	"context"
	"fmt"
	"github.com/tuusuario/hospital-m/config"
	"github.com/tuusuario/hospital-m/models"
	"github.com/gofiber/fiber/v2"

	
)

func GetUsuarios(c *fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), "SELECT id_usuario, nombre, rol, correo FROM usuarios")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener usuarios"})
	}
	defer rows.Close()

	var usuarios []models.Usuario
	for rows.Next() {
		var u models.Usuario
		if err := rows.Scan(&u.IdUsuario, &u.Nombre, &u.Rol, &u.Correo ); err != nil {
			fmt.Println(err)
			continue
		}
		usuarios = append(usuarios, u)
	}

	return c.JSON(usuarios)
}

func CreateUsuario(c *fiber.Ctx) error {
	var u models.Usuario

	if err := c.BodyParser(&u); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	_, err := config.DB.Exec(context.Background(),
		`INSERT INTO usuarios (nombre, rol) VALUES ($1, $2)`,
		u.Nombre, u.Rol,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo crear el usuario"})
	}

	return c.JSON(fiber.Map{"message": "Usuario creado exitosamente"})
}

func UpdateUsuario(c *fiber.Ctx) error {
	id := c.Params("id")
	var u models.Usuario

	if err := c.BodyParser(&u); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	_, err := config.DB.Exec(context.Background(),
		`UPDATE usuarios SET nombre = $1, rol = $2 WHERE id_usuario = $3`,
		u.Nombre, u.Rol, id,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo actualizar el usuario"})
	}

	return c.JSON(fiber.Map{"message": "Usuario actualizado correctamente"})
}

func DeleteUsuario(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := config.DB.Exec(context.Background(),
		"DELETE FROM usuarios WHERE id_usuario = $1", id,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar el usuario"})
	}

	rowsAffected := result.RowsAffected()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al verificar eliminación"})
	}

	if rowsAffected == 0 {
	return c.Status(404).JSON(fiber.Map{"error": "Usuario no encontrado"})
}

	return c.JSON(fiber.Map{"message": "Usuario eliminado exitosamente"})
}
