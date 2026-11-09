package model

import (
	"time"

	"gorm.io/gorm"
)

type ProgramStudi struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Kode      string    `gorm:"size:10;uniqueIndex;not null" json:"kode"`
	Nama      string    `gorm:"size:100;not null" json:"nama"`
	Jenjang   string    `gorm:"size:10;not null" json:"jenjang"`
	CreatedAt time.Time `json:"dibuat_pada"`
	UpdatedAt time.Time `json:"diubah_pada"`
}

func (ProgramStudi) TableName() string {
	return "program_studis"
}

type Mahasiswa struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	ProgramStudiID uint           `gorm:"not null;index" json:"program_studi_id"`
	ProgramStudi   *ProgramStudi  `gorm:"foreignKey:ProgramStudiID" json:"program_studi,omitempty"`
	Nim            string         `gorm:"size:20;uniqueIndex;not null" json:"nim"`
	Nama           string         `gorm:"size:100;not null" json:"nama"`
	Email          string         `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Angkatan       int            `gorm:"not null" json:"angkatan"`
	Ipk            float64        `gorm:"type:decimal(4,2);not null" json:"ipk"`
	Aktif          bool           `gorm:"not null" json:"aktif"`
	CreatedAt      time.Time      `json:"dibuat_pada"`
	UpdatedAt      time.Time      `json:"diubah_pada"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Mahasiswa) TableName() string {
	return "mahasiswas"
}
