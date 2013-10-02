package gospel

import (
	"fmt"
	"testing"
)

// it(message, evaluator) will create this object.
type Example struct {
	T *testing.T
	Index uint
	Description string
	Message string
	Evaluator func(Expect)
}

// Run() is called as soon as it(message, evaluator) is called.
func (example *Example) Run() {
	example.Start()
	example.Evaluator(func(value interface{}) *Expectation {
		return &Expectation{example, value}
	})
}

// Start() is called before running this example.
func (example *Example) Start() {
	if example.Index == 0 {
		fmt.Println(example.Description)
	}
	fmt.Println("\t" + example.Message)
}

// Fail() is called when any of its expectation failed.
func (example *Example) Fail(message string, actual, expected interface{}) {
	example.T.Errorf("Expected `%v` to %s `%v`", expected, message, actual)
}

// Utility type for expect(...) function.
type Expect func(interface{}) *Expectation

// Utility type for it(...) function.
type It func(string, func(Expect))

// Please call this function from your TestXxx function.
func Describe(t *testing.T, description string, exampleGroup func(It)) {
	var index uint
	exampleGroup(func(message string, evaluator func(Expect)) {
		example := Example{t, index, description, message, evaluator}
		example.Run()
		index++
	})
}
