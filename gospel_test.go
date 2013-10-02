package gospel

import "testing"

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
