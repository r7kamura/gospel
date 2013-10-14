package gospel

import "strings"

// func It() creates a new Example object.
type Example struct {
	*ExampleGroup
	Message string
	Evaluator func()
	HasFailure bool
	Formatter
}

// func It() calls this function to run the newly created example.
func (example *Example) Run() {
	example.Started()
	example.Evaluate()
	if !example.HasFailure {
		example.Succeeded()
	}
	previousExample = example
}

// Invokes its Evaluator function.
func (example *Example) Evaluate() {
	example.Evaluator()
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
	example.Root().T.Fail()
	example.HasFailure = true
	example.Formatter.Failed(example, message)
}

// Returns its entire descriptions + message as a string.
func (example *Example) FullDescription() string {
	return strings.Join(append(example.Descriptions(), example.Message), " ")
}

// Returns its ExamplesGroups' descriptions as an array of string.
func (example *Example) Descriptions() []string {
	segments := []string{}
	for _, ancestor := range example.ExampleGroup.ReverseAncestorsAndSelf() {
		segments = append(segments, ancestor.Description)
	}
	return segments
}

