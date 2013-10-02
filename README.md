# Gospel
BDD-style testing library for Golang.

## Install
```
$ go get github.com/r7kamura/gospel
```

## Usage
```go
package main

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestDescribe(t *testing.T) {
	Describe(t, "gospel.Expectation#ToEqual", func(it It) {
		it("compares integers with ==", func(expect Expect) {
			expect(1).ToEqual(1)
		})
		it("compares strings with ==", func(expect Expect) {
			expect("1").ToEqual("1")
		})
	})

	Describe(t, "gospel.Expectation#ToNotEqual", func(it It) {
		it("compares integers with !=", func(expect Expect) {
			expect(1).ToNotEqual(2)
		})
		it("compares strings with !=", func(expect Expect) {
			expect("1").ToNotEqual("2")
		})
	})
}
```

```
$ go test
gospel.Expectation#ToEqual
        compares integers with ==
        compares strings with ==
gospel.Expectation#ToNotEqual
        compares integers with !=
        compares strings with !=
PASS
ok      _/Users/r7kamura/gospel     0.023s
```
