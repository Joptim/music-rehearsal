package note

import (
	"github.com/Joptim/music-rehearsal/theory/natural"
	"testing"
)

func TestNew(t *testing.T) {
	table := []struct {
		in       string
		expected Note
	}{
		{"A0", Note{natural.NewNaturalTestHelper("A", t), 0, 0}},
		{"Bb1", Note{natural.NewNaturalTestHelper("B", t), 1, -1}},
		{"B1", Note{natural.NewNaturalTestHelper("B", t), 1, 0}},
		{"B#1", Note{natural.NewNaturalTestHelper("B", t), 1, 1}},
		{"Cb2", Note{natural.NewNaturalTestHelper("C", t), 2, -1}},
		{"C2", Note{natural.NewNaturalTestHelper("C", t), 2, 0}},
		{"C#2", Note{natural.NewNaturalTestHelper("C", t), 2, 1}},
		{"Db3", Note{natural.NewNaturalTestHelper("D", t), 3, -1}},
		{"D3", Note{natural.NewNaturalTestHelper("D", t), 3, 0}},
		{"D#3", Note{natural.NewNaturalTestHelper("D", t), 3, 1}},
		{"Eb4", Note{natural.NewNaturalTestHelper("E", t), 4, -1}},
		{"E4", Note{natural.NewNaturalTestHelper("E", t), 4, 0}},
		{"E#4", Note{natural.NewNaturalTestHelper("E", t), 4, 1}},
		{"Fb5", Note{natural.NewNaturalTestHelper("F", t), 5, -1}},
		{"F5", Note{natural.NewNaturalTestHelper("F", t), 5, 0}},
		{"F#5", Note{natural.NewNaturalTestHelper("F", t), 5, 1}},
		{"Gb6", Note{natural.NewNaturalTestHelper("G", t), 6, -1}},
		{"G6", Note{natural.NewNaturalTestHelper("G", t), 6, 0}},
		{"G#6", Note{natural.NewNaturalTestHelper("G", t), 6, 1}},
		{"Ab7", Note{natural.NewNaturalTestHelper("A", t), 7, -1}},
	}
	for _, test := range table {
		actual, _ := New(test.in)
		if actual != test.expected {
			t.Errorf("with %s, got %v, expected %v", test.in, actual, test.expected)
		}
	}
}

func TestNew_FailsOnInvalidName(t *testing.T) {
	table := []string{
		"A",
		"A-1",
		"A3b",
		"A#b3",
		"Abb3",
		"Ab#3",
		"Abbb3",
		"AB3",
		"A##3",
		"A###3",
		"H3",
	}
	for _, invalidNote := range table {
		_, err := New(invalidNote)
		if err == nil {
			t.Errorf("with %s, got nil error, expected non-nil error", invalidNote)
		}
	}
}

