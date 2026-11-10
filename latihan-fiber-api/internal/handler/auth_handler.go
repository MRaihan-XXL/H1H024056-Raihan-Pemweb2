package handler

import (
	"errors"
	"log"
	"strings"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"latihan-fiber-api/internal/auth"
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"
	"latihan-fiber-api/internal/validasi"
)

type PermintaanRegistrasi struct {
	Nama      string `json:"nama" validate:"required,min=3,max=100"`
	Email     string `json:"email" validate:"required,email,max=100"`
	KataSandi string `json:"kata_sandi" validate:"required,min=8"`
}

type PermintaanLogin struct {
	Email     string `json:"email" validate:"required,email"`
	KataSandi string `json:"kata_sandi" validate:"required"`
}

func Registrasi(c fiber.Ctx) error {
	var permintaan PermintaanRegistrasi
	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format body tidak valid")
	}

	permintaan.Nama = strings.TrimSpace(permintaan.Nama)
	permintaan.Email = strings.TrimSpace(permintaan.Email)
	galat, err := validasi.Periksa(permintaan)
	if err != nil {
		log.Printf("Gagal memvalidasi permintaan registrasi: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal memvalidasi data")
	}
	if len(galat) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data yang dikirim tidak valid",
			"galat":  galat,
		})
	}
	if len([]byte(permintaan.KataSandi)) > 72 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data yang dikirim tidak valid",
			"galat": []validasi.PesanGalat{{
				Field: "KataSandi",
				Pesan: "Kata sandi tidak boleh lebih dari 72 byte",
			}},
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(permintaan.KataSandi), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Gagal membuat hash kata sandi: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal memproses kata sandi")
	}

	pengguna := model.Pengguna{
		Nama:      permintaan.Nama,
		Email:     permintaan.Email,
		KataSandi: string(hash),
		Peran:     "mahasiswa",
	}
	if err := config.DB.Create(&pengguna).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return fiber.NewError(fiber.StatusConflict, "Email sudah terdaftar")
		}
		log.Printf("Gagal menyimpan pengguna baru: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal menyimpan data pengguna")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Registrasi berhasil",
		"data":   pengguna,
	})
}

func Login(c fiber.Ctx) error {
	var permintaan PermintaanLogin
	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format body tidak valid")
	}

	permintaan.Email = strings.TrimSpace(permintaan.Email)
	galat, err := validasi.Periksa(permintaan)
	if err != nil {
		log.Printf("Gagal memvalidasi permintaan login: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal memvalidasi data")
	}
	if len(galat) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data yang dikirim tidak valid",
			"galat":  galat,
		})
	}

	var pengguna model.Pengguna
	err = config.DB.Where("email = ?", permintaan.Email).First(&pengguna).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusUnauthorized, "Email atau kata sandi tidak sesuai")
	}
	if err != nil {
		log.Printf("Gagal mencari pengguna saat login: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal memproses login")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(pengguna.KataSandi), []byte(permintaan.KataSandi)); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Email atau kata sandi tidak sesuai")
	}

	konf := config.Muat()
	token, err := auth.TerbitkanToken(pengguna.ID, pengguna.Peran, konf.JwtSecret, konf.JwtTtlMenit)
	if err != nil {
		log.Printf("Gagal menerbitkan token: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal menerbitkan token")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Login berhasil",
		"data": fiber.Map{
			"pengguna": pengguna,
			"token":    token,
			"tipe":     "Bearer",
		},
	})
}

func Profil(c fiber.Ctx) error {
	idPengguna, cocok := c.Locals("id_pengguna").(uint)
	if !cocok {
		return fiber.NewError(fiber.StatusUnauthorized, "Identitas pengguna tidak valid")
	}

	var pengguna model.Pengguna
	if err := config.DB.First(&pengguna, idPengguna).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Pengguna tidak ditemukan")
		}
		log.Printf("Gagal mengambil profil pengguna: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil profil pengguna")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   pengguna,
	})
}

func Segarkan(c fiber.Ctx) error {
	klaim, cocok := c.Locals("klaim").(*auth.Klaim)
	if !cocok {
		return fiber.NewError(fiber.StatusUnauthorized, "Klaim token tidak tersedia")
	}

	idPengguna, cocok := c.Locals("id_pengguna").(uint)
	if !cocok {
		return fiber.NewError(fiber.StatusUnauthorized, "Identitas pengguna tidak valid")
	}

	var pengguna model.Pengguna
	if err := config.DB.First(&pengguna, idPengguna).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusUnauthorized, "Pengguna tidak ditemukan")
		}
		log.Printf("Gagal mengambil pengguna untuk refresh token: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal menerbitkan token baru")
	}

	konf := config.Muat()
	tokenBaru, err := auth.TerbitkanToken(pengguna.ID, pengguna.Peran, konf.JwtSecret, konf.JwtTtlMenit)
	if err != nil {
		log.Printf("Gagal menerbitkan token pengganti: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal menerbitkan token baru")
	}
	dicabut, err := repository.CabutToken(klaim.ID, klaim.ExpiresAt.Time)
	if err != nil {
		log.Printf("Gagal mencabut token setelah refresh: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mencabut token lama")
	}
	if !dicabut {
		return fiber.NewError(fiber.StatusUnauthorized, "Token sudah dicabut")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Token berhasil diperbarui",
		"data": fiber.Map{
			"token": tokenBaru,
			"tipe":  "Bearer",
		},
	})
}

func Logout(c fiber.Ctx) error {
	klaim, cocok := c.Locals("klaim").(*auth.Klaim)
	if !cocok {
		return fiber.NewError(fiber.StatusUnauthorized, "Klaim token tidak tersedia")
	}

	dicabut, err := repository.CabutToken(klaim.ID, klaim.ExpiresAt.Time)
	if err != nil {
		log.Printf("Gagal mencabut token saat logout: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mencabut token")
	}
	if !dicabut {
		return fiber.NewError(fiber.StatusUnauthorized, "Token sudah dicabut")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Logout berhasil",
	})
}
