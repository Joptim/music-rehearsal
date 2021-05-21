package mediaplayer

import (
	n "github.com/Joptim/music-rehearsal/theory/note"
	"testing"
)

func TestPrepare(t *testing.T) {
	bf := castToBeepFacade(New(), t)
	notes := []n.Note{
		n.NewTestHelper("A3", t),
		n.NewTestHelper("C3", t),
	}
	if err := bf.Prepare(notes...); err != nil {
		t.Logf("with %v, got error %v, expected nil error", notes, err)
	}
	for _, note := range notes {
		_, exists := bf.buffers[note]
		if !exists {
			t.Logf(
				"note %v not loaded into buffer, expected note to be loaded into buffer",
				note,
			)
		}
	}
}

func TestPrepare_FailsIfSoundDoesNotExist(t *testing.T) {
	bf := castToBeepFacade(New(), t)
	note := n.NewTestHelper("A15", t)
	if err := bf.Prepare(note); err == nil {
		t.Logf("with note %v, got nil error, expected non-nil error", note)
	}
	if _, exists := bf.buffers[note]; exists {
		t.Logf("note %v loaded into buffer, expected note not loaded into buffer", note)
	}
}

func TestPrepare_SkipsAlreadyLoadedSounds(t *testing.T) {
	bf := castToBeepFacade(New(), t)
	note := n.NewTestHelper("A3", t)
	_ = bf.Prepare(note)
	expectedAddress, _ := bf.buffers[note]
	if err := bf.Prepare(note); err != nil {
		t.Errorf("with note %v, got error %v, expected non-nil error", note, err)
	}
	actualAddress, _ := bf.buffers[note]
	if actualAddress != expectedAddress {
		t.Errorf(
			"with note %v, got different addresses %p vs %p, expected address %p",
			note,
			expectedAddress,
			actualAddress,
			expectedAddress,
		)
	}
}

func TestRelease(t *testing.T) {
	bf := castToBeepFacade(New(), t)
	notes := []n.Note{
		n.NewTestHelper("A3", t),
		n.NewTestHelper("C3", t),
	}
	_ = bf.Prepare(notes...)
	bf.Release(notes...)
	for _, note := range notes {
		_, exists := bf.buffers[note]
		if exists {
			t.Logf(
				"note %v not release from buffer, expected note to be release from buffer",
				note,
			)
		}
	}
}

func TestReleaseAll(t *testing.T) {
	bf := castToBeepFacade(New(), t)
	notes := []n.Note{
		n.NewTestHelper("A3", t),
		n.NewTestHelper("C3", t),
	}
	_ = bf.Prepare(notes...)
	bf.ReleaseAll()
	for _, note := range notes {
		_, exists := bf.buffers[note]
		if exists {
			t.Logf(
				"note %v not release from buffer, expected note to be release from buffer",
				note,
			)
		}
	}
}

func castToBeepFacade(player IMediaPlayer, t *testing.T) *beepFacade {
	bf, ok := player.(*beepFacade)
	if !ok {
		t.Fatalf("cannot cast IMediaPlayer %v into *beepFacade", player)
	}
	return bf
}
