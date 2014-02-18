package responses_test

import (
    "encoding/json"

    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi/responses"
)

var _ = Describe("Projects Response", func() {
    var (
        response responses.Projects
    )

    BeforeEach(func() {
        response = responses.Projects{}
        responseBody := `[
           {
               "id": 98,
               "current_iteration_number": 1,
               "name": "Learn About the Force"
           },
           {
               "id": 99,
               "description": "Expeditionary Battle Planetoid",
               "current_iteration_number": 15,
               "name": "Death Star"
           }
        ]`
        json.Unmarshal([]byte(responseBody), &response.Structure)
    })

    It("can be used to unmarshal a response, returning project domain objects", func() {
        projects := response.Projects()
        Expect(len(projects)).To(Equal(2))

        Expect(projects[0].ID).To(Equal(98))
        Expect(projects[0].Name).To(Equal("Learn About the Force"))
        Expect(projects[0].CurrentIteration).To(Equal(1))

        Expect(projects[1].ID).To(Equal(99))
        Expect(projects[1].Name).To(Equal("Death Star"))
        Expect(projects[1].CurrentIteration).To(Equal(15))
        Expect(projects[1].Description).To(Equal("Expeditionary Battle Planetoid"))
    })
})
