Sistema de Citas y Reportes Hospitalarios

Este proyecto es un backend desarrollado en **Go** utilizando el framework **Fiber**, junto con **Supabase** como plataforma de base de datos. Está diseñado para gestionar la información de un hospital, incluyendo usuarios, consultorios, horarios, expedientes médicos, consultas y recetas, todo con un enfoque en la seguridad y privacidad de los datos.

Tecnologías utilizadas

- Go: Lenguaje de programación principal
- Fiber: Framework web inspirado en Express.js
- Supabase: Backend as a Service (BaaS) con PostgreSQL
- JWT: Autenticación basada en tokens

Estructura del proyecto
hospital-m/
│
├── config/ # Configuración de la base de datos
├── controllers/ # Lógica de los endpoints (usuarios, consultas, etc.)
├── middleware/ # Middlewares como autenticación JWT
├── models/ # Modelos de datos
├── routes/ # Definición de rutas
├── utils/ # Funciones auxiliares
├── .env # Variables de entorno
├── go.mod # Módulo de Go
└── main.go # Punto de entrada principal


Variables de entorno (`.env`)
Crea un archivo `.env` en la raíz del proyecto con el siguiente contenido:
DATABASE_URL=postgres://usuario:contraseña@host.supabase.co:5432/nombre_db
JWT_SECRET=secreto
PORT=3000

Funcionalidades principales

- Registro y autenticación de usuarios
- Asignación de roles (paciente, médico, enfermera, admin)
- Gestión de consultorios y horarios
- Creación de expedientes médicos
- Registro de consultas y diagnósticos
- Generación de recetas médicas
- Rutas protegidas con autenticación JWT

Crear un usuario
http
POST /api/usuarios
Obtener usuarios (requiere token JWT)
http
GET /api/usuarios
Authorization: Bearer <token>

Cómo correr el proyecto:
git clone https://github.com/tuusuario/hospital-m.git
cd hospital-m
Dependecias:
go mod tidy
Ejecuta el servidor:
go run main.go



  
