package main

import (
	"fmt"
	"github.com/ennoaragon/pitch-perfect/internal/tuner"
	"github.com/gordonklaus/portaudio"
)

func main() {
    // Endless loop, cycling and always listening,
	// any sound it'll output its note, sometiems itcan be a mix
	// it will display it's findings every second,
	// discard any other aftert the fact
	// visual sould mabye be some cool wave or sheet ascii art idk something fun

	portaudio.Initialize()
	defer portaudio.Terminate()

	t := tuner.NewAnalyzer()
	t.Stream.Start()

	defer t.Stream.Stop()
	defer t.Stream.Close()


	fmt.Println("Listening... Press Ctrl+C to stop.")
	for {
		audioSamples := t.Listener()
		hz := tuner.ProcesorSignal(audioSamples)
		if hz > 20 {
			fmt.Printf("\rDetected Frequency: %f Hz ", hz)
			tuner.Analyze(hz)
		}
	}
}

