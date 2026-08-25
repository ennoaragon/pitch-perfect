package main

import (
	"fmt"

	"github.com/ennoaragon/pitch-perfect/internal/tuner"
	"github.com/gordonklaus/portaudio"
)

func main() {
    fmt.Println("Hello, World!")
    // Endless loop, cycling and always listening,
	// any sound it'll output its note, sometiems itcan be a mix
	// it will display it's findings every second,
	// discard any other aftert the fact

	portaudio.Initialize()
	defer portaudio.Terminate()

	for range 100{
		audioSnippet := tuner.Listener()
		processedAudio := tuner.Procesor(audioSnippet)
		tuner.Analyze(processedAudio)
	}


}

