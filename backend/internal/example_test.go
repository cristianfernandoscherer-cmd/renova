package internal_test

import "testing"

func TestSetup(t *testing.T) {
	if 1 != 1 {
		t.Errorf("expected 1 == 1")
	}
}