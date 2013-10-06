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
	Describe(t, "Expectation#To", func() {
		Context("with Equal", func() {
			It("evaluates actual == expected", func() {
				Expect(1).To(Equal, 1)
			})
		})

		Context("with NotEqual", func() {
			It("evaluates actual != expected", func() {
				Expect(1).To(NotEqual, 2)
			})
		})

		Context("with Exist", func() {
			It("evaluates actual != nil", func() {
				Expect(1).To(Exist)
			})
		})

		Context("with NotExist", func() {
			It("evaluates actual == nil", func() {
				Expect(nil).To(NotExist)
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

![](http://dl.dropboxusercontent.com//u/5978869/image/20131006_224123.png)
