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

func (n Note) addSemitone() (Note, error) {
	nat := n.natural
	octave := n.octave
	accidental := n.accidental
	switch accidental {
	case 0:
		if nat.SemitonesToNext() == 2 {
			accidental = 1
			break
		}
		fallthrough
	case 1:
		nat = n.natural.Next()
		if nat.IsA() {
			octave += 1
		}
		fallthrough
	case -1:
		accidental = 0
	}
	return NewFromParams(nat, octave, accidental)
}

func (n Note) SubtractSemitone() (Note, error) {
	nat := n.natural
	octave := n.octave
	accidental := n.accidental
	switch accidental {
	case 0:
		if nat.SemitonesFromPrev() == 2 {
			accidental = -1
			break
		}
		fallthrough
	case -1:
		if nat.IsA() {
			octave -= 1
		}
		nat = n.natural.Prev()
		fallthrough
	case 1:
		accidental = 0
	}
	return NewFromParams(nat, octave, accidental)
}

func (n Note) AddSemitones(semitones int) (Note, error) {
	note := n
	var err error
	if semitones >= 0 {
		for st := 1; st <= semitones; st++ {
			if note, err = note.addSemitone(); err != nil {
				break
			}
		}
	} else {
		for st := -1; st >= semitones; st-- {
			if note, err = note.SubtractSemitone(); err != nil {
				break
			}
		}
	}
	return note, err
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

func NewFromParams(n natural.Natural, octave, accidental int) (Note, error) {
	if octave < 0 || accidental < -1 || accidental > 1 {
		return Note{}, fmt.Errorf(
			"cannot create note with natural %v, octave %d and accidental %d",
			n,
			octave,
			accidental,
		)
	}
	return Note{natural: n, octave: octave, accidental: accidental}, nil
}
