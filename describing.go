package gospel

import (
	"fmt"
	"testing"
)

// func Describe() function creates a new Describing object.
type Describing struct {
	*testing.T
	DoneExamplesCount uint
	Description string
	SubDescriptions []string
	PreviousSubDescriptions []string
	Result string
}

// Called from Context().
func (describing *Describing) PushSubDescription(subDescription string) {
	describing.SubDescriptions = append(describing.SubDescriptions, subDescription)
}

// Called from Context().
func (describing *Describing) PopSubDescription() {
	describing.SubDescriptions = describing.SubDescriptions[:len(describing.SubDescriptions) - 1]
}

// Prints Describing.Result value to show failures in non verbose mode.
func (describing *Describing) PrintResult() {
	fmt.Print(describing.Result)
}

// Shares a current running Describing object.
var currentDescribing *Describing

// Shares a current running Example object.
var currentExample *Example

// Creates and invokes a new Describing object.
func Describe(t *testing.T, description string, callback func()) {
	currentDescribing = &Describing{
		T: t,
		Description: description,
		SubDescriptions: []string{},
		PreviousSubDescriptions: []string{},
	}
	callback()
	currentDescribing.PrintResult()
}

// Nests example description.
func Context(subDescription string, callback func()) {
	currentDescribing.PushSubDescription(subDescription)
	callback()
	currentDescribing.PopSubDescription()
}

// Creates and invokes a new Example object.
func It(message string, evaluator func()) {
	currentExample = &Example{
		Describing: currentDescribing,
		Message: message,
		Evaluator: evaluator,
		Formatter: newFormatter(),
	}
	currentExample.Run()
}

// Creates a new Expectation object to test a given value.
func Expect(actual interface{}) *Expectation {
	return &Expectation{
		Actual: actual,
		Example: currentExample,
	}
}
