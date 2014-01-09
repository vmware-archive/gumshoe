package types_test

import (
	. "github.com/pivotal/gumshoe/repos/ginkgo"
	. "github.com/pivotal/gumshoe/repos/gomega"

	"testing"
)

func TestTypes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo Types Suite")
}
