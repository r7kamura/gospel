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
	Describe(t, "gospel.Expectation#ToEqual", func(context Context, it It) {
		context("with 1 & 1", func() {
			it("compares integers by ==", func(expect Expect) {
				expect(1).ToEqual(1)
			})
		})
		context("with `1` & `1`", func() {
			it("compares strings by ==", func(expect Expect) {
				expect("1").ToEqual("1")
			})
		})
	})

	Describe(t, "gospel.Expectation#ToNotEqual", func(context Context, it It) {
		context("with 1 & 2", func() {
			it("compares integers by !=", func(expect Expect) {
				expect(1).ToNotEqual(2)
			})
		})
		context("with `1` & `2`", func() {
			it("compares strings by !=", func(expect Expect) {
				expect("1").ToNotEqual("2")
			})
		})
	})
}
```

```
$ go test
```

![](http://dl.dropboxusercontent.com//u/5978869/image/20131003_025631.png)
