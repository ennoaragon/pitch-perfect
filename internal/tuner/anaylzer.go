package tuner

import (
	"fmt"
	"math"
	"math/cmplx"
	"github.com/gordonklaus/portaudio"
	"github.com/madelynnblue/go-dsp/fft"
)

const FFTsize int = 2048 //audio samples
const HZSampleRate int = 44100 //hz
const UniversalConcertPitch = 440 //hz

type Analyzer struct {
	Stream *portaudio.Stream
	Buffer []float32
}

type AnalyzeResult struct {
    Note   string
    Octave int
}


func NewAnalyzer() *Analyzer {
	buffer := make([]float32, FFTsize) // Increased for better resolution
	stream, err := portaudio.OpenDefaultStream(1, 0, float64(HZSampleRate), len(buffer), buffer)

	if err != nil {
		panic(err)
	}
	return &Analyzer{Stream: stream, Buffer: buffer}
}

func (a *Analyzer) Listener() []float32 {
	err := a.Stream.Read()
	if err != nil {
		fmt.Println("Error reading stream:", err)
	}

	return a.Buffer
}

func ProcesorSignal(audioSamples []float32) float64{
	n := len(audioSamples)
	fftInput := make([]complex128, n)

	for i := range n{
		// Apply a basic Hann Window to reduce noise leakage
		window := 0.5 * (1 - math.Cos(2*math.Pi*float64(i)/float64(n-1)))
		fftInput[i] = complex(float64(audioSamples[i])*window, 0)
	}

	freq := fft.FFT(fftInput)

	maxMagnitude := 0.0
	strongestBinIndex := 0

	for i :=  range n/2{
		magnitude := cmplx.Abs(freq[i])

		// Skip Bin 0, i guess it's static noise and not worth checking?
		if i > 0 && magnitude > maxMagnitude {
			maxMagnitude = magnitude
			strongestBinIndex = i
		}
	}

	hz := (float64(strongestBinIndex) * float64(HZSampleRate)) / float64(FFTsize)
	return hz
}

// Analyze checks a frequency against musical standards
func Analyze(freq float64 ) AnalyzeResult {

	// Logic to determine if freq note;
	note, octave := hzToNote(freq)

	analyzed := AnalyzeResult{Note: note, Octave: octave}

	fmt.Printf("Note: %s Octave %d\n", analyzed.Note, octave)

	return  analyzed
}

// https://physics.bu.edu/~duffy/sc528_notes03/scale.html
// fractions found here.
// https://golem.ph.utexas.edu/category/2010/02/a_look_at_the_mathematical_ori.html
// This should also return status if near sharp or in tune or out for example,
// or it could be another func
func hzToNote ( hz float64) (string, int) {

	noteMap := [12]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#","A", "A#", "B" }

	semitoneFromA4 := math.Log2(hz/UniversalConcertPitch) * 12

	roundTone := math.Round(semitoneFromA4)

	midiNote := int(roundTone) + 69

	noteIndex := midiNote % 12

	octave := (midiNote / 12 )-1

	return noteMap[noteIndex], octave
}
