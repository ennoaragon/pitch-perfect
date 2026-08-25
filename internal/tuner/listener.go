package tuner

// Result represents the state of a note
// some array with bytes containing the wavelength
// there has to be some duration, per second
// it should also feed back real time
import (
	"fmt"
	"github.com/gordonklaus/portaudio"
)

type ListenerResult struct {
	AudioSample [] float32
}

// Analyze checks a frequency against musical standards
func Listener() ListenerResult {
	buffer := make([]float32, 512)

	stream, err := portaudio.OpenDefaultStream(1,0, 44100, len(buffer), buffer)

	chk(err)

	defer stream.Close()

	err = stream.Start()
	chk(err)

	defer stream.Stop()

	// we will listen for a bit
	err = stream.Read()

	fmt.Printf("Captured chunk with first sample: %f - %f\n", buffer[0], buffer[511])
	return ListenerResult{AudioSample: buffer}
}

func chk(err error){

	if err != nil {
		panic(err)
	}
}
