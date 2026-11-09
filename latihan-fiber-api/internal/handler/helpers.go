package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func idDariParam(c fiber.Ctx) (uint, error) {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil || id == 0 {
		return 0, fiber.NewError(fiber.StatusBadRequest, "ID harus berupa bilangan positif")
	}
	return uint(id), nil
}

func galatPenyimpanan(err error, pesan string) error {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return fiber.NewError(fiber.StatusNotFound, "Data tidak ditemukan")
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return fiber.NewError(fiber.StatusConflict, "Kode, NIM, atau email sudah digunakan")
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Relasi data yang dipilih tidak valid")
	default:
		return fiber.NewError(fiber.StatusInternalServerError, pesan)
	}
}
