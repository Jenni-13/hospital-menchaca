package controllers

import (
	"context"
	"fmt"
"database/sql" 
	"github.com/gofiber/fiber/v2"
	"github.com/tuusuario/hospital-m/config"
	"github.com/tuusuario/hospital-m/models"
)

func GetExpedientes(c *fiber.Ctx) error {
	  userId, ok := c.Locals("user_id").(int)
    if !ok {
        return c.Status(401).JSON(fiber.Map{"error": "Usuario no autenticado"})
    }
	rows, err := config.DB.Query(context.Background(),
    "SELECT id_expediente, antecedentes, historial_clinico, seguro, id_paciente, estatura, peso, edad FROM expediente WHERE id_paciente = $1",
    userId,
)
if err != nil {
    fmt.Println("ERROR DB:", err) // ⬅️ Añadido para ver el error exacto
    return c.Status(500).JSON(fiber.Map{"error": "Error al obtener expedientes"})
}
	defer rows.Close()

	var lista []models.Expediente
	for rows.Next() {
		var ex models.Expediente
		if err := rows.Scan(&ex.IdExpediente, &ex.Antecedentes, &ex.HistorialClinico, &ex.Seguro, &ex.IdPaciente, &ex.Estatura, &ex.Peso, &ex.Edad,); err != nil {
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
		"INSERT INTO expediente (antecedentes, historial_clinico, seguro, id_paciente, estatura, peso, registro ) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		ex.Antecedentes, ex.HistorialClinico, ex.Seguro, ex.IdPaciente, ex.Estatura, ex.Peso, ex.Edad,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al guardar expediente"})
	}
	return c.JSON(fiber.Map{"message": "Expediente creado"})
}

func UpdateExpediente(c *fiber.Ctx) error {
	id := c.Params("id")
	var exp models.Expediente

	if err := c.BodyParser(&exp); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	_, err := config.DB.Exec(context.Background(),
		`UPDATE expediente 
		 SET peso = $1, estatura = $2, edad = $3
		 WHERE id_expediente = $4`,
		exp.Peso, exp.Estatura, exp.Edad, id,
	)

	if err != nil {
		fmt.Println("ERROR al actualizar:", err) // 👈 Agrega esto para ver el error real
		return c.Status(500).JSON(fiber.Map{"error": "Error al actualizar expediente"})
	}

	return c.JSON(fiber.Map{"message": "Expediente actualizado correctamente"})
}

func DeleteExpediente(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := config.DB.Exec(context.Background(),
		"DELETE FROM expediente WHERE id_expediente = $1", id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar expediente"})
	}

	return c.JSON(fiber.Map{"message": "Expediente eliminado correctamente"})
}

func  GetExpedientesConUsuario(c *fiber.Ctx) error {
    rows, err := config.DB.Query(context.Background(), `
    SELECT 
        e.id_expediente, 
        u.nombre, 
        e.antecedentes, 
        e.historial_clinico, 
        e.seguro, 
        e.id_paciente, 
        e.estatura, 
        e.peso, 
        e.edad
    FROM expediente e
    JOIN usuarioS u ON e.id_paciente = u.id_usuario
`)
if err != nil {
    fmt.Println("ERROR EN CONSULTA SQL:", err) // agrega este log
    return c.Status(500).JSON(fiber.Map{"error": "Error al obtener expedientes"})
}
    defer rows.Close()

    var lista []models.Expediente

    for rows.Next() {
        var ex models.Expediente
        var estatura sql.NullFloat64
        var peso sql.NullFloat64
        var edad sql.NullInt64

        err := rows.Scan(
            &ex.IdExpediente,
            &ex.NombrePaciente,  // << Esto da error si no está en el struct
            &ex.Antecedentes,
            &ex.HistorialClinico,
            &ex.Seguro,
            &ex.IdPaciente,
            &estatura,
            &peso,
            &edad,
        )
        if err != nil {
            fmt.Println("Error al escanear expediente:", err)
            continue
        }

        if estatura.Valid {
            ex.Estatura = estatura.Float64
        }
        if peso.Valid {
            ex.Peso = peso.Float64
        }
        if edad.Valid {
            ex.Edad = int(edad.Int64)
        }

        lista = append(lista, ex)
    }

    return c.JSON(lista)
}
