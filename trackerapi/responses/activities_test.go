package responses_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi/responses"
)

var _ = Describe("Activities Response", func() {
    var (
        response responses.Activities
    )

    BeforeEach(func() {
        response = responses.Activities{}
        responseBody := `[
            {"message": "Wilhuff Tarkin changed iteration 1's length from default to 2 weeks"},
            {"message": "Hey Girl, I'm just like coffee. I'll keep you up all night."}
        ]`
        response.Parse([]byte(responseBody))
    })

    It("can be used to unmarshal a response, returning activity domain objects", func() {
        activities := response.Activities()
        Expect(len(activities)).To(Equal(2))
        Expect(activities[0].Message).To(Equal("Wilhuff Tarkin changed iteration 1's length from default to 2 weeks"))
        Expect(activities[1].Message).To(Equal("Hey Girl, I'm just like coffee. I'll keep you up all night."))
    })
})
