package shapes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestShapes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shapes Suite")
}
