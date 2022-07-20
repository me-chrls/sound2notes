package utils

import (
	"fmt"
	"math"
)

type Note int
type Midi int

func (note Midi) ToNote() Note {
	return Note(note - 20)
}
func (note Note) ToMidi() Midi {
	return Midi(note + 20)
}
func (note Midi) NoteString() string {
	return Note(note - 20).NoteString() // todo:maybe not the best approach
}

func (note Note) NoteString() string {
	if note < 1 {
		return "not a note"
	}
	notes := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	letter := notes[(int(note)-4)%len(notes)]
	octave := math.Ceil((float64(note-49))/float64(len(notes))) + 4
	res := letter
	if len(letter) < 2 {
		res += "."
	} else {
		res += ""
	}
	res += fmt.Sprintf("%f %d", octave, note)
	return res
}

func HzToNoteString(freq float64) string {
	return HzToNote(freq).NoteString()
}

// NoteToHz
// https://en.wikipedia.org/wiki/Piano_key_frequencies
func NoteToHz(note int) float64 {
	return 440 * math.Pow(2, (float64(note-49))/12)
}

// HzToNote
// https://en.wikipedia.org/wiki/Piano_key_frequencies
func HzToNote(freq float64) Note {
	return Note(math.Round(12*math.Log2(freq/440) + 49))
}

func HzToMidi(freq float64) Midi {
	return Midi(math.Round(12*math.Log2(freq/440) + 69))
}

func MidiToHz(note int) float64 {
	return 440 * math.Pow(2, float64(note-69)/12)
}
