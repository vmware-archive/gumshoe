package presenters_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi/domain"
    "github.com/pivotal/gumshoe/trackerapi/presenters"
)

var _ = Describe("Activities Presenter", func() {
    var (
        presenter presenters.Activities
    )

    BeforeEach(func() {
        activities := []domain.Activity{
            domain.Activity{
                Message: "That's no moon, it's a space station.",
            },
            domain.Activity{
                Message: "He's the brains, sweetheart!",
            },
        }
        presenter = presenters.Activities{
            Activities: activities,
        }
    })

    It("returns a string representation of a collection of activities", func() {
        output := presenter.String()
        Expect(output).To(ContainSubstring("Activity:\n"))
        Expect(output).To(ContainSubstring("  That's no moon, it's a space station."))
        Expect(output).To(ContainSubstring("  He's the brains, sweetheart!"))
    })
})
