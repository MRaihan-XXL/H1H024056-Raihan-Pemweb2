package middleware

import (
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/auth"
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/repository"
)

func Autentikasi() fiber.Handler {
	konf := config.Muat()

	return func(c fiber.Ctx) error {
		bagian := strings.Fields(c.Get("Authorization"))
		if len(bagian) != 2 || !strings.EqualFold(bagian[0], "Bearer") {
			return fiber.NewError(fiber.StatusUnauthorized, "Header Authorization tidak valid")
		}

		klaim, err := auth.PeriksaToken(bagian[1], konf.JwtSecret)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Token tidak valid atau sudah kedaluwarsa")
		}

		diblokir, err := repository.TokenDiblokir(klaim.ID)
		if err != nil {
			log.Printf("Gagal memeriksa daftar hitam token: %v", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Gagal memeriksa status token")
		}
		if diblokir {
			return fiber.NewError(fiber.StatusUnauthorized, "Token sudah dicabut")
		}

		idPengguna, err := strconv.ParseUint(klaim.Subject, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Identitas pada token tidak valid")
		}

		c.Locals("id_pengguna", uint(idPengguna))
		c.Locals("klaim", klaim)
		c.Locals("peran", klaim.Peran)
		return c.Next()
	}
}