func TestNewFromParams(t *testing.T) {
	table := []struct {
		natural    string
		octave     int
		accidental int
		expected   Note
	}{
		{"A", 0, 0, Note{natural.NewNaturalTestHelper("A", t), 0, 0}},
		{"B", 1, -1, Note{natural.NewNaturalTestHelper("B", t), 1, -1}},
		{"B", 1, 0, Note{natural.NewNaturalTestHelper("B", t), 1, 0}},
		{"B", 1, 1, Note{natural.NewNaturalTestHelper("B", t), 1, 1}},
		{"C", 2, -1, Note{natural.NewNaturalTestHelper("C", t), 2, -1}},
		{"C", 2, 0, Note{natural.NewNaturalTestHelper("C", t), 2, 0}},
		{"C", 2, 1, Note{natural.NewNaturalTestHelper("C", t), 2, 1}},
		{"D", 3, -1, Note{natural.NewNaturalTestHelper("D", t), 3, -1}},
		{"D", 3, 0, Note{natural.NewNaturalTestHelper("D", t), 3, 0}},
		{"D", 3, 1, Note{natural.NewNaturalTestHelper("D", t), 3, 1}},
		{"E", 4, -1, Note{natural.NewNaturalTestHelper("E", t), 4, -1}},
		{"E", 4, 0, Note{natural.NewNaturalTestHelper("E", t), 4, 0}},
		{"E", 4, 1, Note{natural.NewNaturalTestHelper("E", t), 4, 1}},
		{"F", 5, -1, Note{natural.NewNaturalTestHelper("F", t), 5, -1}},
		{"F", 5, 0, Note{natural.NewNaturalTestHelper("F", t), 5, 0}},
		{"F", 5, 1, Note{natural.NewNaturalTestHelper("F", t), 5, 1}},
		{"G", 6, -1, Note{natural.NewNaturalTestHelper("G", t), 6, -1}},
		{"G", 6, 0, Note{natural.NewNaturalTestHelper("G", t), 6, 0}},
		{"G", 6, 1, Note{natural.NewNaturalTestHelper("G", t), 6, 1}},
		{"A", 7, -1, Note{natural.NewNaturalTestHelper("A", t), 7, -1}},
	}
	for _, test := range table {
		nat := natural.NewNaturalTestHelper(test.natural, t)
		actual, _ := NewFromParams(nat, test.octave, test.accidental)
		if actual != test.expected {
			t.Logf(
				"with %s, %d and %d, got error %v, expected %v",
				test.natural,
				test.octave,
				test.accidental,
				actual,
				test.expected,
			)
		}
	}
}

func TestNewFromParams_FailsOnInvalidParams(t *testing.T) {
	table := []struct {
		natural    string
		octave     int
		accidental int
	}{
		{"A", -1, 0},
		{"A", 3, -2},
		{"A", 3, 2},
	}
	for _, test := range table {
		nat := natural.NewNaturalTestHelper(test.natural, t)
		_, err := NewFromParams(nat, test.octave, test.accidental)
		if err == nil {
			t.Errorf(
				"with %s, %d and %d, got nil error, expected non-nil error",
				test.natural,
				test.octave,
				test.accidental,
			)
		}
	}
}

func TestNote_GetName(t *testing.T) {
	table := []string{
		"A0",
		"Bb1",
		"B1",
		"B#1",
		"Cb2",
		"C2",
		"C#2",
		"Db3",
		"D3",
		"D#3",
		"Eb4",
		"E4",
		"E#4",
		"Fb5",
		"F5",
		"F#5",
		"Gb6",
		"G6",
		"G#6",
		"Ab7",
	}
	for _, name := range table {
		note := NewHelper(name, t)
		actual := note.GetName()
		if actual != name {
			t.Errorf("with %s, got %s, expected %s", name, actual, name)
		}
	}
}

func TestNote_Semitones(t *testing.T) {
	table := []struct {
		name     string
		expected int
	}{
		{"A0", 0},
		{"Bb1", 13},
		{"B1", 14},
		{"B#1", 15},
		{"Cb2", 26},
		{"C2", 27},
		{"C#2", 28},
		{"Db3", 40},
		{"D3", 41},
		{"D#3", 42},
		{"Eb4", 54},
		{"E4", 55},
		{"E#4", 56},
		{"Fb5", 67},
		{"F5", 68},
		{"F#5", 69},
		{"Gb6", 81},
		{"G6", 82},
		{"G#6", 83},
		{"Ab7", 83},
	}
	for _, test := range table {
		note := NewHelper(test.name, t)
		actual := note.semitones()
		if actual != test.expected {
			t.Errorf("with %s, got %v, expected %v", test.name, actual, test.expected)
		}
	}
}

