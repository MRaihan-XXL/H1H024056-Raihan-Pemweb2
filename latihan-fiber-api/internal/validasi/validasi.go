package validasi

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var polaNIM = regexp.MustCompile(`^H[0-9]+$`)

var pemeriksa, galatInisialisasi = newPemeriksa()

type PesanGalat struct {
	Field string `json:"field"`
	Pesan string `json:"pesan"`
}

func newPemeriksa() (*validator.Validate, error) {
	pemeriksa := validator.New()
	if err := pemeriksa.RegisterValidation("nim", func(field validator.FieldLevel) bool {
		return polaNIM.MatchString(field.Field().String())
	}); err != nil {
		return nil, fmt.Errorf("mendaftarkan aturan validasi NIM: %w", err)
	}

	return pemeriksa, nil
}

func Periksa(data any) ([]PesanGalat, error) {
	if galatInisialisasi != nil {
		return nil, galatInisialisasi
	}

	err := pemeriksa.Struct(data)
	if err == nil {
		return nil, nil
	}

	var galatValidasi validator.ValidationErrors
	if !errors.As(err, &galatValidasi) {
		return nil, fmt.Errorf("memvalidasi data: %w", err)
	}

	daftar := make([]PesanGalat, 0, len(galatValidasi))
	for _, item := range galatValidasi {
		daftar = append(daftar, PesanGalat{
			Field: item.Field(),
			Pesan: susunPesan(item),
		})
	}

	return daftar, nil
}

func PeriksaNIM(nim string) ([]PesanGalat, error) {
	return Periksa(struct {
		NIM string `validate:"required,nim"`
	}{NIM: nim})
}

func susunPesan(item validator.FieldError) string {
	switch item.Tag() {
	case "required":
		return "Field ini wajib diisi"
	case "email":
		return "Format email tidak valid"
	case "min":
		return "Nilai minimal adalah " + item.Param()
	case "max":
		return "Nilai maksimal adalah " + item.Param()
	case "nim":
		return "NIM harus diawali huruf H dan diikuti angka"
	default:
		return "Nilai tidak memenuhi aturan " + item.Tag()
	}
}
