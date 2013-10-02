package gospel

import "testing"

func TestDescribe(t *testing.T) {
	Describe(t, "gospel.Expectation#ToEqual", func(context Context, it It) {
		context("with 1 & 1", func() {
			it("compares integers with ==", func(expect Expect) {
				expect(1).ToEqual(1)
			})
		})
		context("with `1` & `1`", func() {
			it("compares strings with ==", func(expect Expect) {
				expect("1").ToEqual("1")
			})
		})
	})

	Describe(t, "gospel.Expectation#ToNotEqual", func(context Context, it It) {
		context("with 1 & 2", func() {
			it("compares integers with !=", func(expect Expect) {
				expect(1).ToNotEqual(2)
			})
		})
		context("with `1` & `2`", func() {
			it("compares strings with !=", func(expect Expect) {
				expect("1").ToNotEqual("2")
			})
		})
	})
}
