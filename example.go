package gospel

import "strings"

// func It() creates a new Example object.
type Example struct {
	*Describing
	Message string
	Evaluator func()
	HasFailure bool
	Formatter
}

// func It() calls this function to run the newly created example.
func (example *Example) Run() {
	example.Started()
	example.Evaluate()
	example.DoneExamplesCount++
	if !example.HasFailure {
		example.Succeeded()
	}
	example.UpdatePreviousSubDescriptions()
}

// Invokes its Evaluator function.
func (example *Example) Evaluate() {
	example.Evaluator()
}

// Copy an array and then assign its slice.
func (example *Example) UpdatePreviousSubDescriptions() {
	example.PreviousSubDescriptions = append(make([]string, 0), example.SubDescriptions...)
}

// Called when started.
func (example *Example) Started() {
	example.Formatter.Started(example)
}

// Called when all of expectations passed.
func (example *Example) Succeeded() {
	example.Formatter.Succeeded(example)
}

// Called when any of expectations failed.
func (example *Example) Failed(message string) {
	example.T.Fail()
	example.HasFailure = true
	example.Formatter.Failed(example, message)
}

// Returns its entire descriptions + message as a string.
func (example *Example) FullDescription() string {
	var segments []string
	segments = append(segments, example.Description)
	segments = append(segments, example.SubDescriptions...)
	segments = append(segments, example.Message)
	return strings.Join(segments, " ")
}
