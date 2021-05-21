package chord

import "testing"

func NewTestHelper(code string, t *testing.T) Chord {
	t.Helper()
	chord, err := New(code)
	if err != nil {
		t.Fatalf("cannot instantiate chord with code %s. Got error %v", code, err)
	}
	return chord
}
