package middleware

import "github.com/gofiber/fiber/v3"

func KunciApi(kunci string) fiber.Handler {
	return func(c fiber.Ctx) error {
		if kunci == "" || c.Get("X-Api-Key") != kunci {
			return fiber.NewError(fiber.StatusUnauthorized, "X-Api-Key tidak sesuai")
		}

		return c.Next()
	}
}
