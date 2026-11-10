package model

import "time"

type Pengguna struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"size:100;not null" json:"nama"`
	Email     string    `gorm:"size:100;uniqueIndex;not null" json:"email"`
	KataSandi string    `gorm:"size:255;not null" json:"-"`
	Peran     string    `gorm:"size:20;not null;default:mahasiswa" json:"peran"`
	CreatedAt time.Time `json:"dibuat_pada"`
	UpdatedAt time.Time `json:"diubah_pada"`
}

func (Pengguna) TableName() string {
	return "penggunas"
}

type TokenDiblokir struct {
	JTI         string    `gorm:"column:jti;primaryKey;size:64" json:"-"`
	Kedaluwarsa time.Time `gorm:"column:kedaluwarsa;not null;index" json:"-"`
}

func (TokenDiblokir) TableName() string {
	return "token_blacklists"
}
