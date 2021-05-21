package chord

import (
	"fmt"
	"github.com/Joptim/music-rehearsal/theory/interval"
	"github.com/Joptim/music-rehearsal/theory/note"
)

type Chord struct {
	code      string
	alias     []string
	name      string
	intervals []interval.Interval
}

func (c Chord) GetAlias() []string {
	return c.alias
}

func (c Chord) GetName() string {
	return c.name
}

func (c Chord) Build(base note.Note) ([]note.Note, error) {
	var err error
	notes := make([]note.Note, len(c.intervals)+1)
	notes[0] = base
	for idx, in := range c.intervals {
		if notes[idx+1], err = base.AddInterval(in); err != nil {
			return []note.Note{},
				fmt.Errorf("cannot build chord %v from base note %v. Got error %v", c, base, err)
		}
	}
	return notes, nil
}

func New(code string) (Chord, error) {
	chord, exists := chords[code]
	if !exists {
		return Chord{}, fmt.Errorf("cannot get chord from code %s", code)
	}
	return chord, nil
}

var chords map[string]Chord

func init() {
	chords = make(map[string]Chord)
	m3 := interval.NewOrPanic("m3")
	M3 := interval.NewOrPanic("M3")
	d5 := interval.NewOrPanic("d5")
	P5 := interval.NewOrPanic("P5")
	A5 := interval.NewOrPanic("A5")
	m7 := interval.NewOrPanic("m7")
	M7 := interval.NewOrPanic("M7")

	chords = map[string]Chord{
		"":       {"", []string{"M", "maj"}, "Major triad", []interval.Interval{M3, P5}},
		"m":      {"m", []string{"min", "-"}, "Minor triad", []interval.Interval{m3, P5}},
		"aug":    {"aug", []string{"+"}, "Augmented triad", []interval.Interval{M3, A5}},
		"dim":    {"dim", []string{"º"}, "Diminished triad", []interval.Interval{m3, d5}},
		"7":      {"7", []string{}, "Dominant seventh", []interval.Interval{M3, P5, m7}},
		"Maj7":   {"Maj7", []string{"M7"}, "Major seventh", []interval.Interval{M3, P5, M7}},
		"m7":     {"m7", []string{}, "Minor seventh", []interval.Interval{m3, P5, m7}},
		"m7(b5)": {"m7(b5)", []string{"m7b5"}, "Half-diminished seventh", []interval.Interval{m3, d5, m7}},
	}
}
