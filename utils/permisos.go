package utils

import (
	"context"
	"fmt"

	"github.com/tuusuario/hospital-m/config"
)

func ObtenerPermisosPorRol(rol string) ([]string, error) {
	var permisos []string

	row := config.DB.QueryRow(context.Background(), "SELECT permisos FROM permisos WHERE rol = $1", rol)

	err := row.Scan(&permisos)
	if err != nil {
		fmt.Println("Error al obtener permisos:", err)
		return nil, err
	}

	return permisos, nil
}
