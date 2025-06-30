package controllers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tuusuario/hospital-m/config"
	"github.com/tuusuario/hospital-m/models"
)

func GetRecetas(c *fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), "SELECT id_receta, fecha, medicamento, dosis, id_medico, id_paciente, id_consultorio FROM recetas")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener recetas"})
	}
	defer rows.Close()

	var lista []models.Receta
	for rows.Next() {
		var r models.Receta
		if err := rows.Scan(&r.IdReceta, &r.Fecha, &r.Medicamento, &r.Dosis, &r.IdMedico, &r.IdPaciente, &r.IdConsultorio); err != nil {
			fmt.Println(err)
			continue
		}
		lista = append(lista, r)
	}
	return c.JSON(lista)
}

func CreateReceta(c *fiber.Ctx) error {
	var r models.Receta
	if err := c.BodyParser(&r); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	_, err := config.DB.Exec(context.Background(),
		"INSERT INTO recetas (fecha, medicamento, dosis, id_medico, id_paciente, id_consultorio) VALUES ($1, $2, $3, $4, $5, $6)",
		r.Fecha, r.Medicamento, r.Dosis, r.IdMedico, r.IdPaciente, r.IdConsultorio,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al guardar receta"})
	}
	return c.JSON(fiber.Map{"message": "Receta creada"})
}
