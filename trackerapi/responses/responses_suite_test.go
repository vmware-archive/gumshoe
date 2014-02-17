package responses_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"

    "testing"
)

func TestResponsess(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Responses Suite")
}
