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
	Result string
}

func (describing *Describing) PrintResult() {
	fmt.Print(describing.Result)
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
			Formatter: newFormatter(),
		}
		example.Run()
	}
	context := func(subDescription string, contextCallback func()) {
		describing.SubDescriptions = append(describing.SubDescriptions, subDescription)
		contextCallback()
		describing.SubDescriptions = describing.SubDescriptions[:len(describing.SubDescriptions) - 1]
	}
	describeCallback(context, it)
	describing.PrintResult()
}
