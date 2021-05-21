package chord

import (
	"github.com/Joptim/music-rehearsal/theory/note"
	"testing"
)

func TestNew_FailsIfChordIsNotDefined(t *testing.T) {
	code := "M4"
	if _, err := New(code); err == nil {
		t.Errorf("with %s, got nil-error, expected non-nil error", code)
	}
}

func TestChord_Build(t *testing.T) {
	table := []struct {
		chord    string
		base     string
		expected []string
	}{
		{"", "C3", []string{"C3", "E3", "G3"}},
		{"m", "C3", []string{"C3", "Eb3", "G3"}},
		{"aug", "C3", []string{"C3", "E3", "G#3"}},
		{"dim", "C3", []string{"C3", "Eb3", "Gb3"}},
		{"7", "C3", []string{"C3", "E3", "G3", "Bb4"}},
		{"Maj7", "C3", []string{"C3", "E3", "G3", "B4"}},
		{"m7", "C3", []string{"C3", "Eb3", "G3", "Bb4"}},
		{"m7(b5)", "C3", []string{"C3", "Eb3", "Gb3", "Bb4"}},
	}
	for _, test := range table {
		chord, _ := New(test.chord)
		base := note.NewTestHelper(test.base, t)
		chordNotes, _ := chord.Build(base)
		if len(chordNotes) != len(test.expected) {
			t.Errorf(
				"with chord \"%s\" and base %s, got %v, expected %v",
				test.chord,
				test.base,
				chordNotes,
				test.expected,
			)
		}
		for i, n := range chordNotes {
			if n.GetName() != test.expected[i] {
				t.Errorf(
					"with chord \"%s\" and base %s, got %v, expected %v",
					test.chord,
					test.base,
					chordNotes,
					test.expected,
				)
				break
			}
		}
	}
}

func TestChord_Build_Fails(t *testing.T) {
	chord := NewTestHelper("", t)
	base := note.NewTestHelper("E#3", t)
	if _, err := chord.Build(base); err == nil {
		t.Errorf("with code \"%s\" and base %s, got nil-error, expected non-nil error", chord.code, base.GetName())
	}
}
