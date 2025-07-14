package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	
)

var DB *pgx.Conn

func ConnectDB() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("❌ Error al conectar a la base de datos:", err)
		os.Exit(1)
	}

	fmt.Println("✅ Conexión a la base de datos exitosa")
	DB = conn
}
