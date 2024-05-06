package calculator_test

import (
	"testing"

	calculator "github.com/YounesOuterbah/goDS.git"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 6
	got := calculator.Add(3, 3)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 3
	got := calculator.Subtract(6, 3)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 8
	got := calculator.Multiply(2, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
