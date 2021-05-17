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
	return fmt.Sprintf("%s%s%d", n.natural.Name(), accidental, n.octave)
}

func (n Note) SemitonesFrom(from Note) int {
	return n.semitones() - from.semitones()
}

// semitones return the number of semitones from A0
func (n Note) semitones() int {
	semitones := n.natural.Semitones()
	return 12*n.octave + semitones + n.accidental
}

func (n Note) addSemitone() Note {
	var note Note
	var err error
	if n.accidental == 1 {
		nextNatural := n.natural.Next()
		if n.natural.IsB() || n.natural.IsE() {
			note, err = NewFromParams(nextNatural, n.octave, 1)
		} else if n.natural.IsG() {
			note, err = NewFromParams(nextNatural, n.octave+1, 0)
		} else {
			note, err = NewFromParams(nextNatural, n.octave, 0)
		}
	} else {
		note, err = NewFromParams(n.natural, n.octave, n.accidental+1)
	}
	if err != nil {
		panic(fmt.Sprintf("cannot add semitone to note %v", n))
	}
	return note
}

func (n Note) subtractSemitone() Note {
	var note Note
	var err error
	if n.accidental == -1 {
		prevNatural := n.natural.Prev()
		if n.natural.IsC() || n.natural.IsF() {
			note, err = NewFromParams(prevNatural, n.octave, -1)
		} else if n.natural.IsA() {
			note, err = NewFromParams(prevNatural, n.octave-1, 0)
		} else {
			note, err = NewFromParams(prevNatural, n.octave, 0)
		}
	} else {
		note, err = NewFromParams(n.natural, n.octave, n.accidental-1)
	}
	if err != nil {
		panic(fmt.Sprintf("cannot subtract semitone to note %v", n))
	}
	return note
}

func (n Note) AddSemitones(semitones int) Note {
	note := n
	if semitones >= 0 {
		for st := 1; st <= semitones; st++ {
			note = note.addSemitone()
		}
	} else {
		for st := -1; st >= semitones; st-- {
			note = note.subtractSemitone()
		}
	}
	return note
}

}

func New(name string) (Note, error) {
	re := regexp.MustCompile(
		"^(?P<natural>[A-G]{1})" +
			"(?P<accidental>b|#)?" +
			"(?P<sign>-?)(?P<octave>[0-9]+)$")
	match := re.FindStringSubmatch(name)
	if len(match) == 0 {
		return Note{}, fmt.Errorf("cannot create note with name %s", name)
	}

	// Natural
	nat, err := natural.New(match[1])
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

	// Sign
	sign := 1
	if match[3] == "-" {
		sign = -1
	}
	// Octave
	octave, err := strconv.Atoi(match[4])
	if err != nil {
		return Note{}, err
	}

	return NewFromParams(nat, sign*octave, accidental)
}

func NewFromParams(n natural.Natural, octave, accidental int) (Note, error) {
	if accidental < -1 || accidental > 1 {
		return Note{}, fmt.Errorf(
			"cannot create note with natural %v, octave %d and accidental %d",
			n,
			octave,
			accidental,
		)
	}
	return Note{natural: n, octave: octave, accidental: accidental}, nil
}