func TestNote_SemitonesFrom(t *testing.T) {
	table := []struct {
		from     string
		note     string
		expected int
	}{
		{"C3", "C3", 0},
		{"C3", "C#3", 1},
		{"C3", "Db3", 1},
		{"C3", "D3", 2},
		{"C3", "D#3", 3},
		{"C3", "Eb3", 3},
		{"C3", "E3", 4},
		{"C3", "F3", 5},
		{"C3", "F#3", 6},
		{"C3", "Gb3", 6},
		{"C3", "G3", 7},
		{"C3", "G#3", 8},
		{"C3", "Ab4", 8},
		{"C3", "A4", 9},
		{"C3", "A#4", 10},
		{"C3", "Bb4", 10},
		{"C3", "B4", 11},
		{"C3", "C4", 12},
		{"C3", "B3", -1},
		{"C3", "Bb3", -2},
		{"C3", "A#3", -2},
		{"C3", "A3", -3},
		{"C3", "Ab3", -4},
		{"C3", "G#2", -4},
		{"C3", "G2", -5},
		{"C3", "Gb2", -6},
		{"C3", "F#2", -6},
		{"C3", "F2", -7},
		{"C3", "E2", -8},
		{"C3", "Eb2", -9},
		{"C3", "D#2", -9},
		{"C3", "D2", -10},
		{"C3", "Db2", -11},
		{"C3", "C#2", -11},
		{"C3", "C2", -12},
	}

	for _, test := range table {
		note := NewHelper(test.note, t)
		fromNote := NewHelper(test.from, t)
		actual := note.SemitonesFrom(fromNote)
		if actual != test.expected {
			t.Errorf("with %s and %s, got %v, expected %v", test.from, test.note, actual, test.expected)
		}
	}
}

func TestNote_AddSemitone(t *testing.T) {
	table := []struct {
		note     Note
		expected Note
	}{
		{NewHelper("C3", t), NewHelper("C#3", t)},
		{NewHelper("C#3", t), NewHelper("D3", t)},
		{NewHelper("D3", t), NewHelper("D#3", t)},
		{NewHelper("D#3", t), NewHelper("E3", t)},
		{NewHelper("E3", t), NewHelper("F3", t)},
		{NewHelper("F3", t), NewHelper("F#3", t)},
		{NewHelper("F#3", t), NewHelper("G3", t)},
		{NewHelper("G3", t), NewHelper("G#3", t)},
		{NewHelper("G#3", t), NewHelper("A4", t)},
		{NewHelper("A4", t), NewHelper("A#4", t)},
		{NewHelper("A#4", t), NewHelper("B4", t)},
		{NewHelper("Ab3", t), NewHelper("A3", t)},
		{NewHelper("Bb3", t), NewHelper("B3", t)},
		{NewHelper("Cb3", t), NewHelper("C3", t)},
		{NewHelper("Db3", t), NewHelper("D3", t)},
		{NewHelper("Eb3", t), NewHelper("E3", t)},
		{NewHelper("Fb3", t), NewHelper("F3", t)},
		{NewHelper("Gb3", t), NewHelper("G3", t)},
	}
	for _, test := range table {
		actual, _ := test.note.addSemitone()
		if test.expected != actual {
			t.Errorf(
				"from %v, got %v, expected %v",
				test.note,
				actual,
				test.expected,
			)
		}
	}
}

func TestNote_SubtractSemitone(t *testing.T) {
	table := []struct {
		note     Note
		expected Note
	}{
		{NewHelper("C3", t), NewHelper("B3", t)},
		{NewHelper("B3", t), NewHelper("Bb3", t)},
		{NewHelper("Bb3", t), NewHelper("A3", t)},
		{NewHelper("A3", t), NewHelper("Ab3", t)},
		{NewHelper("Ab3", t), NewHelper("G2", t)},
		{NewHelper("G2", t), NewHelper("Gb2", t)},
		{NewHelper("Gb2", t), NewHelper("F2", t)},
		{NewHelper("F2", t), NewHelper("E2", t)},
		{NewHelper("E2", t), NewHelper("Eb2", t)},
		{NewHelper("Eb2", t), NewHelper("D2", t)},
		{NewHelper("D2", t), NewHelper("Db2", t)},
		{NewHelper("Db2", t), NewHelper("C2", t)},
		{NewHelper("A#3", t), NewHelper("A3", t)},
		{NewHelper("B#3", t), NewHelper("B3", t)},
		{NewHelper("C#3", t), NewHelper("C3", t)},
		{NewHelper("D#3", t), NewHelper("D3", t)},
		{NewHelper("E#3", t), NewHelper("E3", t)},
		{NewHelper("F#3", t), NewHelper("F3", t)},
		{NewHelper("G#3", t), NewHelper("G3", t)},
	}
	for _, test := range table {
		actual, _ := test.note.SubtractSemitone()
		if test.expected != actual {
			t.Errorf(
				"from %v, got %v, expected %v",
				test.note,
				actual,
				test.expected,
			)
		}
	}
}

