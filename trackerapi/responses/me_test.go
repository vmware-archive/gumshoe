package responses_test

import (
    "encoding/json"

    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi/responses"
)

var _ = Describe("Me Response", func() {
    var (
        response responses.Me
    )

    BeforeEach(func() {
        response = responses.Me{}
        responseBody := `{
            "api_token": "abcde90792f3898ab464cd3412345",
            "name": "Mister Tee",
            "kind": "me",
            "id": 123,
            "email": "mister_tee@pivotallabs.com",
            "initials": "MT",
            "username": "mister_tee",
            "time_zone": {
                "kind": "time_zone",
                "offset": "-08:00",
                "olson_name": "America/Los_Angeles"
            }
        }`
        json.Unmarshal([]byte(responseBody), &response.Structure)
    })

    It("can be used to unmarshal a response, returning a user domain object", func() {
        user := response.User()
        Expect(user.APIToken).To(Equal("abcde90792f3898ab464cd3412345"))
        Expect(user.Name).To(Equal("Mister Tee"))
        Expect(user.Email).To(Equal("mister_tee@pivotallabs.com"))
        Expect(user.Initials).To(Equal("MT"))
        Expect(user.Username).To(Equal("mister_tee"))
        Expect(user.Timezone).To(Equal("America/Los_Angeles"))
    })
})
