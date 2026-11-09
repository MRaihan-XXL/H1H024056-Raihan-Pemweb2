package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func HubungkanBasisData(konf Konfigurasi) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		konf.DbUser,
		konf.DbSandi,
		konf.DbHost,
		konf.DbPort,
		konf.DbNama,
	)

	koneksi, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return fmt.Errorf("membuka koneksi basis data: %w", err)
	}

	sqlDB, err := koneksi.DB()
	if err != nil {
		return fmt.Errorf("mengakses koneksi basis data: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("memeriksa koneksi basis data: %w", err)
	}

	DB = koneksi
	return nil
}
