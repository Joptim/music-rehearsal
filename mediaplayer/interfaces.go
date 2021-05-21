package mediaplayer

import (
	n "github.com/Joptim/music-rehearsal/theory/note"
	"time"
)

type IMediaPlayer interface {
	Prepare(notes ...n.Note) error
	Play(delay time.Duration, note ...n.Note) (<-chan struct{}, error)
	Release(notes ...n.Note)
	ReleaseAll()
}
