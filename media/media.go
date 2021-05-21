package media

import (
	"bytes"
	"embed"
	"fmt"
	n "github.com/Joptim/music-rehearsal/theory/note"
	"io"
)

//go:embed *
var media embed.FS

func GetReadClosers(notes []n.Note) ([]io.ReadCloser, error) {
	output := make([]io.ReadCloser, len(notes))
	for i, note := range notes {
		loadedNote, err := GetReadCloser(note)
		if err != nil {
			return nil, err
		}
		output[i] = loadedNote
	}
	return output, nil
}

func GetReadCloser(note n.Note) (io.ReadCloser, error) {
	if readCloser, err := doGetReadCloser(note); err == nil {
		return readCloser, nil
	}
	equivalent := note.GetEquivalent()
	if equivalent != note {
		if readCloser, err := doGetReadCloser(equivalent); err == nil {
			return readCloser, nil
		}
	}
	return nil, fmt.Errorf("cannot get ReadCloser for note %v", note)
}

func doGetReadCloser(note n.Note) (io.ReadCloser, error) {
	path := fmt.Sprintf("piano/%s.mp3", note.GetName())
	if bytes_, err := media.ReadFile(path); err != nil {
		return nil, err
	} else {
		return io.NopCloser(bytes.NewReader(bytes_)), nil
	}
}
