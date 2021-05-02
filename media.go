package main

import (
	"bytes"
	"embed"
	"fmt"
	"io"
)

//go:embed media
var media embed.FS

func LoadNote(note string) (io.ReadCloser, error) {
	path := fmt.Sprintf("media/piano/%s.mp3", note)
	bytes_, err := media.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return io.NopCloser(bytes.NewReader(bytes_)), nil
}

func LoadNotes(notes []string) ([]io.ReadCloser, error) {
	output := make([]io.ReadCloser, len(notes))
	for i, note := range notes {
		loadedNote, err := LoadNote(note)
		if err != nil {
			return nil, err
		}
		output[i] = loadedNote
	}
	return output, nil
}
