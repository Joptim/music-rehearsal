package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"time"
)

type Player interface {
	Prepare(notes ...string) error
	Play(delay time.Duration, note ...string) (<-chan struct{}, error)
	Release(notes ...string) error
}

type MediaPlayer struct {
	buffers       map[string]*beep.Buffer
	isInitialised bool
}

func (mp *MediaPlayer) init(sampleRate beep.SampleRate, bufferSize int) error {
	if mp.isInitialised {
		return nil
	}
	if err := speaker.Init(sampleRate, bufferSize); err != nil {
		return err
	}
	mp.isInitialised = true
	return nil
}

func (mp *MediaPlayer) Prepare(notes ...string) error {
	for _, note := range notes {
		if _, ok := mp.buffers[note]; ok {
			continue
		}
		file, err := LoadNote(note)
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

//Play plays the given notes with the given delay between them
func (mp *MediaPlayer) Play(delay time.Duration, notes ...string) (<-chan struct{}, error) {
	for _, note := range notes {
		if _, ok := mp.buffers[note]; !ok {
			return nil, fmt.Errorf("note %s not prepared", note)
		}
	}

	done := make(chan struct{})
	go mp.play(done, delay, notes...)
	return done, nil
}

func (mp *MediaPlayer) play(done chan<- struct{}, delay time.Duration, notes ...string) {
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

func (mp *MediaPlayer) Release(notes ...string) error {
	for _, note := range notes {
		delete(mp.buffers, note)
	}
	return nil
}

func NewMediaPlayer() Player {
	return &MediaPlayer{
		buffers:       make(map[string]*beep.Buffer),
		isInitialised: false,
	}
}
