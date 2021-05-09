package note

import (
	"fmt"
	"github.com/Joptim/music-rehearsal/theory/natural"
	"regexp"
	"strconv"
	"strings"
)

type Note struct {
	natural natural.Natural
	octave  int
	// accidental equals -1 for flat notes,
	// 0 for natural notes and 1 for sharp notes.
	accidental int
}

func (n Note) GetName() string {
	var accidental string
	switch n.accidental {
	case -1:
		accidental = "b"
	case 0:
		accidental = ""
	case 1:
		accidental = "#"
	}
	return fmt.Sprintf("%s%s%d", n.natural.GetName(), accidental, n.octave)
}

func (n Note) SemitonesFrom(from Note) int {
	return n.semitones() - from.semitones()
}

// semitones return the number of semitones from A0
func (n Note) semitones() int {
	semitones := n.natural.Semitones()
	return 12*n.octave + semitones + n.accidental
}

func (n Note) AddSemitone() Note {
	note := n
	switch n.accidental {
	case 0:
		if note.natural.SemitonesToNext() == 2 {
			note.accidental = 1
			break
		}
		fallthrough
	case 1:
		note.natural, _ = n.natural.Next()
		if note.natural.IsA() {
			note.octave += 1
		}
		fallthrough
	case -1:
		note.accidental = 0
	}
	return note
}

func (n Note) SubtractSemitone() Note {
	note := n
	switch n.accidental {
	case 0:
		if note.natural.SemitonesFromPrev() == 2 {
			note.accidental = -1
			break
		}
		fallthrough
	case -1:
		note.natural, _ = n.natural.Prev()
		if n.natural.IsA() {
			note.octave -= 1
		}
		fallthrough
	case 1:
		note.accidental = 0
	}
	return note
}

func (n Note) AddSemitones(semitones int) Note {
	note := n
	if semitones >= 0 {
		for st := 1; st <= semitones; st++ {
			note = note.AddSemitone()
		}
	} else {
		for st := -1; st >= semitones; st-- {
			note = note.SubtractSemitone()
		}
	}
	return note
}

func New(name string) (Note, error) {
	re := regexp.MustCompile(
		"^(?P<natural>[A-G]{1})" +
			"(?P<accidental>b|#)?" +
			"(?P<octave>[0-7]{1})$")
	match := re.FindStringSubmatch(name)
	if len(match) == 0 {
		return Note{}, fmt.Errorf("cannot create note with name %s", name)
	}

	// Natural
	nat, err := natural.NewNatural(match[1])
	if err != nil {
		return Note{}, err
	}

	// Accidentals
	accidental := 0
	if strings.Contains(match[2], "#") {
		accidental = 1
	}
	if strings.Contains(match[2], "b") {
		accidental = -1
	}

	// Octave
	octave, err := strconv.Atoi(match[3])
	if err != nil {
		return Note{}, err
	}

	return Note{
		natural:    nat,
		octave:     octave,
		accidental: accidental,
	}, nil
}
