package handler

import (
	"errors"
	"net/mail"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"
)

type PermintaanMahasiswa struct {
	ProgramStudiID uint    `json:"program_studi_id"`
	Nim            string  `json:"nim"`
	Nama           string  `json:"nama"`
	Email          string  `json:"email"`
	Angkatan       int     `json:"angkatan"`
	Ipk            float64 `json:"ipk"`
	Aktif          *bool   `json:"aktif"`
}

type PembaruanMahasiswa struct {
	ProgramStudiID *uint    `json:"program_studi_id"`
	Nim            *string  `json:"nim"`
	Nama           *string  `json:"nama"`
	Email          *string  `json:"email"`
	Angkatan       *int     `json:"angkatan"`
	Ipk            *float64 `json:"ipk"`
	Aktif          *bool    `json:"aktif"`
}

func DaftarMahasiswa(c fiber.Ctx) error {
	halaman, err := angkaQuery(c, "halaman", 1)
	if err != nil || halaman < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "Parameter halaman harus bilangan positif")
	}
	perHalaman, err := angkaQuery(c, "per_halaman", 10)
	if err != nil || perHalaman < 1 || perHalaman > 100 {
		return fiber.NewError(fiber.StatusBadRequest, "Parameter per_halaman harus antara 1 dan 100")
	}
	angkatan, err := angkaQuery(c, "angkatan", 0)
	if err != nil || angkatan < 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Parameter angkatan harus berupa bilangan positif")
	}

	filter := repository.FilterMahasiswa{
		Cari:       c.Query("cari"),
		Angkatan:   angkatan,
		Halaman:    halaman,
		PerHalaman: perHalaman,
	}
	daftar, total, err := repository.AmbilMahasiswa(filter)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mahasiswa")
	}

	totalHalaman := int(total) / perHalaman
	if int(total)%perHalaman != 0 {
		totalHalaman++
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   daftar,
		"meta": fiber.Map{
			"halaman":       halaman,
			"per_halaman":   perHalaman,
			"total":         total,
			"total_halaman": totalHalaman,
		},
	})
}

func DetailMahasiswa(c fiber.Ctx) error {
	id, err := idDariParam(c)
	if err != nil {
		return err
	}

	mahasiswa, err := repository.AmbilMahasiswaPerID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mahasiswa")
	}

	return c.JSON(fiber.Map{"sukses": true, "data": mahasiswa})
}

func CariBerdasarkanNim(c fiber.Ctx) error {
	mahasiswa, err := repository.AmbilMahasiswaPerNIM(c.Params("nim"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mahasiswa")
	}
	return c.JSON(fiber.Map{"sukses": true, "data": mahasiswa})
}

func BuatMahasiswa(c fiber.Ctx) error {
	var permintaan PermintaanMahasiswa
	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format body tidak valid")
	}
	if err := validasiMahasiswa(permintaan); err != nil {
		return err
	}

	aktif := true
	if permintaan.Aktif != nil {
		aktif = *permintaan.Aktif
	}
	mahasiswa := model.Mahasiswa{
		ProgramStudiID: permintaan.ProgramStudiID,
		Nim:            strings.TrimSpace(permintaan.Nim),
		Nama:           strings.TrimSpace(permintaan.Nama),
		Email:          strings.TrimSpace(permintaan.Email),
		Angkatan:       permintaan.Angkatan,
		Ipk:            permintaan.Ipk,
		Aktif:          aktif,
	}
	if err := repository.SimpanMahasiswa(&mahasiswa); err != nil {
		return galatPenyimpanan(err, "Gagal menyimpan data mahasiswa")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil dibuat",
		"data":   mahasiswa,
	})
}

