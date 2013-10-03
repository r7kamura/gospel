package gospel

import "testing"

func TestDescribe(t *testing.T) {
	Describe(t, "gospel.Expectation#ToEqual", func() {
		Context("with 1 & 1", func() {
			It("compares integers by ==", func() {
				Expect(1).ToEqual(2)
				Expect(1).ToEqual(1)
			})
		})
		Context("with `1` & `1`", func() {
			It("compares strings by ==", func() {
				Expect("1").ToEqual("1")
			})
		})
	})

	Describe(t, "gospel.Expectation#ToNotEqual", func() {
		Context("with 1 & 2", func() {
			It("compares integers by !=", func() {
				Expect(1).ToNotEqual(2)
			})
		})
		Context("with `1` & `2`", func() {
			It("compares strings by !=", func() {
				Expect("1").ToNotEqual("2")
			})
		})
	})

	Describe(t, "gospel.Expectation#ToExist", func() {
		Context("with 1", func() {
			It("checks existence by non-equivalence with nil", func() {
				Expect(1).ToExist()
			})
		})
		Context("with `1`", func() {
			It("checks existence by non-equivalence with nil", func() {
				Expect("1").ToExist()
			})
		})
	})

	Describe(t, "gospel.Expectation#ToNotExist", func() {
		Context("with nil", func() {
			It("checks existence by equivalence with nil", func() {
				Expect(nil).ToNotExist()
			})
		})
	})
}
