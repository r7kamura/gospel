package gospel

import (
	"fmt"
	"reflect"
)

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

func isNil(value interface{}) bool {
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return val.IsNil()
	}
	return value == nil
}

// Checks if actual != nil.
func Exist(values ...interface{}) (failureMessage string) {
	if isNil(values[0]) {
		failureMessage = fmt.Sprintf("Expected `%v` is not exist.", values[0])
	}
	return
}

// Checks if actual == nil.
func NotExist(values ...interface{}) (failureMessage string) {
	if !isNil(values[0]) {
		failureMessage = fmt.Sprintf("Expected `%v` is exist.", values[0])
	}
	return
}

// Checks deep equailty of actual and expected.
func Same(values ...interface{}) (failureMessage string) {
	if !reflect.DeepEqual(values[0], values[1]) {
		failureMessage = fmt.Sprintf("Expected `%v` to equal `%v`", values[0], values[1])
	}
	return
}

func NotSame(values ...interface{}) (failureMessage string) {
	if reflect.DeepEqual(values[0], values[1]) {
		failureMessage = fmt.Sprintf("Expected `%v` to not equal `%v`", values[0], values[1])
	}
	return
}
