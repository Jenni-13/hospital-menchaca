package controllers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tuusuario/hospital-m/config"
	"github.com/tuusuario/hospital-m/models"
)

func GetExpedientes(c *fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), "SELECT id_expediente, antecedentes, historial_clinico, seguro, id_paciente FROM expediente")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener expedientes"})
	}
	defer rows.Close()

	var lista []models.Expediente
	for rows.Next() {
		var ex models.Expediente
		if err := rows.Scan(&ex.IdExpediente, &ex.Antecedentes, &ex.HistorialClinico, &ex.Seguro, &ex.IdPaciente); err != nil {
			fmt.Println(err)
			continue
		}
		lista = append(lista, ex)
	}
	return c.JSON(lista)
}

func CreateExpediente(c *fiber.Ctx) error {
	var ex models.Expediente
	if err := c.BodyParser(&ex); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	_, err := config.DB.Exec(context.Background(),
		"INSERT INTO expediente (antecedentes, historial_clinico, seguro, id_paciente) VALUES ($1, $2, $3, $4)",
		ex.Antecedentes, ex.HistorialClinico, ex.Seguro, ex.IdPaciente,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al guardar expediente"})
	}
	return c.JSON(fiber.Map{"message": "Expediente creado"})
}
