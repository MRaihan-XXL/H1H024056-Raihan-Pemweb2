package handler

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"
)

type PermintaanMatakuliah struct {
	Kode     string `json:"kode"`
	Nama     string `json:"nama"`
	SKS      int    `json:"sks"`
	Semester int    `json:"semester"`
}

func DaftarMatakuliah(c fiber.Ctx) error {
	daftar, err := repository.AmbilMatakuliah()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mata kuliah")
	}
	return c.JSON(fiber.Map{"sukses": true, "data": daftar})
}

func DetailMatakuliah(c fiber.Ctx) error {
	id, err := idDariParam(c)
	if err != nil {
		return err
	}
	matakuliah, err := repository.AmbilMatakuliahPerID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Mata kuliah tidak ditemukan")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mata kuliah")
	}
	return c.JSON(fiber.Map{"sukses": true, "data": matakuliah})
}

func BuatMatakuliah(c fiber.Ctx) error {
	var permintaan PermintaanMatakuliah
	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format body tidak valid")
	}
	if err := validasiMatakuliah(permintaan, true); err != nil {
		return err
	}

	matakuliah := model.Matakuliah{
		Kode:     strings.TrimSpace(permintaan.Kode),
		Nama:     strings.TrimSpace(permintaan.Nama),
		SKS:      permintaan.SKS,
		Semester: permintaan.Semester,
	}
	if err := repository.SimpanMatakuliah(&matakuliah); err != nil {
		return galatPenyimpanan(err, "Gagal menyimpan data mata kuliah")
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mata kuliah berhasil dibuat",
		"data":   matakuliah,
	})
}

func PerbaruiMatakuliah(c fiber.Ctx) error {
	id, err := idDariParam(c)
	if err != nil {
		return err
	}
	if _, err := repository.AmbilMatakuliahPerID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Mata kuliah tidak ditemukan")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mata kuliah")
	}

	var permintaan PermintaanMatakuliah
	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format body tidak valid")
	}
	if err := validasiMatakuliah(permintaan, false); err != nil {
		return err
	}

	perubahan := make(map[string]any)
	if strings.TrimSpace(permintaan.Kode) != "" {
		perubahan["kode"] = strings.TrimSpace(permintaan.Kode)
	}
	if strings.TrimSpace(permintaan.Nama) != "" {
		perubahan["nama"] = strings.TrimSpace(permintaan.Nama)
	}
	if permintaan.SKS > 0 {
		perubahan["sks"] = permintaan.SKS
	}
	if permintaan.Semester > 0 {
		perubahan["semester"] = permintaan.Semester
	}
	if len(perubahan) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Tidak ada field yang dapat diperbarui")
	}

	if err := repository.PerbaruiMatakuliah(id, perubahan); err != nil {
		return galatPenyimpanan(err, "Gagal memperbarui data mata kuliah")
	}
	matakuliah, err := repository.AmbilMatakuliahPerID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mata kuliah terbaru")
	}
	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mata kuliah berhasil diperbarui",
		"data":   matakuliah,
	})
}

func HapusMatakuliah(c fiber.Ctx) error {
	id, err := idDariParam(c)
	if err != nil {
		return err
	}
	matakuliah, err := repository.AmbilMatakuliahPerID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Mata kuliah tidak ditemukan")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mata kuliah")
	}
	if err := repository.HapusMatakuliah(&matakuliah); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal menghapus data mata kuliah")
	}
	return c.JSON(fiber.Map{"sukses": true, "pesan": "Mata kuliah berhasil dihapus"})
}

func validasiMatakuliah(permintaan PermintaanMatakuliah, wajib bool) error {
	if wajib && (strings.TrimSpace(permintaan.Kode) == "" ||
		strings.TrimSpace(permintaan.Nama) == "" ||
		permintaan.SKS <= 0 ||
		permintaan.Semester <= 0) {
		return fiber.NewError(fiber.StatusBadRequest, "Kode, nama, SKS, dan semester wajib diisi")
	}
	if permintaan.SKS < 0 || permintaan.Semester < 0 {
		return fiber.NewError(fiber.StatusBadRequest, "SKS dan semester tidak boleh negatif")
	}
	return nil
}
