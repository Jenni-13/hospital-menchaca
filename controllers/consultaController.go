package controllers

import (
	"context"
	"fmt"
    "database/sql" 
	"github.com/gofiber/fiber/v2"
	"github.com/tuusuario/hospital-m/config"
	"github.com/tuusuario/hospital-m/models"
)

func GetConsultas(c *fiber.Ctx) error {
    userId, ok := c.Locals("user_id").(int)
	fmt.Println("🧪 user_id extraído del token:", userId, "ok:", ok)
    if !ok {
        return c.Status(401).JSON(fiber.Map{"error": "Usuario no autenticado"})
    }

    rows, err := config.DB.Query(context.Background(),
        "SELECT id_consulta, tipo, diagnostico, costo, id_paciente, id_medico, id_consultorio, id_horario, fecha FROM consultas WHERE id_paciente=$1", userId)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Error al obtener consultas"})
    }
    defer rows.Close()

    var lista []fiber.Map
    for rows.Next() {
        var c models.Consulta
        if err := rows.Scan(&c.IdConsulta, &c.Tipo, &c.Diagnostico, &c.Costo,
            &c.IdPaciente, &c.IdMedico, &c.IdConsultorio, &c.IdHorario, &c.Fecha); err != nil {
            continue
        }

        // Convertir fecha si no es nula
        fechaStr := ""
        if c.Fecha.Valid {
            fechaStr = c.Fecha.Time.Format("2006-01-02")
        }

        lista = append(lista, fiber.Map{
            "id_consulta":    c.IdConsulta,
            "tipo":           c.Tipo,
            "diagnostico":    c.Diagnostico,
            "costo":          c.Costo,
            "id_paciente":    c.IdPaciente,
            "id_medico":      c.IdMedico,
            "id_consultorio": c.IdConsultorio,
            "id_horario":     c.IdHorario,
            "fecha":          fechaStr,
        })
    }

    return c.JSON(lista)
}

func CreateConsulta(c *fiber.Ctx) error {
    var input struct {
        Tipo      string `json:"tipo"`
        IdHorario int    `json:"id_horario"`
        Fecha     string `json:"fecha"`
    }

    if err := c.BodyParser(&input); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
    }

    idPaciente, ok := c.Locals("user_id").(int)
    if !ok {
        return c.Status(401).JSON(fiber.Map{"error": "Token inválido"})
    }

    var idMedico, idConsultorio int
    err := config.DB.QueryRow(context.Background(),
        "SELECT id_medico, id_consultorio FROM horarios WHERE id_horario = $1",
        input.IdHorario).Scan(&idMedico, &idConsultorio)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Horario no válido"})
    }

    _, err = config.DB.Exec(context.Background(),
        `INSERT INTO consultas (tipo, fecha, diagnostico, costo, id_paciente, id_medico, id_consultorio, id_horario)
         VALUES ($1, $2, '', 0.0, $3, $4, $5, $6)`,
        input.Tipo, input.Fecha, idPaciente, idMedico, idConsultorio, input.IdHorario)

    if err != nil {
        fmt.Println("Error al guardar consulta:", err)
        return c.Status(500).JSON(fiber.Map{"error": "Error al guardar consulta"})
    }

    return c.JSON(fiber.Map{"message": "Consulta agendada con éxito"})
}

func UpdateConsulta(c *fiber.Ctx) error {
    id := c.Params("id")
    var input struct {
        Tipo     string  `json:"tipo"`
        Costo    float64 `json:"costo"`
        Fecha    string  `json:"fecha"`
        IdHorario int    `json:"id_horario"`
    }

    if err := c.BodyParser(&input); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
    }

    _, err := config.DB.Exec(context.Background(),
        `UPDATE consultas
         SET tipo = $1, costo = $2, fecha = $3, id_horario = $4
         WHERE id_consulta = $5`,
        input.Tipo, input.Costo, input.Fecha, input.IdHorario, id)

    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Error al actualizar consulta"})
    }

    return c.JSON(fiber.Map{"message": "Consulta actualizada correctamente"})
}


func DeleteConsulta(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := config.DB.Exec(context.Background(),
		"DELETE FROM consultas WHERE id_consulta = $1", id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar consulta"})
	}

	return c.JSON(fiber.Map{"message": "Consulta eliminada correctamente"})
}

func GetTodasConsultas(c *fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), `
		SELECT 
			c.id_consulta, 
			c.tipo, 
			c.diagnostico, 
			c.costo, 
			c.id_paciente, 
			up.nombre AS nombre_paciente,
			c.id_medico, 
			um.nombre AS nombre_medico,
			c.id_consultorio, 
			c.id_horario, 
			c.fecha
		FROM consultas c
		JOIN usuarios up ON c.id_paciente = up.id_usuario
		JOIN usuarios um ON c.id_medico = um.id_usuario
	`)
	if err != nil {
		fmt.Println("Error al obtener todas las consultas:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener todas las consultas"})
	}
	defer rows.Close()

	var lista []fiber.Map
	for rows.Next() {
		var cons models.Consulta
		var nombrePaciente, nombreMedico string
		var fechaStr sql.NullTime

		if err := rows.Scan(
			&cons.IdConsulta,
			&cons.Tipo,
			&cons.Diagnostico,
			&cons.Costo,
			&cons.IdPaciente,
			&nombrePaciente,
			&cons.IdMedico,
			&nombreMedico,
			&cons.IdConsultorio,
			&cons.IdHorario,
			&fechaStr,
		); err != nil {
			fmt.Println("Error escaneando fila:", err)
			continue
		}

		fechaTexto := ""
		if fechaStr.Valid {
			fechaTexto = fechaStr.Time.Format("2006-01-02")
		}

		lista = append(lista, fiber.Map{
			"id_consulta":     cons.IdConsulta,
			"tipo":            cons.Tipo,
			"diagnostico":     cons.Diagnostico,
			"costo":           cons.Costo,
			"id_paciente":     cons.IdPaciente,
			"nombre_paciente": nombrePaciente,
			"id_medico":       cons.IdMedico,
			"nombre_medico":   nombreMedico,
			"id_consultorio":  cons.IdConsultorio,
			"id_horario":      cons.IdHorario,
			"fecha":           fechaTexto,
		})
	}

	return c.JSON(lista)
}

func GetCitasDelDia(c *fiber.Ctx) error {
	userId, ok := c.Locals("user_id").(int)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No autorizado",
		})
	}

	// Traer solo citas del médico actual y del día de hoy
	rows, err := config.DB.Query(context.Background(), `
		SELECT c.id_consulta, c.tipo, c.fecha, u.nombre AS nombre_paciente
		FROM consultas c
		JOIN usuarios u ON c.id_paciente = u.id_usuario
		WHERE c.id_medico = $1 AND DATE(c.fecha) = CURRENT_DATE
	`, userId)
	if err != nil {
		fmt.Println("Error al obtener citas del día:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener citas"})
	}
	defer rows.Close()

	var citas []fiber.Map
	for rows.Next() {
		var id int
		var tipo string
		var fecha time.Time
		var nombrePaciente string

		if err := rows.Scan(&id, &tipo, &fecha, &nombrePaciente); err != nil {
			continue
		}

		citas = append(citas, fiber.Map{
			"id_consulta":     id,
			"tipo":            tipo,
			"fecha":           fecha.Format("2006-01-02"),
			"nombre_paciente": nombrePaciente,
		})
	}

	return c.JSON(citas)
}
