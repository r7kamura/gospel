package gospel

import (
	"fmt"
	"strings"
)

// it(message, evaluator) will create this object.
type Example struct {
	*Describing
	Message string
	Evaluator func(Expect)
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
	differenceIsFound := false
	for i, subscription := range example.SubDescriptions {
		if !differenceIsFound && i <= len(example.PreviousSubDescriptions) - 1 {
			if subscription == example.PreviousSubDescriptions[i] {
				continue
			}
		}
		fullMessage += strings.Repeat("\t", i + 1) + subscription
		differenceIsFound = true
	}
	fmt.Println(fullMessage)
	example.PreviousSubDescriptions = append(make([]string, 0), example.SubDescriptions...)
}

// Done() is called at end of this example.
func (example *Example) Done() {
	example.DoneExamplesCount++
}

// Succeeded() is called when any of its expectation failed.
func (example *Example) Succeeded(message string, actual, expected interface{}) {
	fmt.Println(green(example.MessageLine()))
}

// Failed() is called when any of its expectation failed.
func (example *Example) Failed(message string, actual, expected interface{}) {
	fmt.Println(red(example.MessageLine()))
	example.Describing.T.Errorf("Expected `%v` to %s `%v`", expected, message, actual)
}

// Utility method to return message line with proper tabs.
func (example *Example) MessageLine() string {
	return strings.Repeat("\t", len(example.SubDescriptions) + 1) + example.Message
}

// Add red terminal ANSI color
func red(str string) string {
	return "\033[31m\033[1m" + str + "\033[0m"
}

// Add green terminal ANSI color
func green(str string) string {
	return "\033[32m\033[1m" + str + "\033[0m"
}
