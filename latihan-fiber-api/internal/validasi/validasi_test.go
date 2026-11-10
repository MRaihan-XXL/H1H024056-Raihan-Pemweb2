package validasi

import "testing"

func TestPeriksaNIM(t *testing.T) {
	tests := []struct {
		nim   string
		valid bool
	}{
		{nim: "H12345678", valid: true},
		{nim: "12345678", valid: false},
		{nim: "h12345678", valid: false},
		{nim: "H12345A78", valid: false},
		{nim: "H", valid: false},
		{nim: "", valid: false},
	}

	for _, test := range tests {
		t.Run(test.nim, func(t *testing.T) {
			galat, err := PeriksaNIM(test.nim)
			if err != nil {
				t.Fatalf("PeriksaNIM() error = %v", err)
			}
			if valid := len(galat) == 0; valid != test.valid {
				t.Errorf("PeriksaNIM(%q) valid = %v, want %v", test.nim, valid, test.valid)
			}
		})
	}
}
