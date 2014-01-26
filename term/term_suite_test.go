package term_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"

    "testing"
)

func TestTerm(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Term Suite")
}
