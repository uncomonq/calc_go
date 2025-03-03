package tests

import (
	"testing"

	"github.com/uncomonq/calc_go/internal/calculation"
)

func TestCompute(t *testing.T) {
	tests := []struct {
		op        string
		a, b      float64
		expected  float64
		shouldErr bool
	}{
		{"+", 10, 5, 15, false},
		{"-", 52, 10, 42, false},
		{"*", 6, 8, 48, false},
		{"/", 10, 2, 5, false},
		{"/", 10, 0, 0, true},
		{"^", 2, 3, 0, true},
	}
	for _, tc := range tests {
		result, err := calculation.Compute(tc.op, tc.a, tc.b)
		if tc.shouldErr && err == nil {
			t.Errorf("Ожидалась ошибка для операции %s", tc.op)
		}
		if !tc.shouldErr && result != tc.expected {
			t.Errorf("Compute(%s, %f, %f) = %f; ожидалось %f", tc.op, tc.a, tc.b, result, tc.expected)
		}
	}
}