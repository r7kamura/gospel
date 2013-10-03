package gospel

// func Expect() creates a new Expectation object.
type Expectation struct {
	*Example
	Actual interface{}
}

// Checks if actual == expected.
func (expectation *Expectation) ToEqual(expected interface{}) {
	expectation.To("equal", Equal, expected)
}

// Checks if actual != expected.
func (expectation *Expectation) ToNotEqual(expected interface{}) {
	expectation.To("not equal", NotEqual, expected)
}

// Checks if actual != nil.
func (expectation *Expectation) ToExist() {
	expectation.To("not equal", NotEqual, nil)
}

// Checks if actual == nil.
func (expectation *Expectation) ToNotExist() {
	expectation.To("equal", Equal, nil)
}

// Checks if matcher(actual, expected) == true only if all of previous expectations passed.
func (expectation *Expectation) To(message string, matcher Matcher, expected interface{}) {
	if !expectation.Example.HasFailure {
		if !matcher(expectation.Actual, expected) {
			expectation.Example.Failed(message, expectation.Actual, expected)
		}
	}
}

// Utility type to define a matcher function.
type Matcher func(interface{}, interface{}) bool

// For To(...) function.
func Equal(actual, expected interface{}) bool {
	return actual == expected
}

// For To(...) function.
func NotEqual(actual, expected interface{}) bool {
	return actual != expected
}
