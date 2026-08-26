package tuner

import (
	"fmt"
	"math"
	"math/cmplx"
	"github.com/gordonklaus/portaudio"
	"github.com/madelynnblue/go-dsp/fft"
)

const FFTsize int = 2048
const HZSampleRate int = 44100

type Analyzer struct {
	Stream *portaudio.Stream
	Buffer []float32
}

type Result struct {
    Note   string
    Status string // "In Tune", "Flat", "Sharp"
	Frequency float64
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
// this will take in the processed audio and output the note
// also include if the note is flat or sharp or tune
func Analyze(freq float64 ) Result {

	// Logic to determine if freq is flat/sharp/in-tune;
	// make some hz to note algo
	// then check if the note is of or not real
	//note := hzToNote(freq)

	analyzed := Result{Note: "A", Status: "Flat", Frequency: freq}

	fmt.Printf("Note: %s", analyzed.Note)

	return  analyzed
}

