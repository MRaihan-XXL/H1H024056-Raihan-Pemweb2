package handler

import (
	"sort"

	"github.com/gofiber/fiber/v3"
	"latihan-fiber-api/internal/model"
)

func DaftarMahasiswa(c fiber.Ctx) error {
	urut := c.Query("urut", "")

	if urut != "" && urut != "nama" && urut != "ipk" {
		return fiber.NewError(fiber.StatusBadRequest, "Parameter urut harus bernilai nama atau ipk")
	}

	hasil := make([]model.Mahasiswa, len(model.Daftar))
	copy(hasil, model.Daftar)

	if urut == "nama" {
		sort.SliceStable(hasil, func(i, j int) bool {
			return hasil[i].Nama < hasil[j].Nama
		})
	}

	if urut == "ipk" {
		sort.SliceStable(hasil, func(i, j int) bool {
			return hasil[i].Ipk > hasil[j].Ipk
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"urut":   urut,
		"data":   hasil,
	})
}

func CariBerdasarkanNim(c fiber.Ctx) error {
	nim := c.Params("nim")

	for i := 0; i < len(model.Daftar); i++ {
		if model.Daftar[i].Nim == nim {
			return c.JSON(fiber.Map{
				"sukses": true,
				"data":   model.Daftar[i],
			})
		}
	}

	return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
}
