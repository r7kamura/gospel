package gospel

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/mattn/go-colorable"
)

// Decides Formatter type.
var verboseMode bool
var Output = colorable.NewColorableStdout()

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

type Formatter interface {
	Started(*Example)
	Failed(*Example, string)
	Succeeded(*Example)
}

type DotFormatter struct {}

// Does nothing.
func (formatter *DotFormatter) Started(example *Example) {
}

func (formatter *DotFormatter) Failed(example *Example, message string) {
	fmt.Fprintf(Output, red("F"))
	root := example.ExampleGroup.Root()
	if root.Result == "" {
		root.Result += "\n\n"
	}
	_, filename, line, _ := runtime.Caller(3)
	buffer, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(buffer), "\n")[line-2:line+2]
	root.Result += fmt.Sprintf(
		red("  %s\n") +
		grey("  %s\n") +
		grey("  %s:%d\n") +
		grey("  %4d.%s\n") +
		grey("  %4d.%s\n") +
		grey("  %4d.%s\n") +
		"\n",
		example.FullDescription(),
		message,
		filename, line,
		line - 1, strings.Replace(lines[0], "\t", "  ", -1),
		line + 0, strings.Replace(lines[1], "\t", "  ", -1),
		line + 1, strings.Replace(lines[2], "\t", "  ", -1),
	)
}

func (formatter *DotFormatter) Succeeded(example *Example) {
	fmt.Fprintf(Output, green("."))
}

type DocumentFormatter struct {}

func (formatter *DocumentFormatter) Started(example *Example) {
	var previousDescriptions []string
	if previousExample != nil {
		previousDescriptions = previousExample.Descriptions()
	}
	currentDescriptions := example.Descriptions()
	fullMessage := ""
	for i, description := range currentDescriptions {
		if previousExample == nil || i > len(previousDescriptions) - 1 || description != previousDescriptions[i] {
			fullMessage += strings.Repeat("  ", i) + "\033[0m" + description + "\n"
		}
	}
	fmt.Fprint(Output, fullMessage)
}

func (formatter *DocumentFormatter) Succeeded(example *Example) {
	margin := strings.Repeat("  ", len(example.ExampleGroup.Ancestors()) + 1)
	fmt.Fprintln(Output, margin + green(example.Message))
}

func (formatter *DocumentFormatter) Failed(example *Example, message string) {
	_, filename, line, _ := runtime.Caller(3)
	buffer, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(buffer), "\n")[line-2:line+2]
	margin := strings.Repeat("  ", len(example.ExampleGroup.Ancestors()) + 1)
	fmt.Fprint(Output,
		red("%s%s\n") +
		grey("%s%s\n") +
		grey("%s%s:%d\n") +
		grey("%s%4d.%s\n") +
		grey("%s%4d.%s\n") +
		grey("%s%4d.%s\n"),
		margin, example.Message,
		margin, message,
		margin, filename, line,
		margin, line - 1, strings.Replace(lines[0], "\t", "  ", -1),
		margin, line + 0, strings.Replace(lines[1], "\t", "  ", -1),
		margin, line + 1, strings.Replace(lines[2], "\t", "  ", -1),
	)
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
