package repository

import (
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"

	"gorm.io/gorm"
)

type FilterMahasiswa struct {
	Cari       string
	Angkatan   int
	Halaman    int
	PerHalaman int
}

func AmbilMahasiswa(filter FilterMahasiswa) ([]model.Mahasiswa, int64, error) {
	daftar := make([]model.Mahasiswa, 0)
	var total int64

	kueri := config.DB.Model(&model.Mahasiswa{})
	if filter.Cari != "" {
		pola := "%" + filter.Cari + "%"
		kueri = kueri.Where("nama LIKE ? OR nim LIKE ?", pola, pola)
	}
	if filter.Angkatan > 0 {
		kueri = kueri.Where("angkatan = ?", filter.Angkatan)
	}

	if err := kueri.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	lompat := (filter.Halaman - 1) * filter.PerHalaman
	if err := kueri.Preload("ProgramStudi").
		Order("nama ASC").
		Limit(filter.PerHalaman).
		Offset(lompat).
		Find(&daftar).Error; err != nil {
		return nil, 0, err
	}

	return daftar, total, nil
}

func AmbilMahasiswaPerID(id uint) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa
	err := config.DB.Preload("ProgramStudi").First(&mahasiswa, id).Error
	return mahasiswa, err
}

func AmbilMahasiswaPerNIM(nim string) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa
	err := config.DB.Preload("ProgramStudi").Where("nim = ?", nim).First(&mahasiswa).Error
	return mahasiswa, err
}

func SimpanMahasiswa(mahasiswa *model.Mahasiswa) error {
	return config.DB.Create(mahasiswa).Error
}

func PerbaruiMahasiswa(id uint, perubahan map[string]any) error {
	return config.DB.Model(&model.Mahasiswa{}).Where("id = ?", id).Updates(perubahan).Error
}

func HapusMahasiswa(mahasiswa *model.Mahasiswa) error {
	return config.DB.Delete(mahasiswa).Error
}

func PulihkanMahasiswa(id uint) error {
	hasil := config.DB.Unscoped().
		Model(&model.Mahasiswa{}).
		Where("id = ? AND deleted_at IS NOT NULL", id).
		Update("deleted_at", nil)
	if hasil.Error != nil {
		return hasil.Error
	}
	if hasil.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
