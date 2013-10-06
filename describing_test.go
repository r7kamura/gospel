package gospel

import "testing"

func TestDescribe(t *testing.T) {
	Describe(t, "gospel.Expectation#ToEqual", func() {
		Context("with 1 & 1", func() {
			It("compares integers by ==", func() {
				Expect(1).To(Equal, 1)
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

	Describe(t, "gospel.Expectation#ToExist", func() {
		Context("with 1", func() {
			It("checks existence by non-equivalence with nil", func() {
				Expect(1).To(Exist)
			})
		})
		Context("with `1`", func() {
			It("checks existence by non-equivalence with nil", func() {
				Expect("1").To(Exist)
			})
		})
	})

	Describe(t, "gospel.Expectation#ToNotExist", func() {
		Context("with nil", func() {
			It("checks existence by equivalence with nil", func() {
				Expect(nil).To(NotExist)
			})
		})
	})
}
