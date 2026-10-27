package router

import (
	"github.com/gofiber/fiber/v3"
	"latihan-fiber-api/internal/handler"
	"latihan-fiber-api/internal/middleware"
)

func Daftarkan(app *fiber.App, kunciAPI string) {
	api := app.Group("/api")

	api.Get("/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"sukses": true,
			"pesan":  "API Fiber aktif",
		})
	})

	mahasiswa := api.Group("/mahasiswa")
	mahasiswa.Use(middleware.KunciApi(kunciAPI))

	mahasiswa.Get("", handler.DaftarMahasiswa)
	mahasiswa.Get("/nim/:nim", handler.CariBerdasarkanNim)
}
