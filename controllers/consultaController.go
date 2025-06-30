package controllers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tuusuario/hospital-m/config"
	"github.com/tuusuario/hospital-m/models"
)

func GetConsultas(c *fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), "SELECT id_consulta, tipo, diagnostico, costo, id_paciente, id_medico, id_consultorio, id_horario FROM consultas")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener consultas"})
	}
	defer rows.Close()

	var lista []models.Consulta
	for rows.Next() {
		var cs models.Consulta
		if err := rows.Scan(&cs.IdConsulta, &cs.Tipo, &cs.Diagnostico, &cs.Costo, &cs.IdPaciente, &cs.IdMedico, &cs.IdConsultorio, &cs.IdHorario); err != nil {
			fmt.Println(err)
			continue
		}
		lista = append(lista, cs)
	}
	return c.JSON(lista)
}

func CreateConsulta(c *fiber.Ctx) error {
	var cs models.Consulta
	if err := c.BodyParser(&cs); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	_, err := config.DB.Exec(context.Background(),
		"INSERT INTO consultas (tipo, diagnostico, costo, id_paciente, id_medico, id_consultorio, id_horario) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		cs.Tipo, cs.Diagnostico, cs.Costo, cs.IdPaciente, cs.IdMedico, cs.IdConsultorio, cs.IdHorario,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al guardar consulta"})
	}
	return c.JSON(fiber.Map{"message": "Consulta creada"})
}
