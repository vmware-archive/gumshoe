package reporters_test

import (
	. "github.com/pivotal/gumshoe/repos/ginkgo"
	. "github.com/pivotal/gumshoe/repos/gomega"

	"testing"
)

func TestReporters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reporters Suite")
}
