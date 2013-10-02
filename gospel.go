package gospel

import (
	"fmt"
	"strings"
	"testing"
)

// it(message, evaluator) will create this object.
type Example struct {
	*Describing
	Message string
	Evaluator func(Expect)
}

// Run() is called as soon as it(message, evaluator) is called.
func (example *Example) Run() {
	example.Started()
	example.Evaluate()
	example.Done()
}

// Evaluate() invokes its Evaluator function.
func (example *Example) Evaluate() {
	example.Evaluator(func(value interface{}) *Expectation {
		return &Expectation{example, value}
	})
}

// Started() is called at start of this example, printing doc-style example message.
func (example *Example) Started() {
	fullMessage := ""
	if example.DoneExamplesCount == 0 {
		fullMessage += example.Description + "\n"
	}
	differenceIsFound := false
	for i, subscription := range example.SubDescriptions {
		if !differenceIsFound && i <= len(example.PreviousSubDescriptions) - 1 {
			if subscription == example.PreviousSubDescriptions[i] {
				continue
			}
		}
		fullMessage += strings.Repeat("\t", i + 1) + subscription + "\n"
		differenceIsFound = true
	}
	fullMessage += strings.Repeat("\t", len(example.SubDescriptions) + 1) + example.Message
	fmt.Println(fullMessage)
	example.PreviousSubDescriptions = append(make([]string, 0), example.SubDescriptions...)
}

// Done() is called at end of this example.
func (example *Example) Done() {
	example.DoneExamplesCount++
}

// Failed() is called when any of its expectation failed.
func (example *Example) Failed(message string, actual, expected interface{}) {
	example.Describing.T.Errorf("Expected `%v` to %s `%v`", expected, message, actual)
}

// Utility type for context(...) function.
type Context func(string, func())

// Utility type for expect(...) function.
type Expect func(interface{}) *Expectation

// Utility type for it(...) function.
type It func(string, func(Expect))

// Describe(...) will create this object to manage its examples.
type Describing struct {
	*testing.T
	DoneExamplesCount uint
	Description string
	SubDescriptions []string
	PreviousSubDescriptions []string
}

// Please call Describe(...) from your TestXxx function.
func Describe(t *testing.T, description string, describeCallback func(Context, It)) {
	describing := &Describing{
		T: t,
		Description: description,
		SubDescriptions: make([]string, 0),
		PreviousSubDescriptions: make([]string, 0),
	}
	it := func(message string, evaluator func(Expect)) {
		example := Example{describing, message, evaluator}
		example.Run()
	}
	context := func(subDescription string, contextCallback func()) {
		describing.SubDescriptions = append(describing.SubDescriptions, subDescription)
		contextCallback()
		describing.SubDescriptions = describing.SubDescriptions[:len(describing.SubDescriptions) - 1]
	}
	describeCallback(context, it)
}
