package controllers

import (
	"context"
	"fmt"
	"github.com/tuusuario/hospital-m/config"
	"github.com/tuusuario/hospital-m/models"

	"github.com/gofiber/fiber/v2"
)

// GET /horarios
func GetHorarios(c *fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), "SELECT id_horario, turno, id_consultorio, id_medico FROM horarios")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener horarios"})
	}
	defer rows.Close()

	var lista []models.Horario
	for rows.Next() {
		var h models.Horario
		if err := rows.Scan(&h.IdHorario, &h.Turno, &h.IdConsultorio, &h.IdMedico); err != nil {
			fmt.Println(err)
			continue
		}
		lista = append(lista, h)
	}

	return c.JSON(lista)
}

// POST /horarios
func CreateHorario(c *fiber.Ctx) error {
	var h models.Horario
	if err := c.BodyParser(&h); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	_, err := config.DB.Exec(context.Background(),
		"INSERT INTO horarios (turno, id_consultorio, id_medico) VALUES ($1, $2, $3)",
		h.Turno, h.IdConsultorio, h.IdMedico)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al crear horario"})
	}

	return c.JSON(fiber.Map{"message": "Horario creado"})
}

// PUT /horarios/:id
func UpdateHorario(c *fiber.Ctx) error {
	id := c.Params("id")
	var h models.Horario

	if err := c.BodyParser(&h); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	_, err := config.DB.Exec(context.Background(),
		`UPDATE horarios SET turno = $1, id_consultorio = $2, id_medico = $3 WHERE id_horario = $4`,
		h.Turno, h.IdConsultorio, h.IdMedico, id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al actualizar horario"})
	}

	return c.JSON(fiber.Map{"message": "Horario actualizado"})
}

// DELETE /horarios/:id
func DeleteHorario(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := config.DB.Exec(context.Background(),
		"DELETE FROM horarios WHERE id_horario = $1", id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar horario"})
	}

	return c.JSON(fiber.Map{"message": "Horario eliminado"})
}

func GetHorariosDisponibles(c *fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), `
		SELECT h.id_horario, h.turno, u.nombre AS medico, c.nombre AS consultorio
		FROM horarios h
		JOIN usuarios u ON u.id_usuario = h.id_medico
		JOIN consultorios c ON c.id_consultorio = h.id_consultorio
	`)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener horarios disponibles"})
	}
	defer rows.Close()

	var horarios []fiber.Map
	for rows.Next() {
		var id int
		var turno, medico, consultorio string
		if err := rows.Scan(&id, &turno, &medico, &consultorio); err == nil {
			horarios = append(horarios, fiber.Map{
				"id":          id,
				"turno":       turno,
				"medico":      medico,
				"consultorio": consultorio,
			})
		}
	}

	return c.JSON(fiber.Map{"horarios": horarios})
}

