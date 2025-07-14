package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tuusuario/hospital-m/config"
)


func GetLogs(c *fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), "SELECT id_log, id_usuario, accion, ip, resultado, fecha FROM logs")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener logs"})
	}
	defer rows.Close()

	var logs []map[string]interface{}
	for rows.Next() {
		var idLog, idUsuario int
		var accion, ip, resultado string
		var fecha time.Time

		if err := rows.Scan(&idLog, &idUsuario, &accion, &ip, &resultado, &fecha); err != nil {
			continue
		}
		logs = append(logs, fiber.Map{
			"id_log":     idLog,
			"id_usuario": idUsuario,
			"accion":     accion,
			"ip":         ip,
			"resultado":  resultado,
			"fecha":      fecha,
		})
	}

	return c.JSON(logs)
}
