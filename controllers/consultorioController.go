package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/tuusuario/hospital-m/config"
	"github.com/tuusuario/hospital-m/models"
)

func GetConsultorios(c *fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), "SELECT * FROM consultorios")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener consultorios"})
	}
	defer rows.Close()

	var lista []models.Consultorio
	for rows.Next() {
		var con models.Consultorio
		if err := rows.Scan(&con.IdConsultorio, &con.Tipo, &con.Ubicacion, &con.Nombre, &con.IdMedico); err != nil {
			continue
		}
		lista = append(lista, con)
	}
	return c.JSON(lista)
}

func CreateConsultorio(c *fiber.Ctx) error {
	var con models.Consultorio
	if err := c.BodyParser(&con); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	_, err := config.DB.Exec(context.Background(),
		"INSERT INTO consultorios (tipo, ubicacion, nombre, id_medico) VALUES ($1, $2, $3, $4)",
		con.Tipo, con.Ubicacion, con.Nombre, con.IdMedico,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al guardar consultorio"})
	}
	return c.JSON(fiber.Map{"message": "Consultorio creado"})
}

func UpdateConsultorio(c *fiber.Ctx) error {
	id := c.Params("id")
	var cs models.Consultorio

	if err := c.BodyParser(&cs); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	_, err := config.DB.Exec(context.Background(),
		`UPDATE consultorios 
		SET tipo = $1, ubicacion = $2, nombre = $3, id_medico = $4 
		WHERE id_consultorio = $5`,
		cs.Tipo, cs.Ubicacion, cs.Nombre, cs.IdMedico, id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al actualizar consultorio"})
	}

	return c.JSON(fiber.Map{"message": "Consultorio actualizado correctamente"})
}


func DeleteConsultorio(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := config.DB.Exec(context.Background(),
		"DELETE FROM consultorios WHERE id_consultorio = $1", id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar consultorio"})
	}

	return c.JSON(fiber.Map{"message": "Consultorio eliminado correctamente"})
}

