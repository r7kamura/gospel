package gospel

import "fmt"

// func Expect() creates a new Expectation object.
type Expectation struct {
	*Example
	Actual interface{}
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

// Checks if actual == expected.
func Equal(values ...interface{}) (failureMessage string) {
	if values[0] != values[1] {
		failureMessage = fmt.Sprintf("Expected `%v` to equal `%v`", values[0], values[1])
	}
	return
}

// Checks if actual != expected.
func NotEqual(values ...interface{}) (failureMessage string) {
	if values[0] == values[1] {
		failureMessage = fmt.Sprintf("Expected `%v` to not equal `%v`", values[0], values[1])
	}
	return
}

// Checks if actual != nil.
func Exist(values ...interface{}) (failureMessage string) {
	return NotEqual(values[0], nil)
}

// Checks if actual == nil.
func NotExist(values ...interface{}) (failureMessage string) {
	return Equal(values[0], nil)
}
