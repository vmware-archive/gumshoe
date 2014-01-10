package trackerapi_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi"
)

var _ = Describe("Resolver", func() {

    Describe("ActivityRequestURL", func() {
        It("returns a string for the activity endpoint with the given project ID", func() {
            resolver := &trackerapi.Resolver{
                TrackerDomain: "http://localhost:9000",
            }
            Expect(resolver.ActivityRequestURL(345)).To(Equal("http://localhost:9000/projects/345/activity?limit=5"))
        })
    })

})
