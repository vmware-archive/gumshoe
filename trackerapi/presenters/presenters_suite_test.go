package presenters_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"

    "testing"
)

func TestPresenters(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Presenters Suite")
}
