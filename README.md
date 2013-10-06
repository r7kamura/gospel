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
	Describe(t, "gospel.Expectation#ToEqual", func() {
		Context("with 1 & 1", func() {
			It("compares integers by ==", func() {
				Expect(1).To(Equal, 1)
			})
		})
		Context("with `1` & `1`", func() {
			It("compares strings by ==", func() {
				Expect("1").To(Equal, "1")
			})
		})
	})

	Describe(t, "gospel.Expectation#ToNotEqual", func() {
		Context("with 1 & 2", func() {
			It("compares integers by !=", func() {
				Expect(1).To(NotEqual, 2)
			})
		})
		Context("with `1` & `2`", func() {
			It("compares strings by !=", func() {
				Expect(1).To(NotEqual, "2")
			})
		})
	})
}
```

```
$ go test
```

![](http://dl.dropboxusercontent.com//u/5978869/image/20131003_083821.png)

```
$ go test -v
```

![](http://dl.dropboxusercontent.com//u/5978869/image/20131003_083718.png)
