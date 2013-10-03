package gospel

import (
	"fmt"
	"os"
	"testing"
)

// Decides Formatter type.
var verboseMode bool

// Checks if -v option specified or not.
func init() {
	for _, argument := range os.Args {
		if argument == "-test.v=true" {
			verboseMode = true
			return
		}
	}
}

// Factory method to create a Formatter.
func newFormatter() Formatter {
	if verboseMode {
		return &DocumentFormatter{}
	} else {
		return &DotFormatter{}
	}
}

// Describe(...) will create this object to manage its examples.
type Describing struct {
	*testing.T
	DoneExamplesCount uint
	Description string
	SubDescriptions []string
	PreviousSubDescriptions []string
	Result string
}

// Print describing's Result value. This value is set if any example failed without -v option.
func (describing *Describing) PrintResult() {
	fmt.Print(describing.Result)
}

// Share a current running Describing object.
var currentDescribing *Describing

// Share a current running Example object.
var currentExample *Example

// Please call Describe(...) from your TestXxx function.
func Describe(t *testing.T, description string, describeCallback func()) {
	currentDescribing = &Describing{
		T: t,
		Description: description,
		SubDescriptions: make([]string, 0),
		PreviousSubDescriptions: make([]string, 0),
	}
	describeCallback()
	currentDescribing.PrintResult()
}

func Context(subDescription string, contextCallback func()) {
	currentDescribing.SubDescriptions = append(currentDescribing.SubDescriptions, subDescription)
	contextCallback()
	currentDescribing.SubDescriptions = currentDescribing.SubDescriptions[:len(currentDescribing.SubDescriptions) - 1]
}

func It(message string, evaluator func()) {
	currentExample = &Example{
		Describing: currentDescribing,
		Message: message,
		Evaluator: evaluator,
		Formatter: newFormatter(),
	}
	currentExample.Run()
}

func Expect(actual interface{}) *Expectation {
	return &Expectation{
		Actual: actual,
		Example: currentExample,
	}
}
