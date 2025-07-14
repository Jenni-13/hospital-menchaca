package controllers

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tuusuario/hospital-m/config"
	"github.com/tuusuario/hospital-m/models"
	"github.com/tuusuario/hospital-m/utils"
	"golang.org/x/crypto/bcrypt"
	"github.com/pquerna/otp/totp"


)

type LoginInput struct {
	Correo   string `json:"correo"`
	Password string `json:"password"`
}


// Login autentica al usuario y genera tokens
func Login(c *fiber.Ctx) error {
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, 400, utils.F01DatosInvalidos, "Datos inválidos")
	}

	// Buscar usuario por correo
	var usuario models.Usuario
	err := config.DB.QueryRow(context.Background(),
		"SELECT id_usuario, nombre, rol, password FROM usuarios WHERE correo = $1",
		input.Correo,
	).Scan(&usuario.IdUsuario, &usuario.Nombre, &usuario.Rol, &usuario.Password)

	if err != nil {
		utils.GuardarLog(0, "login", c.IP(), "fallido: usuario no encontrado")
		return utils.ErrorResponse(c, 401, utils.F02UsuarioNoEncontrado, "Usuario no encontrado")
	}

	
	usuario.Password = strings.TrimSpace(usuario.Password)
	err = bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(input.Password))
	if err != nil {
		utils.GuardarLog(usuario.IdUsuario, "login", c.IP(), "fallido: contraseña incorrecta")
		return utils.ErrorResponse(c, 401, "F03", "Contraseña incorrecta")
	}

	utils.GuardarLog(usuario.IdUsuario, "login", c.IP(), "exitoso")

permisos, err := utils.ObtenerPermisosPorRol(usuario.Rol)
if err != nil {
	return utils.ErrorResponse(c, 500, "F05", "Error al obtener permisos del rol")
}
	// Access Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       usuario.IdUsuario,
		"nombre":   usuario.Nombre,
		"rol":      usuario.Rol,
		"permisos": permisos,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // 2 horas
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return utils.ErrorResponse(c, 500, "F04", "No se pudo generar token de acceso")
	}

	// Refresh Token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"id":  usuario.IdUsuario,
	"rol": usuario.Rol, // <-- Añade esto
	"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
})
	refreshTokenString, _ := refreshToken.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))

	return c.JSON(fiber.Map{
		"statusCode": 200,
		"intCode":    "S01",
		"data": fiber.Map{
			"access_token":  tokenString,
			"refresh_token": refreshTokenString,
		},
	})
}

// RefreshToken genera un nuevo token de acceso usando un refresh token válido
  func RefreshToken(c *fiber.Ctx) error {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.BodyParser(&body); err != nil || body.RefreshToken == "" {
		return utils.ErrorResponse(c, 400, "R01", "Token de refresco requerido")
	}

	// Parsear y validar refresh token
	token, err := jwt.Parse(body.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return utils.ErrorResponse(c, 401, "R02", "Refresh token inválido")
	}

	claims := token.Claims.(jwt.MapClaims)
	id := int(claims["id"].(float64))
	rol := claims["rol"].(string) //

    permisos, err := utils.ObtenerPermisosPorRol(rol)
    if err != nil {
	return utils.ErrorResponse(c, 500, "F05", "Error al obtener permisos del rol")
    }

	// Nuevo Access Token
	newAccess := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"rol":      rol,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
		"permisos": permisos,
	})
	accessString, _ := newAccess.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return c.JSON(fiber.Map{
		"statusCode": 200,
		"intCode":    "S02",
		"data": fiber.Map{
	    "access_token": accessString,
		},
	})
}

func VerifyMFA(c *fiber.Ctx) error {
	var input struct {
		Correo string `json:"correo"`
		Code   string `json:"code"` // Código generado por el Authenticator
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	var usuario models.Usuario
	err := config.DB.QueryRow(context.Background(),
		"SELECT id_usuario, nombre, rol, mfa_secret FROM usuarios WHERE correo = $1",
		input.Correo).Scan(&usuario.IdUsuario, &usuario.Nombre, &usuario.Rol, &usuario.MfaSecret)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Usuario no encontrado"})
	}

	// Validar el código MFA
	valid := totp.Validate(input.Code, usuario.MfaSecret)
	if !valid {
		return c.Status(401).JSON(fiber.Map{"error": "Código MFA inválido"})
	}

	// Generar token JWT si MFA es válido
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     usuario.IdUsuario,
		"nombre": usuario.Nombre,
		"rol":    usuario.Rol,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo generar token"})
	}

	return c.JSON(fiber.Map{"access_token": tokenString})
}
