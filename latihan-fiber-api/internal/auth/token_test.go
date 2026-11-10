package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestTerbitkanDanPeriksaToken(t *testing.T) {
	secret := "rahasia-uji-yang-panjang-minimal-32-byte"
	token, err := TerbitkanToken(42, "mahasiswa", secret, 10)
	if err != nil {
		t.Fatalf("TerbitkanToken() error = %v", err)
	}

	klaim, err := PeriksaToken(token, secret)
	if err != nil {
		t.Fatalf("PeriksaToken() error = %v", err)
	}
	if klaim.Subject != "42" || klaim.Peran != "mahasiswa" || klaim.ID == "" || klaim.ExpiresAt == nil {
		t.Fatalf("klaim token tidak lengkap: %+v", klaim)
	}
}

func TestPeriksaTokenMenolakRahasiaBerbeda(t *testing.T) {
	token, err := TerbitkanToken(42, "mahasiswa", "rahasia-uji-yang-panjang-minimal-32-byte", 10)
	if err != nil {
		t.Fatalf("TerbitkanToken() error = %v", err)
	}

	if _, err := PeriksaToken(token, "rahasia-uji-berbeda-yang-panjang-minimal-32-byte"); err == nil {
		t.Fatal("PeriksaToken() menerima token dengan rahasia yang berbeda")
	}
}

func TestPeriksaTokenMenolakTokenKedaluwarsa(t *testing.T) {
	secret := "rahasia-uji-yang-panjang-minimal-32-byte"
	sekarang := time.Now().Add(-2 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Klaim{
		Peran: "mahasiswa",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "42",
			ID:        "id-token-uji",
			IssuedAt:  jwt.NewNumericDate(sekarang),
			ExpiresAt: jwt.NewNumericDate(sekarang.Add(time.Minute)),
			Issuer:    penerbitToken,
		},
	})
	nilaiToken, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("SignedString() error = %v", err)
	}

	if _, err := PeriksaToken(nilaiToken, secret); err == nil {
		t.Fatal("PeriksaToken() menerima token yang sudah kedaluwarsa")
	}
}

func TestTerbitkanTokenMenolakRahasiaPendek(t *testing.T) {
	if _, err := TerbitkanToken(42, "mahasiswa", "rahasia-pendek", 10); err == nil {
		t.Fatal("TerbitkanToken() menerima rahasia yang kurang dari 32 byte")
	}
}
