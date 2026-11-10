package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestRuteAuthTerlindungiMemerlukanToken(t *testing.T) {
	t.Setenv("JWT_SECRET", "rahasia-uji-router")

	app := fiber.New()
	Daftarkan(app, "")

	tests := []struct {
		method string
		path   string
	}{
		{method: http.MethodGet, path: "/api/auth/profil"},
		{method: http.MethodPost, path: "/api/auth/segarkan"},
		{method: http.MethodPost, path: "/api/auth/logout"},
	}

	for _, test := range tests {
		t.Run(test.method+" "+test.path, func(t *testing.T) {
			request := httptest.NewRequest(test.method, test.path, nil)
			response, err := app.Test(request)
			if err != nil {
				t.Fatalf("app.Test() error = %v", err)
			}
			defer response.Body.Close()
			if response.StatusCode != fiber.StatusUnauthorized {
				t.Errorf("status = %d, want %d", response.StatusCode, fiber.StatusUnauthorized)
			}
		})
	}
}

func TestRuteRegistrasiTerbukaMenolakKataSandiTerlaluPanjang(t *testing.T) {
	t.Setenv("JWT_SECRET", "rahasia-uji-router-yang-panjang-minimal-32-byte")

	app := fiber.New()
	Daftarkan(app, "")

	body := fmt.Sprintf(
		`{"nama":"Test User","email":"test@example.com","kata_sandi":"%s"}`,
		strings.Repeat("a", 73),
	)
	request := httptest.NewRequest(http.MethodPost, "/api/auth/registrasi", strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != fiber.StatusUnprocessableEntity {
		t.Errorf("status = %d, want %d", response.StatusCode, fiber.StatusUnprocessableEntity)
	}
}
