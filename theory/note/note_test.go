package note

import "testing"

func TestNew(t *testing.T){
	var table = []struct {
		in string
		expected Note
	}{
		{"A4", Note{"A4", 0, 0}},
		{"Ab4", Note{"A4", 0, 1}},
		{"Abb4", Note{"A4", 0, 2}},
		{"A#4", Note{"A4", 1, 0}},
		{"A##4", Note{"A4", 2, 0}},
		{"E#4", Note{"E4", 1, 0}},
		{"Fb4", Note{"F4", 0, 1}},
	}
	for _, test := range table {
		actual, err := New(test.in)
		if err != nil {
			t.Logf("got error %s, expected %v", err, test.expected)
			t.FailNow()
		}
		if actual != test.expected {
			t.Errorf("got %v, expected %v", actual, test.expected)
		}
	}
}

func TestNewFailsOnInvalidName(t *testing.T){
	var table = []string {"A", "A-1", "A3b", "A#b3", "Ab#3", "Abbb3", "AB3", "A###3", "H3"}
	for _, invalidNote := range table {
		_, err := New(invalidNote)
		if err == nil {
			t.Errorf("got nil error for invalid note %s, expected non-nil error", invalidNote)
		}
	}
}