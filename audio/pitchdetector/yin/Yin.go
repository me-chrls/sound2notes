package yin

/*
YIN Pitch Detection Algorithm
at https://github.com/cheesedosa/yingo
*/

import (
	"github.com/go-audio/wav"
	"log"
	"math"
	"os"
)

//import "fmt"

var YIN_SAMPLING_RATE float32 = 44100

type Yin struct {
	BufferSize  int
	yinBuffer   *[]float32
	probability float32
	Threshold   float32
}

// API

func (y *Yin) YinInit(bufSize int, thresh float32) {
	y.BufferSize = bufSize
	y.Threshold = thresh

	buff := make([]float32, y.BufferSize/2)
	y.yinBuffer = &buff

}

func (y *Yin) GetPitch(d *[]float32) float32 {

	tauEstimate := -1
	var pitchInHertz float32 = -1

	y.yinDiff(d)

	y.yinCMND()

	tauEstimate = y.yinAbsThresh()

	if tauEstimate != -1 {
		pitchInHertz = YIN_SAMPLING_RATE / y.yinPI(tauEstimate)
	}

	return pitchInHertz

}

func (y *Yin) GetProb() float32 {
	return y.probability
}

//Yin private methods

// Step1: ACF

//Step2: Improving on the autocorrelation function using amplitude difference for each window at different time shifts tau

func (y *Yin) yinDiff(data *[]float32) {

	var delta float32
	if len(*data) < 10 {
		return
	}
	for tau := 0; tau < y.BufferSize/2; tau++ {
		//fmt.Println(tau)
		for i := 0; i < y.BufferSize/2; i++ {
			//fmt.Println(i , tau, len(*data))
			delta = (*data)[i] - (*data)[i+tau]

			(*y.yinBuffer)[tau] += delta * delta
		}

	}

}

//Step3: Cummulative Mean Normal Difference to deal with zero-lag errors post difference function. Set the first zero-lag difference
//       to 1 to deal with too high errors.

func (y *Yin) yinCMND() {

	var runningSum float32
	(*y.yinBuffer)[0] = 1

	for tau := 1; tau < y.BufferSize/2; tau++ {
		runningSum += (*y.yinBuffer)[tau]
		(*y.yinBuffer)[tau] *= float32(tau) / runningSum
	}
}

//Step4: Thresholding to pick the frst dip(the difference) lower than the threshold to reduce octave errors.

func (y *Yin) yinAbsThresh() int {

	var tau int
	for tau = 2; tau < y.BufferSize/2; tau++ {
		if (*y.yinBuffer)[tau] < y.Threshold {
			for tau+1 < y.BufferSize/2 && (*y.yinBuffer)[tau+1] < (*y.yinBuffer)[tau] {
				tau++
			}

			y.probability = 1 - y.Threshold
			break
		}
	}

	if tau == y.BufferSize/2 || (*y.yinBuffer)[tau] >= y.Threshold {
		tau = -1
		y.probability = 0
	}

	return tau

}

//Step5: The process is carried out for integer time-shifts (multiples of sampling rate). However, there may be a better
//       overlap at a non-integer time-shift (tau). Fit a parabolic curve to get
//       a better non-integer estimate.

func (y *Yin) yinPI(tauEstimate int) float32 {

	var betterTau float32
	var x0, x2 int

	if tauEstimate < 0 {
		x0 = tauEstimate
	} else {
		x0 = tauEstimate - 1
	}

	if tauEstimate+1 < y.BufferSize/2 {
		x2 = tauEstimate + 1
	} else {
		x2 = tauEstimate
	}

	if x0 == tauEstimate {
		if (*y.yinBuffer)[tauEstimate] <= (*y.yinBuffer)[x2] {
			betterTau = float32(tauEstimate)
		} else {
			betterTau = float32(x0)
		}
	} else if x2 == tauEstimate {
		if (*y.yinBuffer)[tauEstimate] <= (*y.yinBuffer)[x0] {
			betterTau = float32(tauEstimate)
		} else {
			betterTau = float32(x0)
		}
	} else {
		var s0, s1, s2 float32
		s0 = (*y.yinBuffer)[x0]
		s1 = (*y.yinBuffer)[tauEstimate]
		s2 = (*y.yinBuffer)[x2]

		betterTau = float32(tauEstimate) + (s2-s0)/(2*(2*s1-s2-s0))

	}

	return betterTau
}

type Pitch struct {
	HopStamp         int
	Detectedpitch    float32
	PitchProbability float32
	StdFrequency     float32
	MidiNumber       int
}

func MonoAnalyser(f string, bufferapproximate bool, hopSize int) <-chan Pitch {
	// Open the file
	rawFile, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer rawFile.Close()

	// Create a decoder for the file
	wavDecoder := wav.NewDecoder(rawFile)
	pcm, err := wavDecoder.FullPCMBuffer()
	if err != nil {
		log.Fatal(err)
	}

	// Maps note names -> hz values
	freqArray := loadFreqArray()

	// Load the data from the wave file into memory directly
	intBuffer := pcm.AsIntBuffer().Data
	pcmArray := make([]float32, len(intBuffer))

	// Yin alg expects float32s
	for i, val := range intBuffer {
		pcmArray[i] = float32(val)
	}

	// Get the number of iterations we need to do
	iterations := len(pcmArray) / hopSize
	pch := make(chan Pitch, iterations)

	for i := 0; i < iterations; i++ {
		// Init yin, create variables
		yin := Yin{}
		yin.YinInit(hopSize, float32(0.05))
		batch := make([]float32, hopSize)
		batch = pcmArray[i*hopSize : (i*hopSize + hopSize)]
		pitch := Pitch{HopStamp: i}

		// Do the Yin alg computations, add to our data structure
		pitch.Detectedpitch = yin.GetPitch(&batch)
		pitch.PitchProbability = yin.GetProb()
		pitch.StdFrequency, pitch.MidiNumber = moarData(pitch.Detectedpitch, freqArray)

		pch <- pitch // Send to the channel
	}

	close(pch)
	return pch
}
func loadFreqArray() *[88]float32 {
	// Loop over all MIDI values
	var frArr [88]float32
	for i := 21; i <= 108; i++ {
		//midi to frequency; A4 = 440 Hz

		//very dirty; could use 1 << (exponent)or 1 >> (exponent) after casting to unsigned ints; not decided
		frArr[i-21] = float32(math.Pow(2, float64((i-69))/float64(12.0)) * 440)

	}
	return &frArr
}

func basicAbs(x float32) float32 {

	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}

func moarData(p float32, freqArray *[88]float32) (float32, int) {

	pitch := p

	if pitch == -1 {
		return 0, 0
	}

	smallestDiff := basicAbs(pitch - (*freqArray)[0])
	var stdFrequency float32
	var midiNumber int
	for n, val := range *freqArray {
		xDiff := basicAbs(pitch - val)
		if xDiff < smallestDiff {
			smallestDiff = xDiff
			stdFrequency = val
			midiNumber = n + 21
		}
	}

	return stdFrequency, midiNumber

}
