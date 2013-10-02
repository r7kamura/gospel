package gospel

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
)

// it(message, evaluator) will create this object.
type Example struct {
	*Describing
	Message string
	Evaluator func(Expect)
	HasFailure bool
}

// Run() is called as soon as it(message, evaluator) is called.
func (example *Example) Run() {
	example.Started()
	example.Evaluate()
	example.Done()
}

// Evaluate() invokes its Evaluator function.
func (example *Example) Evaluate() {
	example.Evaluator(func(value interface{}) *Expectation {
		return &Expectation{example, value}
	})
}

// Started() is called at start of this example, printing doc-style example message.
func (example *Example) Started() {
	fullMessage := ""
	if example.DoneExamplesCount == 0 {
		fullMessage += example.Description + "\n"
	}
	if len(example.SubDescriptions) > 0 {
		differenceIsFound := false
		for i, subscription := range example.SubDescriptions {
			if !differenceIsFound && i <= len(example.PreviousSubDescriptions) - 1 {
				if subscription == example.PreviousSubDescriptions[i] {
					continue
				}
			}
			fullMessage += strings.Repeat("  ", i + 1) + subscription + "\n"
			differenceIsFound = true
		}
	}
	fmt.Print(fullMessage)
	example.PreviousSubDescriptions = append(make([]string, 0), example.SubDescriptions...)
}

// Done() is called at end of this example.
func (example *Example) Done() {
	example.DoneExamplesCount++
	if !example.HasFailure {
		example.Succeeded()
	}
}

// Succeeded() is called when all of its expectations are passed.
func (example *Example) Succeeded() {
	fmt.Println(example.LeftMargin() + green(example.Message))
}

// Failed() is called when any of its expectation failed.
func (example *Example) Failed(message string, actual, expected interface{}) {
	_, filename, line, _ := runtime.Caller(3)
	buffer, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(buffer), "\n")[line-2:line+2]
	fmt.Printf(
		red("%s%s\n") +
		grey("%sExpected `%v` to %s `%v`\n") +
		grey("%s%s:%d\n") +
		grey("%s%4d.%s\n") +
		grey("%s%4d.%s\n") +
		grey("%s%4d.%s\n"),
		example.LeftMargin(), example.Message,
		example.LeftMargin(), actual, message, expected,
		example.LeftMargin(), filename, line,
		example.LeftMargin(), line - 1, lines[0],
		example.LeftMargin(), line + 0, lines[1],
		example.LeftMargin(), line + 1, lines[2],
	)
	example.Describing.T.Fail()
	example.HasFailure = true
}

// Utility method to put margin.
func (example *Example) LeftMargin() string {
	return strings.Repeat("  ", len(example.SubDescriptions) + 1)
}

// Add red terminal ANSI color
func red(str string) string {
	return "\033[31m\033[1m" + str + "\033[0m"
}

// Add green terminal ANSI color
func green(str string) string {
	return "\033[32m\033[1m" + str + "\033[0m"
}

// Add grey terminal ANSI color
func grey(str string) string {
	return "\x1B[90m" + str + "\033[0m"
}
