package main

import (
	"fmt"
	"os"
	"sound2notes/audio/pitchdetector/yin"
	"sound2notes/audio/utils"
)

func main() {
	//file, err := os.Open("audio/audio.wav")
	//file, err := os.Open("audio/sine.wav")
	file, err := os.Open("audio/C4v80_1.wav")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("GetPitch")

	analyser := yin.MonoAnalyser("audio/C4v80_1.wav", true, 2048)
	printPitches(analyser)
}

func printPitches(pitches <-chan yin.Pitch) {
	for pitch := range pitches {
		if pitch.Detectedpitch > 0 {
		}
		fmt.Printf("hz value: %v estimated note: %v\n", pitch.Detectedpitch, utils.Midi(pitch.MidiNumber).NoteString())
	}
}

func midiNoteToNoteName(midiNumber int) string {
	noteNames := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	if midiNumber < 21 {
		return ""
	}
	return noteNames[(midiNumber-21)%len(noteNames)]
}
