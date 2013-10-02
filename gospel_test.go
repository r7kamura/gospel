package gospel

import "testing"

func TestDescribe(t *testing.T) {
	Describe(t, "gospel.Expectation#ToEqual", func(context Context, it It) {
		context("with 1 & 1", func() {
			it("compares integers by ==", func(expect Expect) {
				expect(1).ToEqual(1)
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

	Describe(t, "gospel.Expectation#ToExist", func(context Context, it It) {
		context("with 1", func() {
			it("checks existence by non-equivalence with nil", func(expect Expect) {
				expect(1).ToExist()
			})
		})
		context("with `1`", func() {
			it("checks existence by non-equivalence with nil", func(expect Expect) {
				expect("1").ToExist()
			})
		})
	})

	Describe(t, "gospel.Expectation#ToNotExist", func(context Context, it It) {
		context("with nil", func() {
			it("checks existence by equivalence with nil", func(expect Expect) {
				expect(nil).ToNotExist()
			})
		})
	})
}
