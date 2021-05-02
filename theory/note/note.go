package note

import (
	"fmt"
	"regexp"
	"strings"
)

type Note struct {
	Natural string
	SharpAccidentals int
	FlatAccidentals int
}

func New(name string) (Note, error){
	re := regexp.MustCompile(
	"^(?P<natural>[A-G]{1})" +
		"(?P<accidental>bb|b|##|#)?" +
		"(?P<octave>[0-7]{1})$")
	match := re.FindStringSubmatch(name)
	if len(match) == 0 {
		return Note{}, fmt.Errorf("cannot create note with name %s", name)
	}

	accidentals := match[2]
	sharpAccidentals, flatAccidentals := 0, 0
	if strings.Contains(accidentals, "#") {
		sharpAccidentals = len(accidentals)
	}
	if strings.Contains(accidentals, "b") {
		flatAccidentals = len(accidentals)
	}
	return Note {
		Natural: fmt.Sprintf("%s%s", match[1], match[3]),
		SharpAccidentals: sharpAccidentals,
		FlatAccidentals: flatAccidentals,
	}, nil
}