package interval

import "testing"

func TestNew_FailsIfIntervalDoesNotExist(t *testing.T) {
	code := "x4"
	if _, err := New(code); err == nil {
		t.Errorf("with %s, got nil-error, expected non-nil error", code)
	}
}
