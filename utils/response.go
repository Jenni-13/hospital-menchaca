package utils

import "github.com/gofiber/fiber/v2"

// Códigos internos
const (
	S01LoginExitoso       = "S01"
	S02RegistroExitoso    = "S02"
	F01DatosInvalidos     = "F01"
	F02UsuarioNoEncontrado = "F02"
	F03ContraseñaIncorrecta = "F03"
	W01AdvertenciaGeneral = "W01"
	A01AlertaGeneral       = "A01"
)

// SuccessResponse genera una respuesta exitosa con datos
func SuccessResponse(c *fiber.Ctx, intCode, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"intCode":    intCode,
		"message":    message,
		"data":       data,
	})
}

// ErrorResponse genera una respuesta de error sin datos
func ErrorResponse(c *fiber.Ctx, statusCode int, intCode, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"statusCode": statusCode,
		"intCode":    intCode,
		"message":    message,
	})
}
