package reporters

import (
    "github.com/pivotal/gumshoe/repos/ginkgo/config"
    "github.com/pivotal/gumshoe/repos/ginkgo/types"
)

type Reporter interface {
    SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary)
    ExampleWillRun(exampleSummary *types.ExampleSummary)
    ExampleDidComplete(exampleSummary *types.ExampleSummary)
    SpecSuiteDidEnd(summary *types.SuiteSummary)
}
