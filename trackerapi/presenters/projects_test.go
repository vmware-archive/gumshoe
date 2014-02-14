package presenters_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi/domain"
    "github.com/pivotal/gumshoe/trackerapi/presenters"
)

var _ = Describe("Projects Presenter", func() {
    var (
        presenter presenters.Projects
    )

    BeforeEach(func() {
        presenter = presenters.Projects{
            Projects: []domain.Project{
                domain.Project{
                    ID:               123,
                    Name:             "Death Star",
                    CurrentIteration: 12,
                },
                domain.Project{
                    ID:               456,
                    Name:             "Rebel Alliance",
                    CurrentIteration: 4,
                },
            },
        }
    })

    It("returns a string representation of a project collection", func() {
        output := presenter.String()
        Expect(output).To(ContainSubstring("Death Star"))
        Expect(output).To(ContainSubstring("Rebel Alliance"))
    })
})
