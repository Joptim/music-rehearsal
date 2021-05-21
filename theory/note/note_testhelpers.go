package note

import "testing"

func NewTestHelper(name string, t *testing.T) Note {
	t.Helper()
	note, err := New(name)
	if err != nil {
		t.Fatalf("cannot instantiate note with name %s. Got error %v", name, err)
	}
	return note
}
