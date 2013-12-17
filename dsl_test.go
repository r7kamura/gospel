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

			It("evaluates actual == expected with struct", func() {
				t0 := &T{e0: 0, e1: []string{"a", "b", "c"}}
				t1 := &T{e0: 0, e1: []string{"a", "b", "c"}}
				Expect(t0).To(Equal, t1)
			})

			It("evaluates actual == expected with array", func() {
				Expect([2]string{"x", "y"}).To(Equal, [2]string{"x", "y"})
			})

			It("evaluates actual == expected with slice", func() {
				Expect([]string{"x", "y"}).To(Equal, []string{"x", "y"})
			})

			It("evaluates actual == expected with map", func() {
				m0 := map[string]string{"x": "foo", "y": "bar"}
				m1 := map[string]string{"x": "foo", "y": "bar"}
				Expect(m0).To(Equal, m1)
			})

			It("evaluates actual == expected with nil", func() {
				Expect(nil).To(Equal, nil)
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

			It("evaluates actual != expected with struct", func() {
				t0 := &T{e0: 0, e1: []string{"a", "b", "c"}}
				t1 := &T{e0: 0, e1: []string{"a", "b"}}
				Expect(t0).To(NotEqual, t1)
			})

			It("evaluates actual != expected with array", func() {
				Expect([2]string{"x", "y"}).To(NotEqual, [2]string{"x", "z"})
			})

			It("evaluates actual != expected with slice", func() {
				Expect([]string{"x", "y"}).To(NotEqual, []string{"x", "z"})
			})

			It("evaluates actual != expected with map", func() {
				m0 := map[string]string{"x": "foo", "y": "bar"}
				m1 := map[string]string{"x": "foo", "y": "baz"}
				Expect(m0).To(NotEqual, m1)
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
