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
		actual, err := New(test.name)
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
		_, err := New(name)
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
		natural := NewTestHelper(test.natural, t)
		actual := natural.Next()
		if actual.Name() != test.expected {
			t.Errorf("with %s, got %s, expected %s", test.natural, actual.Name(), test.expected)
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
		natural := NewTestHelper(test.natural, t)
		actual := natural.Prev()
		if actual.Name() != test.expected {
			t.Errorf("with %s, got %s, expected %s", test.natural, actual.Name(), test.expected)
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
		actual := NewTestHelper(test.name, t).Semitones()
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
		natural := NewTestHelper(test.natural, t)
		base := NewTestHelper(test.base, t)
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
		natural := NewTestHelper(test.name, t)
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
		natural := NewTestHelper(test.name, t)
		actual := natural.SemitonesFromPrev()
		if actual != test.expected {
			t.Errorf("with %s, got %d, expected %d", test.name, actual, test.expected)
		}
	}
}

func TestNatural_GetName(t *testing.T) {
	table := []string{"A", "B", "C", "D", "E", "F", "G"}
	for _, expected := range table {
		natural := NewTestHelper(expected, t)
		actual := natural.Name()
		if actual != expected {
			t.Errorf("with %s, got %s, expected %s", expected, actual, expected)
		}
	}
}

func TestNatural_AddIntervalSize(t *testing.T) {
	table := []struct {
		natural  string
		size     int
		expected string
	}{
		{"C", 1, "C"},
		{"C", 2, "D"},
		{"C", 3, "E"},
		{"C", 4, "F"},
		{"C", 5, "G"},
		{"C", 6, "A"},
		{"C", 7, "B"},
		{"C", 8, "C"},
		{"C", -1, "C"},
		{"C", -2, "B"},
		{"C", -3, "A"},
		{"C", -4, "G"},
		{"C", -5, "F"},
		{"C", -6, "E"},
		{"C", -7, "D"},
		{"C", -8, "C"},
	}
	for _, test := range table {
		natural := NewTestHelper(test.natural, t)
		actual, _ := natural.AddIntervalSize(test.size)
		if actual.Name() != test.expected {
			t.Errorf(
				"with %s an %d, got %s, expected %s",
				test.natural,
				test.size,
				actual.Name(),
				test.expected,
			)
		}
	}
}

func TestNatural_AddIntervalSize_FailsIfSizeIsInvalid(t *testing.T) {
	natural := NewTestHelper("C", t)
	if _, err := natural.AddIntervalSize(0); err == nil {
		t.Errorf("with C and 0, expected non-nil error, got nil error")
	}
}

func TestNatural_IsA(t *testing.T) {
	A, B := NewTestHelper("A", t), NewTestHelper("B", t)
	if !A.IsA() {
		t.Errorf("with %s, got false, expected true", A.name)
	}
	if B.IsA() {
		t.Errorf("with %s, got true, expected false", B.name)
	}
}

func TestNatural_IsB(t *testing.T) {
	B, C := NewTestHelper("B", t), NewTestHelper("C", t)
	if !B.IsB() {
		t.Errorf("with %s, got false, expected true", B.name)
	}
	if C.IsB() {
		t.Errorf("with %s, got true, expected false", C.name)
	}
}

func TestNatural_IsC(t *testing.T) {
	C, D := NewTestHelper("C", t), NewTestHelper("D", t)
	if !C.IsC() {
		t.Errorf("with %s, got false, expected true", C.name)
	}
	if D.IsC() {
		t.Errorf("with %s, got true, expected false", D.name)
	}
}

func TestNatural_IsD(t *testing.T) {
	D, E := NewTestHelper("D", t), NewTestHelper("E", t)
	if !D.IsD() {
		t.Errorf("with %s, got false, expected true", D.name)
	}
	if E.IsD() {
		t.Errorf("with %s, got true, expected false", E.name)
	}
}

func TestNatural_IsE(t *testing.T) {
	E, A := NewTestHelper("E", t), NewTestHelper("A", t)
	if !E.IsE() {
		t.Errorf("with %s, got false, expected true", E.name)
	}
	if A.IsE() {
		t.Errorf("with %s, got true, expected false", A.name)
	}
}

func TestAllNaturals(t *testing.T) {
	actual := AllNaturals()
	expected := []Natural{
		NewTestHelper("A", t),
		NewTestHelper("B", t),
		NewTestHelper("C", t),
		NewTestHelper("D", t),
		NewTestHelper("E", t),
		NewTestHelper("F", t),
		NewTestHelper("G", t),
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
