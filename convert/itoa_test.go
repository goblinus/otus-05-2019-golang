package convert

import (
	"testing"
)

func TestItoa(t *testing.T) {
	if result := Itoa(0); result != "0" {
		t.Errorf("Должно быть '0', получили %s", result)
	}

	if result := Itoa(100); result != "100" {
		t.Errorf("Должно быть '0', получили %s", result)
	}

	if result := Itoa(-100); result != "-100" {
		t.Errorf("Должно быть '0', получили %s", result)
	}
}
