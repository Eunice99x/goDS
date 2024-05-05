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
