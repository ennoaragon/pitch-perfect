package tuner

// Result represents the state of a note
// some array with bytes containing the wavelength
// there has to be some duration, per second
// it should also feed back real time

type ListenerResult struct {
	audio string
}

// Analyze checks a frequency against musical standards
func Listener() ListenerResult {
	// Logic to determine if freq is flat/sharp/in-tune
	// will be an endless loop
	return ListenerResult{audio: "booo"}
}
