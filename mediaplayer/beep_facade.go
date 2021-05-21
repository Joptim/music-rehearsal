package mediaplayer

import (
	"fmt"
	"github.com/Joptim/music-rehearsal/media"
	n "github.com/Joptim/music-rehearsal/theory/note"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"time"
)

type beepFacade struct {
	buffers       map[n.Note]*beep.Buffer
	isInitialised bool
}

func (mp *beepFacade) init(sampleRate beep.SampleRate, bufferSize int) error {
	if mp.isInitialised {
		return nil
	}
	if err := speaker.Init(sampleRate, bufferSize); err != nil {
		return err
	}
	mp.isInitialised = true
	return nil
}

// Prepare loads note sounds from media folder into a buffer.
// Calling this method with a note whose sound is already loaded takes no action.
func (mp *beepFacade) Prepare(notes ...n.Note) error {
	for _, note := range notes {
		if _, exists := mp.buffers[note]; exists {
			continue
		}
		file, err := media.GetReadCloser(note)
		if err != nil {
			return err
		}
		streamer, format, err := mp3.Decode(file)
		if err != nil {
			return err
		}

		// Todo: Initialise speaker elsewhere
		if err := mp.init(format.SampleRate, format.SampleRate.N(time.Second/10)); err != nil {
			return err
		}

		// Todo: Resample streamers to a same SampleRate
		buffer := beep.NewBuffer(format)
		buffer.Append(streamer)
		if err := streamer.Close(); err != nil {
			return nil
		}
		mp.buffers[note] = buffer
	}
	return nil
}

// Play plays the given notes with the given delay between them
func (mp *beepFacade) Play(delay time.Duration, notes ...n.Note) (<-chan struct{}, error) {
	for _, note := range notes {
		if _, ok := mp.buffers[note]; !ok {
			return nil, fmt.Errorf("note %s not prepared", note.GetName())
		}
	}

	done := make(chan struct{})
	go mp.play(done, delay, notes...)
	return done, nil
}

func (mp *beepFacade) play(done chan<- struct{}, delay time.Duration, notes ...n.Note) {
	signal := make(chan struct{})
	for i, note := range notes {
		buffer := mp.buffers[note]
		streamer := buffer.Streamer(0, buffer.Len())
		callbackStreamer := beep.Callback(func() {
			signal <- struct{}{}
		})
		speaker.Play(beep.Seq(streamer, callbackStreamer))
		if i < len(notes)-1 {
			<-time.After(delay)
		}
	}
	for i := 0; i < len(notes); i++ {
		<-signal
	}
	close(done)
}

// Release frees note sounds memory from buffer.
// Calling this method with a note whose sound is already released takes no action.
func (mp *beepFacade) Release(notes ...n.Note) {
	for _, note := range notes {
		delete(mp.buffers, note)
	}
}

// ReleaseAll frees all note sounds memory from buffer.
func (mp *beepFacade) ReleaseAll() {
	for key := range mp.buffers {
		delete(mp.buffers, key)
	}
}

func New() IMediaPlayer {
	return &beepFacade{
		buffers:       make(map[n.Note]*beep.Buffer),
		isInitialised: false,
	}
}
