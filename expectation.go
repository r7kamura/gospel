package gospel

import "fmt"

// func Expect() creates a new Expectation object.
type Expectation struct {
	*Example
	Actual interface{}
}

// Checks if actual == expected.
func (expectation *Expectation) ToEqual(expected interface{}) {
	expectation.To(Equal, expected)
}

// Checks if actual != expected.
func (expectation *Expectation) ToNotEqual(expected interface{}) {
	expectation.To(NotEqual, expected)
}

// Checks if actual != nil.
func (expectation *Expectation) ToExist() {
	expectation.To(NotEqual, nil)
}

// Checks if actual == nil.
func (expectation *Expectation) ToNotExist() {
	expectation.To(Equal, nil)
}

// Checks if a given matcher satisfies actual & optional values.
func (expectation *Expectation) To(matcher Matcher, expected ...interface{}) {
	if !expectation.Example.HasFailure {
		values := append([]interface{}{expectation.Actual}, expected...)
		if failureMessage := matcher(values...); failureMessage != "" {
			expectation.Example.Failed(failureMessage)
		}
	}
}

// Returns a non-empty failure message if given actual & optional values are not matched.
type Matcher func(...interface{}) string

func Equal(values ...interface{}) (failureMessage string) {
	if values[0] != values[1] {
		failureMessage = fmt.Sprintf("Expected `%v` to equal `%v`", values[0], values[1])
	}
	return
}

func NotEqual(values ...interface{}) (failureMessage string) {
	if values[0] == values[1] {
		failureMessage = fmt.Sprintf("Expected `%v` to not equal `%v`", values[0], values[1])
	}
	return
}
