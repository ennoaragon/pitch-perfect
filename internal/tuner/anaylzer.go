package tuner

import "fmt"

// Result represents the state of a note
type Result struct {
    Note   string
    Status string // "In Tune", "Flat", "Sharp"
	Frequency float32
}

// Analyze checks a frequency against musical standards
// this will take in the processed audio and output the note
// also include if the note is flat or sharp or tune
func Analyze(freq float32) Result {
	analyzed := Result{Note: "A", Status: "Flat", Frequency: freq}

	fmt.Printf("Note: %s, Status: %s, Freq: %.4f", analyzed.Note, analyzed.Status, analyzed.Frequency)

	return  analyzed
    // Logic to determine if freq is flat/sharp/in-tune
}
