package config

import (
	"fmt"

	"latihan-fiber-api/internal/model"
)

func Migrasi() error {
	if err := DB.AutoMigrate(
		&model.ProgramStudi{},
		&model.Mahasiswa{},
		&model.Matakuliah{},
	); err != nil {
		return fmt.Errorf("menjalankan migrasi: %w", err)
	}

	return nil
}

func DataAwal() error {
	var jumlah int64
	if err := DB.Model(&model.ProgramStudi{}).Count(&jumlah).Error; err != nil {
		return fmt.Errorf("memeriksa data program studi: %w", err)
	}

	if jumlah > 0 {
		return nil
	}

	daftar := []model.ProgramStudi{
		{Kode: "TK", Nama: "Teknik Komputer", Jenjang: "S1"},
		{Kode: "IF", Nama: "Informatika", Jenjang: "S1"},
		{Kode: "TE", Nama: "Teknik Elektro", Jenjang: "S1"},
	}

	if err := DB.Create(&daftar).Error; err != nil {
		return fmt.Errorf("membuat data awal program studi: %w", err)
	}

	return nil
}
