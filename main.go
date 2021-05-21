package main

import (
	"github.com/Joptim/music-rehearsal/mediaplayer"
	n "github.com/Joptim/music-rehearsal/theory/note"
	"log"
	"time"
)

func main() {

	notesNames := []string{"A3", "B3", "C4", "D4", "E4", "F4", "G4", "A4"}
	notes := make([]n.Note, len(notesNames))
	var err error
	for i, name := range notesNames {
		if notes[i], err = n.New(name); err != nil {
			panic(err)
		}
	}

	mp := mediaplayer.New()
	err = mp.Prepare(notes...)
	if err != nil {
		log.Fatal(err)
	}

	signal, err := mp.Play(400*time.Millisecond, notes...)
	if err != nil {
		log.Fatal(err)
	}

	<-signal

	mp.Release(notes...)
}
