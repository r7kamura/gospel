package gospel

import "testing"

var (
	currentExampleGroup *ExampleGroup
	currentExample *Example
	previousExample *Example
)

func Describe(t *testing.T, description string, callback func()) {
	currentExampleGroup = &ExampleGroup{
		T: t,
		Description: description,
		BeforeFilters: []func(){},
	}
	callback()
	currentExampleGroup.PrintResult()
	currentExampleGroup = nil
}

func Context(description string, callback func()) {
	currentExampleGroup = &ExampleGroup{
		Description: description,
		Parent: currentExampleGroup,
		BeforeFilters: []func(){},
	}
	callback()
	currentExampleGroup = currentExampleGroup.Parent
}

func It(message string, evaluator func()) {
	currentExampleGroup.RunBeforeFilters()
	currentExample = &Example{
		ExampleGroup: currentExampleGroup,
		Message: message,
		Evaluator: evaluator,
		Formatter: newFormatter(),
	}
	currentExample.Run()
	currentExampleGroup.RunAfterFilters()
}

func Before(filter func()) {
	currentExampleGroup.BeforeFilters = append(currentExampleGroup.BeforeFilters, filter)
}

func After(filter func()) {
	currentExampleGroup.AfterFilters = append(currentExampleGroup.AfterFilters, filter)
}

func Expect(actual interface{}) *Expectation {
	return &Expectation{
		Actual: actual,
		Example: currentExample,
	}
}
