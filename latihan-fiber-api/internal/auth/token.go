package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const penerbitToken = "pemweb2-fiber-api"
const panjangMinimalRahasia = 32

type Klaim struct {
	Peran string `json:"peran"`
	jwt.RegisteredClaims
}

func TerbitkanToken(idPengguna uint, peran string, rahasia string, ttlMenit int) (string, error) {
	if idPengguna == 0 || peran == "" {
		return "", errors.New("identitas pengguna dan peran wajib diisi")
	}
	if len([]byte(rahasia)) < panjangMinimalRahasia {
		return "", errors.New("JWT_SECRET harus memiliki panjang minimal 32 byte")
	}
	if ttlMenit < 1 {
		return "", errors.New("masa berlaku token harus lebih dari nol")
	}

	idToken, err := buatIDToken()
	if err != nil {
		return "", fmt.Errorf("membuat ID token: %w", err)
	}

	sekarang := time.Now()
	klaim := Klaim{
		Peran: peran,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(uint64(idPengguna), 10),
			ID:        idToken,
			IssuedAt:  jwt.NewNumericDate(sekarang),
			ExpiresAt: jwt.NewNumericDate(sekarang.Add(time.Duration(ttlMenit) * time.Minute)),
			Issuer:    penerbitToken,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, klaim)
	return token.SignedString([]byte(rahasia))
}

func PeriksaToken(nilaiToken string, rahasia string) (*Klaim, error) {
	if len([]byte(rahasia)) < panjangMinimalRahasia {
		return nil, errors.New("JWT_SECRET harus memiliki panjang minimal 32 byte")
	}

	token, err := jwt.ParseWithClaims(
		nilaiToken,
		&Klaim{},
		func(token *jwt.Token) (any, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, errors.New("metode penandatanganan tidak didukung")
			}
			return []byte(rahasia), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuer(penerbitToken),
		jwt.WithIssuedAt(),
	)
	if err != nil {
		return nil, err
	}

	klaim, cocok := token.Claims.(*Klaim)
	if !cocok || !token.Valid {
		return nil, errors.New("token tidak valid")
	}
	if klaim.Subject == "" || klaim.ID == "" || klaim.Peran == "" || klaim.ExpiresAt == nil {
		return nil, errors.New("klaim token tidak lengkap")
	}

	return klaim, nil
}

func buatIDToken() (string, error) {
	var nilai [32]byte
	if _, err := rand.Read(nilai[:]); err != nil {
		return "", err
	}
	return hex.EncodeToString(nilai[:]), nil
}
