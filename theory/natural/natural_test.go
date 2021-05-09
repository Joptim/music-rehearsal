package natural

import (
	"testing"
)

func TestNewNatural(t *testing.T) {
	table := []struct {
		name     string
		expected Natural
	}{
		{"A", Natural{"A"}},
		{"B", Natural{"B"}},
		{"C", Natural{"C"}},
		{"D", Natural{"D"}},
		{"E", Natural{"E"}},
		{"F", Natural{"F"}},
		{"G", Natural{"G"}},
	}
	for _, test := range table {
		actual, err := NewNatural(test.name)
		if err != nil {
			t.Errorf("with %s, got error %v, expected %v", test.name, err, test.expected)
		}
		if actual != test.expected {
			t.Errorf("with %s, got %v, expected %v", test.name, actual, test.expected)
		}

	}
}

func TestNewNatural_FailsOnInvalidName(t *testing.T) {
	table := []string{"A0", "Bb4", "C#", "H", "I"}
	for _, name := range table {
		_, err := NewNatural(name)
		if err == nil {
			t.Errorf("with %s, got nil error, expected non-nil error", name)
		}
	}
}

func TestNatural_Next(t *testing.T) {
	table := []struct {
		natural  string
		expected string
	}{
		{"A", "B"},
		{"B", "C"},
		{"C", "D"},
		{"D", "E"},
		{"E", "F"},
		{"F", "G"},
		{"G", "A"},
	}
	for _, test := range table {
		natural := NewNaturalTestHelper(test.natural, t)
		actual, err := natural.Next()
		if err != nil {
			t.Errorf("with %s, got error %v, expected %s", test.natural, err, test.expected)
		}
		if actual.GetName() != test.expected {
			t.Errorf("with %s, got %s, expected %s", test.natural, actual.GetName(), test.expected)
		}
	}
}

func TestNatural_Prev(t *testing.T) {
	table := []struct {
		natural  string
		expected string
	}{
		{"A", "G"},
		{"B", "A"},
		{"C", "B"},
		{"D", "C"},
		{"E", "D"},
		{"F", "E"},
		{"G", "F"},
	}
	for _, test := range table {
		natural := NewNaturalTestHelper(test.natural, t)
		actual, err := natural.Prev()
		if err != nil {
			t.Errorf("with %s, got error %v, expected %s", test.natural, err, test.expected)
		}
		if actual.GetName() != test.expected {
			t.Errorf("with %s, got %s, expected %s", test.natural, actual.GetName(), test.expected)
		}
	}
}

func TestNatural_Semitones(t *testing.T) {
	table := []struct {
		name     string
		expected int
	}{
		{"A", 0},
		{"B", 2},
		{"C", 3},
		{"D", 5},
		{"E", 7},
		{"F", 8},
		{"G", 10},
	}
	for _, test := range table {
		actual := NewNaturalTestHelper(test.name, t).Semitones()
		if test.expected != actual {
			t.Errorf("with %s, got %v, expected %d", test.name, actual, test.expected)
		}
	}
}

func TestNatural_SemitonesBasedOn(t *testing.T) {
	table := []struct {
		base     string
		natural  string
		expected int
	}{
		{"C", "C", 0},
		{"C", "D", 2},
		{"C", "E", 4},
		{"C", "F", 5},
		{"C", "G", 7},
		{"C", "A", 9},
		{"C", "B", 11},
	}
	for _, test := range table {
		natural := NewNaturalTestHelper(test.natural, t)
		base := NewNaturalTestHelper(test.base, t)
		actual := natural.SemitonesBasedOn(base)
		if actual != test.expected {
			t.Errorf(
				"with %s and %s, got %d, expected %d",
				test.base,
				test.natural,
				actual,
				test.expected,
			)
		}
	}
}

func TestNatural_SemitonesToNext(t *testing.T) {
	table := []struct {
		name     string
		expected int
	}{
		{"A", 2},
		{"B", 1},
		{"C", 2},
		{"D", 2},
		{"E", 1},
		{"F", 2},
		{"G", 2},
	}
	for _, test := range table {
		natural := NewNaturalTestHelper(test.name, t)
		actual := natural.SemitonesToNext()
		if actual != test.expected {
			t.Errorf("with %s, got %d, expected %d", test.name, actual, test.expected)
		}
	}
}

func TestNatural_SemitonesFromPrev(t *testing.T) {
	table := []struct {
		name     string
		expected int
	}{
		{"A", 2},
		{"B", 2},
		{"C", 1},
		{"D", 2},
		{"E", 2},
		{"F", 1},
		{"G", 2},
	}
	for _, test := range table {
		natural := NewNaturalTestHelper(test.name, t)
		actual := natural.SemitonesFromPrev()
		if actual != test.expected {
			t.Errorf("with %s, got %d, expected %d", test.name, actual, test.expected)
		}
	}
}

func TestNatural_GetName(t *testing.T) {
	table := []string{"A", "B", "C", "D", "E", "F", "G"}
	for _, expected := range table {
		natural := NewNaturalTestHelper(expected, t)
		actual := natural.GetName()
		if actual != expected {
			t.Errorf("with %s, got %s, expected %s", expected, actual, expected)
		}
	}
}

func TestNatural_IsA(t *testing.T) {
	table := []struct {
		name     string
		expected bool
	}{
		{"A", true},
		{"B", false},
		{"C", false},
		{"D", false},
		{"E", false},
		{"F", false},
		{"G", false},
	}
	for _, test := range table {
		natural := NewNaturalTestHelper(test.name, t)
		actual := natural.IsA()
		if actual != test.expected {
			t.Errorf("with %s, got %t, expected %t", test.name, actual, test.expected)
		}
	}
}

func TestAllNaturals(t *testing.T) {
	actual := AllNaturals()
	expected := []Natural{
		NewNaturalTestHelper("A", t),
		NewNaturalTestHelper("B", t),
		NewNaturalTestHelper("C", t),
		NewNaturalTestHelper("D", t),
		NewNaturalTestHelper("E", t),
		NewNaturalTestHelper("F", t),
		NewNaturalTestHelper("G", t),
	}
	if !testAreEqual(actual, expected, t) {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}

func testAreEqual(a, b []Natural, t *testing.T) bool {
	t.Helper()
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
