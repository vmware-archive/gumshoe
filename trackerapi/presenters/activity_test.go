package presenters_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi/domain"
    "github.com/pivotal/gumshoe/trackerapi/presenters"
)

var _ = Describe("Activity Presenter", func() {
    var (
        presenter presenters.Activity
    )

    BeforeEach(func() {
        activity := domain.Activity{
            Message: "That's no moon, it's a space station.",
        }
        presenter = presenters.Activity{
            Activity: activity,
        }
    })

    It("returns a string representation of an activity", func() {
        output := presenter.String()
        Expect(output).To(ContainSubstring("That's no moon, it's a space station."))
    })
})
