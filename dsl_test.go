package gospel

import "testing"

type T struct {
	e0 int
	e1 []string
}

func TestDescribe(t *testing.T) {
	Describe(t, "Expectation#To", func() {
		Context("with Equal", func() {
			It("evaluates actual == expected", func() {
				Expect(1).To(Equal, 1)
			})
		})

		Context("with NotEqual", func() {
			Before(func() {
				// Called before each examples in this Context.
			})

			After(func() {
				// Called after each examples in this Context.
			})

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

		Context("with Same", func() {
			It("evaluates actual == expected", func() {
				t0 := &T{e0: 0, e1: []string{"a", "b", "c"}}
				t1 := &T{e0: 0, e1: []string{"a", "b", "c"}}
				Expect(t0).To(NotEqual, t1)
				Expect(t0).To(Same, t1)
			})
		})

		Context("with NotSame", func() {
			It("evaluates actual != expected", func() {
				t0 := &T{e0: 0, e1: []string{"a", "b", "c"}}
				t1 := &T{e0: 0, e1: []string{"a", "b"}}
				Expect(t0).To(NotEqual, t1)
				Expect(t0).To(NotSame, t1)
			})
		})
	})
}
