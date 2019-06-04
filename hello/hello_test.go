package hello

import "testing"

func HelloTest(t *testing.T) {
	if got := Hello(); got != "Hello" {
		t.Errorf("got: %s", got)
	}
}
