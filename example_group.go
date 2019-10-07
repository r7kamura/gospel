package gospel

import (
	"fmt"
	"testing"
)

// Describe or Context creates this object.
type ExampleGroup struct {
	*testing.T
	Parent *ExampleGroup
	Description string
	Result string
	BeforeFilters []func()
	AfterFilters []func()
}

// In non-verbose mode, we print test results for each Describe.
func (group *ExampleGroup) PrintResult() {
	fmt.Fprint(Output, group.Result)
}

// Returns its ancestor ExampleGroups in ascending order according to distance.
func (group *ExampleGroup) Ancestors() []*ExampleGroup {
	ancestors := []*ExampleGroup{}
	for parent := group.Parent; parent != nil; parent = parent.Parent {
		ancestors = append(ancestors, parent)
	}
	return ancestors
}

// Returns its ancestor ExampleGroups in descending order according to distance.
func (group *ExampleGroup) ReverseAncestors() []*ExampleGroup {
	reversed := []*ExampleGroup{}
	ancestors := group.Ancestors()
	for i := len(ancestors) - 1; i >= 0; i-- {
		reversed = append(reversed, ancestors[i])
	}
	return reversed
}

// Gets root ExampleGroup defined by Describe func.
func (group *ExampleGroup) Root() *ExampleGroup {
	if group.Parent == nil {
		return group
	}
	ancestors := group.Ancestors()
	return ancestors[len(ancestors) - 1]
}

// Runs all BeforeFilters defined to its ancestors and itself.
func (group *ExampleGroup) RunBeforeFilters() {
	for _, ancestor := range group.ReverseAncestorsAndSelf() {
		for _, beforeFilter := range ancestor.BeforeFilters {
			beforeFilter()
		}
	}
}

// Runs all AfterFilters defined to its ancestors and itself.
func (group *ExampleGroup) RunAfterFilters() {
	for _, ancestor := range group.ReverseAncestorsAndSelf() {
		for _, afterFilter := range ancestor.AfterFilters {
			afterFilter()
		}
	}
}

// Merges its ancestors and itself.
func (group *ExampleGroup) ReverseAncestorsAndSelf() []*ExampleGroup {
	return append(group.ReverseAncestors(), group)
}
