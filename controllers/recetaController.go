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

    idMedico, ok := c.Locals("user_id").(int)
    if !ok {
        return c.Status(401).JSON(fiber.Map{"error": "Token inválido"})
    }

    _, err := config.DB.Exec(context.Background(),
        "INSERT INTO recetas (fecha, medicamento, dosis, id_medico, id_paciente, id_consultorio) VALUES ($1, $2, $3, $4, $5, $6)",
        r.Fecha, r.Medicamento, r.Dosis, idMedico, r.IdPaciente, r.IdConsultorio)

    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Error al guardar receta"})
    }

    return c.JSON(fiber.Map{"message": "Receta creada"})
}


func UpdateReceta(c *fiber.Ctx) error {
	id := c.Params("id")
	var r models.Receta
	if err := c.BodyParser(&r); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	_, err := config.DB.Exec(context.Background(),
		"UPDATE recetas SET fecha=$1, medicamento=$2, dosis=$3 WHERE id_receta=$4",
		r.Fecha, r.Medicamento, r.Dosis, id,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al actualizar receta"})
	}

	return c.JSON(fiber.Map{"message": "Receta actualizada"})
}

func DeleteReceta(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := config.DB.Exec(context.Background(), "DELETE FROM recetas WHERE id_receta=$1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar receta"})
	}
	return c.JSON(fiber.Map{"message": "Receta eliminada"})
}

