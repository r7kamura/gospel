package gospel

import "testing"

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