func TestNote_SubtractSemitone_FailsIfOutOfRange(t *testing.T) {
	Ab0 := NewHelper("Ab0", t)
	if _, err := Ab0.SubtractSemitone(); err == nil {
		t.Errorf("with %v, got error %v, expected non-nil error", Ab0.GetName(), err)
	}
}

func TestNote_AddSemitones(t *testing.T) {
	table := []struct {
		note      Note
		semitones int
		expected  Note
	}{
		{NewHelper("C3", t), 0, NewHelper("C3", t)},
		{NewHelper("C3", t), 1, NewHelper("C#3", t)},
		{NewHelper("C3", t), 2, NewHelper("D3", t)},
		{NewHelper("C3", t), 3, NewHelper("D#3", t)},
		{NewHelper("C3", t), 4, NewHelper("E3", t)},
		{NewHelper("C3", t), 5, NewHelper("F3", t)},
		{NewHelper("C3", t), 6, NewHelper("F#3", t)},
		{NewHelper("C3", t), 7, NewHelper("G3", t)},
		{NewHelper("C3", t), 8, NewHelper("G#3", t)},
		{NewHelper("C3", t), 9, NewHelper("A4", t)},
		{NewHelper("C3", t), 10, NewHelper("A#4", t)},
		{NewHelper("C3", t), 11, NewHelper("B4", t)},
		{NewHelper("C3", t), 12, NewHelper("C4", t)},
		{NewHelper("C3", t), -1, NewHelper("B3", t)},
		{NewHelper("C3", t), -2, NewHelper("Bb3", t)},
		{NewHelper("C3", t), -3, NewHelper("A3", t)},
		{NewHelper("C3", t), -4, NewHelper("Ab3", t)},
		{NewHelper("C3", t), -5, NewHelper("G2", t)},
		{NewHelper("C3", t), -6, NewHelper("Gb2", t)},
		{NewHelper("C3", t), -7, NewHelper("F2", t)},
		{NewHelper("C3", t), -8, NewHelper("E2", t)},
		{NewHelper("C3", t), -9, NewHelper("Eb2", t)},
		{NewHelper("C3", t), -10, NewHelper("D2", t)},
		{NewHelper("C3", t), -11, NewHelper("Db2", t)},
		{NewHelper("C3", t), -12, NewHelper("C2", t)},
	}
	for _, test := range table {
		actual, _ := test.note.AddSemitones(test.semitones)
		if test.expected != actual {
			t.Errorf(
				"from %v and %d, got %v, expected %v",
				test.note,
				test.semitones,
				actual,
				test.expected,
			)
		}
	}
}

func TestNote_AddSemitones_FailsIfOutOfRange(t *testing.T) {
	A0 := NewHelper("A0", t)
	if _, err := A0.AddSemitones(-5); err == nil {
		t.Errorf("with %v, got error %v, expected non-nil error", A0.GetName(), err)
	}
}

func NewHelper(name string, t *testing.T) Note {
	t.Helper()
	note, err := New(name)
	if err != nil {
		t.Fatalf("cannot instantiate note with name %s. Got error %v", name, err)
	}
	return note
}
