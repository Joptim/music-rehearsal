package main

import (
	"log"
	"time"
)

func main() {

	notes := []string{"A3", "B3", "C4", "D4", "E4", "F4", "G4", "A4"}
	mp := NewMediaPlayer()

	err := mp.Prepare(notes...)
	if err != nil {
		log.Fatal(err)
	}

	signal, err := mp.Play(400*time.Millisecond, notes...)
	if err != nil {
		log.Fatal(err)
	}

	<-signal

	err = mp.Release(notes...)
	if err != nil {
		log.Fatal(err)
	}
}
