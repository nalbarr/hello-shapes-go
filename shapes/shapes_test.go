package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var square Square
var circle Circle
var triangle Triangle

// NA. Not TDD testable
// var rectangle shapes.rectangle

func init() {
}

func TestSquareArea(t *testing.T) {
	square := Square{Side: 4.0}
	actual := square.Area()
	assert.Equal(t, actual, 16.0)
}

func TestCircleArea(t *testing.T) {
	circle := Circle{Radius: 3.0}
	actual := circle.Area()
	assert.Equal(t, actual, 9.42477796076938)
}

func TestTriangleArea(t *testing.T) {
	triangle := Triangle{Base: 3.0, Height: 5.0}
	actual := triangle.Area()
	assert.Equal(t, actual, 7.5)
}
