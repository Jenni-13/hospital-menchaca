package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Middleware que valida que el JWT sea válido
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		if tokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token faltante"})
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token inválido"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Claims inválidos"})
		}

		// ✅ Guardamos claims y también el id del usuario
		c.Locals("usuario", claims)

		if idFloat, ok := claims["id"].(float64); ok {
			c.Locals("user_id", int(idFloat)) // <- Aquí lo importante
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "ID no válido en token"})
		}
c.Locals("usuario", claims)
		return c.Next()
	}
}


// Middleware que valida si el usuario tiene un permiso específico
func PermisoRequerido(permiso string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("usuario").(jwt.MapClaims)

		// Verificamos que los permisos estén en el token
		permisosClaim, ok := claims["permisos"]
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Permisos no disponibles en el token"})
		}

		permisos, ok := permisosClaim.([]interface{})
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Permisos en formato incorrecto"})
		}

		// Buscar el permiso requerido
		for _, p := range permisos {
			if p.(string) == permiso {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "No tienes el permiso necesario: " + permiso})
	}
}


// Middleware que valida si el rol está permitido
func RolesAllowed(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("usuario").(jwt.MapClaims)
		userRole := claims["rol"].(string)

		for _, role := range roles {
			if role == userRole {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "No tienes permiso para acceder a esta ruta"})
	}
}


func RoleRequired(role string) fiber.Handler {
	return RolesAllowed(role)
}
