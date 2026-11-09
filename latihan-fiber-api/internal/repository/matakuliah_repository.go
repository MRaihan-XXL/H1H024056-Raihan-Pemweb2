package repository

import (
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

func AmbilMatakuliah() ([]model.Matakuliah, error) {
	daftar := make([]model.Matakuliah, 0)
	err := config.DB.Order("kode ASC").Find(&daftar).Error
	return daftar, err
}

func AmbilMatakuliahPerID(id uint) (model.Matakuliah, error) {
	var matakuliah model.Matakuliah
	err := config.DB.First(&matakuliah, id).Error
	return matakuliah, err
}

func SimpanMatakuliah(matakuliah *model.Matakuliah) error {
	return config.DB.Create(matakuliah).Error
}

func PerbaruiMatakuliah(id uint, perubahan map[string]any) error {
	return config.DB.Model(&model.Matakuliah{}).Where("id = ?", id).Updates(perubahan).Error
}

func HapusMatakuliah(matakuliah *model.Matakuliah) error {
	return config.DB.Delete(matakuliah).Error
}
