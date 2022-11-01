package shapes_test

import (
	"hello-shapes-go/shapes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Shapes", func() {
	var circle *shapes.Circle

	BeforeEach(func() {
		circle = &shapes.Circle{
			Radius: 3.0,
		}
	})

	Describe("Get areas of shapes", func() {
		Context("with a radius 3.0", func() {
			It("should be a Triangle", func() {
				Expect(circle.Area()).To(Equal(9.42477796076938))
			})
		})
	})

})
