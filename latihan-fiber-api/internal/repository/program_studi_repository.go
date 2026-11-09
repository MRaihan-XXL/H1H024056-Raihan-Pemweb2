package repository

import "latihan-fiber-api/internal/config"

type ProgramStudiDenganJumlahMahasiswa struct {
	ID              uint   `json:"id"`
	Kode            string `json:"kode"`
	Nama            string `json:"nama"`
	Jenjang         string `json:"jenjang"`
	JumlahMahasiswa int64  `json:"jumlah_mahasiswa"`
}

func AmbilProgramStudi() ([]ProgramStudiDenganJumlahMahasiswa, error) {
	daftar := make([]ProgramStudiDenganJumlahMahasiswa, 0)
	err := config.DB.Table("program_studis").
		Select("program_studis.id, program_studis.kode, program_studis.nama, program_studis.jenjang, COUNT(mahasiswas.id) AS jumlah_mahasiswa").
		Joins("LEFT JOIN mahasiswas ON mahasiswas.program_studi_id = program_studis.id AND mahasiswas.deleted_at IS NULL").
		Group("program_studis.id, program_studis.kode, program_studis.nama, program_studis.jenjang").
		Order("program_studis.nama ASC").
		Scan(&daftar).Error
	return daftar, err
}
