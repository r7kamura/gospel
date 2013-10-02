package gospel

// expect(actual) will create this object.
type Expectation struct {
	*Example
	Actual interface{}
}

// Matcher method: actual == expected.
func (expectation *Expectation) ToEqual(expected interface{}) {
	expectation.To("equal", Equal, expected)
}

// Matcher method: actual != expected.
func (expectation *Expectation) ToNotEqual(expected interface{}) {
	expectation.To("not equal", NotEqual, expected)
}

// Matcher method: actual != nil.
func (expectation *Expectation) ToExist() {
	expectation.To("exist", NotEqual, nil)
}

// Matcher method: actual == nil.
func (expectation *Expectation) ToNotExist() {
	expectation.To("not exist", Equal, nil)
}

// All-purpose matcher method to compare values with a given `matcher`.
func (expectation *Expectation) To(message string, matcher Matcher, expected interface{}) {
	if !expectation.Example.HasFailure {
		if !matcher(expectation.Actual, expected) {
			expectation.Example.Failed(message, expectation.Actual, expected)
		}
	}
}

// Utility type to define matcher function's form.
type Matcher func(interface{}, interface{}) bool

// For To(...) function.
func Equal(actual, expected interface{}) bool {
	return actual == expected
}

// For To(...) function.
func NotEqual(actual, expected interface{}) bool {
	return actual != expected
}
