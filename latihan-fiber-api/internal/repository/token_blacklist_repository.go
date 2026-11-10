package repository

import (
	"time"

	"gorm.io/gorm/clause"

	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

func TokenDiblokir(jti string) (bool, error) {
	var jumlah int64
	err := config.DB.Model(&model.TokenDiblokir{}).
		Where("jti = ? AND kedaluwarsa > ?", jti, time.Now()).
		Count(&jumlah).Error
	if err != nil {
		return false, err
	}

	return jumlah > 0, nil
}

func CabutToken(jti string, kedaluwarsa time.Time) (bool, error) {
	if err := config.DB.Where("kedaluwarsa <= ?", time.Now()).
		Delete(&model.TokenDiblokir{}).Error; err != nil {
		return false, err
	}

	hasil := config.DB.Clauses(clause.OnConflict{DoNothing: true}).
		Create(&model.TokenDiblokir{
			JTI:         jti,
			Kedaluwarsa: kedaluwarsa,
		})
	return hasil.RowsAffected > 0, hasil.Error
}
