package presenters_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi/domain"
    "github.com/pivotal/gumshoe/trackerapi/presenters"
)

var _ = Describe("Project Presenter", func() {
    var (
        presenter presenters.Project
    )

    BeforeEach(func() {
        project := domain.Project{
            ID:               123,
            Name:             "Death Star",
            CurrentIteration: 12,
        }
        presenter = presenters.Project{
            Project: project,
        }
    })

    It("returns a string representation of a project", func() {
        output := presenter.String()
        Expect(output).To(ContainSubstring("Death Star (123)"))
        Expect(output).To(ContainSubstring("  Current Iteration : 12"))

        Expect(output).NotTo(ContainSubstring("Expeditionary Battle Planetoid"))
        presenter.Project.Description = "Expeditionary Battle Planetoid"
        output = presenter.String()
        Expect(output).To(ContainSubstring("  Description       : Expeditionary Battle Planetoid"))
    })
})
