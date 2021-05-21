package interval

import "fmt"

type Interval struct {
	code      string
	name      string
	size      int
	semitones int
}

func (i Interval) Size() int {
	return i.size
}

func (i Interval) Semitones() int {
	return i.semitones
}

var intervals map[string]Interval

func init() {
	intervals = map[string]Interval{
		"P1": {"P1", "Perfect unison", 1, 0},
		"d2": {"d2", "Diminished second", 2, 0},
		"m2": {"m2", "Minor second", 2, 1},
		"A1": {"A1", "Augmented unison", 1, 1},
		"M2": {"M2", "Major second", 2, 2},
		"d3": {"d3", "Diminished third", 3, 2},
		"m3": {"m3", "Minor third", 3, 3},
		"A2": {"A2", "Augmented second", 2, 3},
		"M3": {"M3", "Major third", 3, 4},
		"d4": {"d4", "Diminished fourth", 4, 4},
		"P4": {"P4", "Perfect fourth", 4, 5},
		"A3": {"A3", "Augmented third", 3, 5},
		"d5": {"d5", "Diminished fifth", 5, 6},
		"A4": {"A4", "Augmented fourth", 4, 6},
		"P5": {"P5", "Perfect fifth", 5, 7},
		"d6": {"d6", "Diminished sixth", 6, 7},
		"m6": {"m6", "Minor sixth", 6, 8},
		"A5": {"A5", "Augmented fifth", 5, 8},
		"M6": {"M6", "Major sixth", 6, 9},
		"d7": {"d7", "Diminished seventh", 7, 9},
		"m7": {"m7", "Minor seventh", 7, 10},
		"A6": {"A6", "Augmented sixth", 6, 10},
		"M7": {"M7", "Major seventh", 7, 11},
		"d8": {"d8", "Diminished octave", 8, 11},
		"P8": {"P8", "Perfect octave", 8, 12},
		"A7": {"A7", "Augmented seventh", 7, 12},
	}
}

func New(code string) (Interval, error) {
	interval, exists := intervals[code]
	if !exists {
		return Interval{}, fmt.Errorf("cannot get interval from code %s", code)
	}
	return interval, nil
}

func NewOrPanic(code string) Interval {
	interval, err := New(code)
	if err != nil {
		panic(err)
	}
	return interval
}
