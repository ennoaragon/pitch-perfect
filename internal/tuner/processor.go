package tuner

import "fmt"

// this is where we will do the logic in which we process some freq and it's pitch
// then output it's values here.
func Procesor( snippet  ListenerResult) float32 {

	fmt.Printf("Process audio clip, heard: %f", snippet.AudioSample)

	return snippet.AudioSample[0]
}
