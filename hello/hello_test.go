package hello

import "testing"

func TestHello(t *testing.T) {
	if got := Hello(); got != "Hello" {
		t.Errorf("got: %s", got)
	}
}
