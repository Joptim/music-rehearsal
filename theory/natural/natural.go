package natural

import (
	"fmt"
)

var naturalA Natural
var naturalsNames []string
var naturals map[string]Natural

type Natural struct {
	name string
}

func (n Natural) Name() string {
	return n.name
}

func (n Natural) Next() Natural {
	for pos, name := range naturalsNames {
		if n.name == name {
			nextPos := (pos + 1) % len(naturalsNames)
			nextNaturalName := naturalsNames[nextPos]
			return naturals[nextNaturalName]
		}
	}
	panic(fmt.Sprintf("cannot find next Natural from %v", n))
}

func (n Natural) Prev() Natural {
	// range naturalsNames in reverse order
	for pos := len(naturalsNames) - 1; pos >= 0; pos-- {
		name := naturalsNames[pos]
		if n.name == name {
			prevPos := (pos - 1) % len(naturalsNames)
			if prevPos < 0 {
				prevPos += len(naturalsNames)
			}
			prevNaturalName := naturalsNames[prevPos]
			return naturals[prevNaturalName]
		}
	}
	panic(fmt.Sprintf("cannot find previous Natural from %v", n))
}

// Semitones return the number of Semitones from A
func (n Natural) Semitones() int {
	semitones := 0
	current := naturalA
	for current != n {
		next := current.Next()
		if (next.name == "C" && current.name == "B") ||
			(next.name == "F" && current.name == "E") {
			semitones += 1
		} else {
			semitones += 2
		}
		current = next
	}
	return semitones
}

// SemitonesBasedOn return the number of Semitones
// of a Natural note from another Natural note
func (n Natural) SemitonesBasedOn(from Natural) int {
	diff := n.Semitones() - from.Semitones()
	if diff < 0 {
		diff += 12
	}
	return diff
}

func (n Natural) SemitonesToNext() int {
	return n.Next().SemitonesBasedOn(n)
}

func (n Natural) SemitonesFromPrev() int {
	return n.SemitonesBasedOn(n.Prev())
}

func (n Natural) AddIntervalSize(size int) (Natural, error) {
	if size == 0 {
		return Natural{}, fmt.Errorf("cannot add interval size %d to natural %v", size, n)
	} else if size > 0 {
		size -= 1
	} else {
		size += 1
	}
	// Find natural position
	for pos, name := range naturalsNames {
		if name == n.name {
			pos = (pos + size) % len(naturalsNames)
			if pos < 0 {
				pos += len(naturalsNames)
			}
			return naturals[naturalsNames[pos]], nil
		}
	}
	panic(fmt.Sprintf("cannot add interval size %d to natural %v", size, n))
}

func (n Natural) IsA() bool {
	return n.name == "A"
}

func New(name string) (Natural, error) {
	if natural, exists := naturals[name]; exists {
		return natural, nil
	} else {
		return Natural{}, fmt.Errorf("cannot find Natural with name %s", name)
	}
}

func AllNaturals() []Natural {
	all := make([]Natural, len(naturalsNames))
	for i, name := range naturalsNames {
		all[i] = naturals[name]
	}
	return all
}

func init() {
	// init naturalsNames
	naturalsNames = []string{"A", "B", "C", "D", "E", "F", "G"}

	// init naturals
	naturals = make(map[string]Natural)
	for _, name := range naturalsNames {
		nat := Natural{name}
		naturals[name] = nat
	}

	// init naturalA
	var err error
	if naturalA, err = New("A"); err != nil {
		panic(err)
	}
}
