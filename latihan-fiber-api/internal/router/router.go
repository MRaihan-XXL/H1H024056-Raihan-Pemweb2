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

	auth := api.Group("/auth")
	auth.Post("/registrasi", handler.Registrasi)
	auth.Post("/login", handler.Login)

	authTerproteksi := api.Group("/auth")
	authTerproteksi.Use(middleware.Autentikasi())
	authTerproteksi.Get("/profil", handler.Profil)
	authTerproteksi.Post("/segarkan", handler.Segarkan)
	authTerproteksi.Post("/logout", handler.Logout)

	mahasiswa := api.Group("/mahasiswa")
	mahasiswa.Use(middleware.KunciApi(kunciAPI))

	mahasiswa.Get("", handler.DaftarMahasiswa)
	mahasiswa.Post("", handler.BuatMahasiswa)
	mahasiswa.Get("/nim/:nim", handler.CariBerdasarkanNim)
	mahasiswa.Post("/:id/pulihkan", handler.PulihkanMahasiswa)
	mahasiswa.Get("/:id", handler.DetailMahasiswa)
	mahasiswa.Put("/:id", handler.PerbaruiMahasiswa)
	mahasiswa.Delete("/:id", handler.HapusMahasiswa)

	matakuliah := api.Group("/matakuliah")
	matakuliah.Get("", handler.DaftarMatakuliah)
	matakuliah.Post("", handler.BuatMatakuliah)
	matakuliah.Get("/:id", handler.DetailMatakuliah)
	matakuliah.Put("/:id", handler.PerbaruiMatakuliah)
	matakuliah.Delete("/:id", handler.HapusMatakuliah)

	api.Get("/program-studi", handler.DaftarProgramStudi)
}