func PerbaruiMahasiswa(c fiber.Ctx) error {
	id, err := idDariParam(c)
	if err != nil {
		return err
	}
	if _, err := repository.AmbilMahasiswaPerID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mahasiswa")
	}

	var permintaan PembaruanMahasiswa
	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format body tidak valid")
	}
	if err := validasiPembaruanMahasiswa(permintaan); err != nil {
		return err
	}

	perubahan := make(map[string]any)
	if permintaan.ProgramStudiID != nil {
		perubahan["program_studi_id"] = *permintaan.ProgramStudiID
	}
	if permintaan.Nim != nil {
		perubahan["nim"] = strings.TrimSpace(*permintaan.Nim)
	}
	if permintaan.Nama != nil {
		perubahan["nama"] = strings.TrimSpace(*permintaan.Nama)
	}
	if permintaan.Email != nil {
		perubahan["email"] = strings.TrimSpace(*permintaan.Email)
	}
	if permintaan.Angkatan != nil {
		perubahan["angkatan"] = *permintaan.Angkatan
	}
	if permintaan.Ipk != nil {
		perubahan["ipk"] = *permintaan.Ipk
	}
	if permintaan.Aktif != nil {
		perubahan["aktif"] = *permintaan.Aktif
	}
	if len(perubahan) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Tidak ada field yang dapat diperbarui")
	}

	if err := repository.PerbaruiMahasiswa(id, perubahan); err != nil {
		return galatPenyimpanan(err, "Gagal memperbarui data mahasiswa")
	}
	mahasiswa, err := repository.AmbilMahasiswaPerID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mahasiswa terbaru")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil diperbarui",
		"data":   mahasiswa,
	})
}

func HapusMahasiswa(c fiber.Ctx) error {
	id, err := idDariParam(c)
	if err != nil {
		return err
	}
	mahasiswa, err := repository.AmbilMahasiswaPerID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mahasiswa")
	}
	if err := repository.HapusMahasiswa(&mahasiswa); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal menghapus data mahasiswa")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil dihapus",
	})
}

func PulihkanMahasiswa(c fiber.Ctx) error {
	id, err := idDariParam(c)
	if err != nil {
		return err
	}
	if err := repository.PulihkanMahasiswa(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Data mahasiswa terhapus tidak ditemukan")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal memulihkan data mahasiswa")
	}

	mahasiswa, err := repository.AmbilMahasiswaPerID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mahasiswa yang dipulihkan")
	}
	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil dipulihkan",
		"data":   mahasiswa,
	})
}

func validasiMahasiswa(permintaan PermintaanMahasiswa) error {
	if permintaan.ProgramStudiID == 0 ||
		strings.TrimSpace(permintaan.Nim) == "" ||
		strings.TrimSpace(permintaan.Nama) == "" ||
		strings.TrimSpace(permintaan.Email) == "" ||
		permintaan.Angkatan <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Program studi, NIM, nama, email, dan angkatan wajib diisi")
	}
	if permintaan.Email != "" {
		alamat, err := mail.ParseAddress(strings.TrimSpace(permintaan.Email))
		if err != nil || alamat.Address != strings.TrimSpace(permintaan.Email) {
			return fiber.NewError(fiber.StatusBadRequest, "Format email tidak valid")
		}
	}
	if permintaan.Ipk < 0 || permintaan.Ipk > 4 {
		return fiber.NewError(fiber.StatusBadRequest, "IPK harus berada pada rentang 0 sampai 4")
	}
	return nil
}

func validasiPembaruanMahasiswa(permintaan PembaruanMahasiswa) error {
	if permintaan.ProgramStudiID != nil && *permintaan.ProgramStudiID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Program studi harus berupa ID positif")
	}
	if permintaan.Nim != nil && strings.TrimSpace(*permintaan.Nim) == "" {
		return fiber.NewError(fiber.StatusBadRequest, "NIM tidak boleh kosong")
	}
	if permintaan.Nama != nil && strings.TrimSpace(*permintaan.Nama) == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Nama tidak boleh kosong")
	}
	if permintaan.Email != nil {
		alamat, err := mail.ParseAddress(strings.TrimSpace(*permintaan.Email))
		if err != nil || alamat.Address != strings.TrimSpace(*permintaan.Email) {
			return fiber.NewError(fiber.StatusBadRequest, "Format email tidak valid")
		}
	}
	if permintaan.Angkatan != nil && *permintaan.Angkatan <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Angkatan harus berupa bilangan positif")
	}
	if permintaan.Ipk != nil && (*permintaan.Ipk < 0 || *permintaan.Ipk > 4) {
		return fiber.NewError(fiber.StatusBadRequest, "IPK harus berada pada rentang 0 sampai 4")
	}
	return nil
}

func angkaQuery(c fiber.Ctx, nama string, bawaan int) (int, error) {
	nilai := c.Query(nama)
	if nilai == "" {
		return bawaan, nil
	}
	return strconv.Atoi(nilai)
}
