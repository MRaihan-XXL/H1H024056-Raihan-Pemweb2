package model

type Matakuliah struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Kode     string `gorm:"size:20;uniqueIndex;not null" json:"kode"`
	Nama     string `gorm:"size:100;not null" json:"nama"`
	SKS      int    `gorm:"not null" json:"sks"`
	Semester int    `gorm:"not null" json:"semester"`
}

func (Matakuliah) TableName() string {
	return "matakuliahs"
}
