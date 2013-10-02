package gospel

import "strings"

// it(message, evaluator) will create this object.
type Example struct {
	*Describing
	Message string
	Evaluator func(Expect)
	HasFailure bool
	Formatter
}

// Run() is called as soon as it(message, evaluator) is called.
func (example *Example) Run() {
	example.Started()
	example.Evaluate()
	example.DoneExamplesCount++
	if !example.HasFailure {
		example.Succeeded()
	}
	example.UpdatePreviousSubDescriptions()
}

// Evaluate() invokes its Evaluator function.
func (example *Example) Evaluate() {
	example.Evaluator(func(value interface{}) *Expectation {
		return &Expectation{example, value}
	})
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
func (example *Example) Failed(message string, actual, expected interface{}) {
	example.Describing.T.Fail()
	example.HasFailure = true
	example.Formatter.Failed(example, message, actual, expected)
}

// Returns its entire descriptions + message as a string.
func (example *Example) FullDescription() string {
	var segments []string
	segments = append(segments, example.Description)
	segments = append(segments, example.SubDescriptions...)
	segments = append(segments, example.Message)
	return strings.Join(segments, " ")
}
