package media

import (
	n "github.com/Joptim/music-rehearsal/theory/note"
	"testing"
)

func TestGetReadCloser(t *testing.T) {
	table := []string{"A3", "F3", "E#3", "Fb3"}
	for _, name := range table {
		note := n.NewTestHelper(name, t)
		if _, err := GetReadCloser(note); err != nil {
			t.Errorf("with %s, got error \"%v\", expected nil error", table, err)
		}
	}
}

func TestGetReadCloser_FailsIfNoteSoundDoesNotExist(t *testing.T) {
	table := []string{"A-5", "F12"}
	for _, name := range table {
		note := n.NewTestHelper(name, t)
		if _, err := GetReadCloser(note); err == nil {
			t.Errorf("with %s, got nil error, expected non-nil error", table)
		}
	}
}

func TestGetReadClosers(t *testing.T) {
	notes := []n.Note{
		n.NewTestHelper("C3", t),
		n.NewTestHelper("E3", t),
		n.NewTestHelper("G3", t),
	}
	if _, err := GetReadClosers(notes); err != nil {
		t.Errorf("with %v, got error \"%v\", expected nil error", notes, err)
	}
}

func TestGetReadClosers_FailsIfNoteSoundDoesNotExist(t *testing.T) {
	notes := []n.Note{
		n.NewTestHelper("C3", t),
		n.NewTestHelper("E3", t),
		n.NewTestHelper("G15", t),
	}
	if _, err := GetReadClosers(notes); err == nil {
		t.Errorf("with %v, got nil error, expected non-nil error", notes)
	}
}
