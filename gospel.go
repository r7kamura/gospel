package gospel

import "testing"

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
		example := Example{
			Describing: describing,
			Message: message,
			Evaluator: evaluator,
		}
		example.Run()
	}
	context := func(subDescription string, contextCallback func()) {
		describing.SubDescriptions = append(describing.SubDescriptions, subDescription)
		contextCallback()
		describing.SubDescriptions = describing.SubDescriptions[:len(describing.SubDescriptions) - 1]
	}
	describeCallback(context, it)
}
