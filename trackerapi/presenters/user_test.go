package presenters_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi/domain"
    "github.com/pivotal/gumshoe/trackerapi/presenters"
)

var _ = Describe("User Presenter", func() {
    var (
        presenter presenters.User
    )

    BeforeEach(func() {
        user := domain.User{
            Name:     "Mister Tee",
            Username: "mister_tee",
            Email:    "mister_tee@example.com",
            APIToken: "abc-123",
            Timezone: "America/Los_Angeles",
            Initials: "T",
        }
        presenter = presenters.User{
            User: user,
        }
    })

    It("returns a string representation of a user", func() {
        output := presenter.String()
        Expect(output).To(ContainSubstring("Mister Tee (mister_tee)"))
        Expect(output).To(ContainSubstring("  Email     : mister_tee@example.com"))
        Expect(output).To(ContainSubstring("  API Token : abc-123"))
        Expect(output).To(ContainSubstring("  Timezone  : America/Los_Angeles"))
        Expect(output).To(ContainSubstring("  Initials  : T"))
    })
})
