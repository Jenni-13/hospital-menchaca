package utils

import (
	"context"
	"log"
	"time"

	"github.com/tuusuario/hospital-m/config"
)

func GuardarLog(idUsuario int, accion, ip, resultado string) {
	_, err := config.DB.Exec(
		context.Background(),
		"INSERT INTO logs (id_usuario, accion, ip, resultado, fecha) VALUES ($1, $2, $3, $4, $5)",
		idUsuario, accion, ip, resultado, time.Now(),
	)

	if err != nil {
		log.Println("❌ Error al guardar log:", err)
	}
}
