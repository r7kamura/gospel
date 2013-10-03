package gospel

import "testing"

func TestDescribe(t *testing.T) {
	Describe(t, "gospel.Expectation#ToEqual", func() {
		Context("wIth 1 & 1", func() {
			It("compares integers by ==", func() {
				Expect(1).ToEqual(2)
				Expect(1).ToEqual(1)
			})
		})
		Context("wIth `1` & `1`", func() {
			It("compares strings by ==", func() {
				Expect("1").ToEqual("1")
			})
		})
	})

	Describe(t, "gospel.Expectation#ToNotEqual", func() {
		Context("wIth 1 & 2", func() {
			It("compares integers by !=", func() {
				Expect(1).ToNotEqual(2)
			})
		})
		Context("wIth `1` & `2`", func() {
			It("compares strings by !=", func() {
				Expect("1").ToNotEqual("2")
			})
		})
	})

	Describe(t, "gospel.Expectation#ToExist", func() {
		Context("wIth 1", func() {
			It("checks existence by non-equivalence wIth nil", func() {
				Expect(1).ToExist()
			})
		})
		Context("wIth `1`", func() {
			It("checks existence by non-equivalence wIth nil", func() {
				Expect("1").ToExist()
			})
		})
	})

	Describe(t, "gospel.Expectation#ToNotExist", func() {
		Context("wIth nil", func() {
			It("checks existence by equivalence wIth nil", func() {
				Expect(nil).ToNotExist()
			})
		})
	})
}
