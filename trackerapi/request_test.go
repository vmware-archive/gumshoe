package trackerapi_test

import (
    "strings"

    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi"
)

var _ = Describe("Request#Execute", func() {
    var (
        json     string
        client   *trackerapi.Client
        ts       *TestServer
        resolver *trackerapi.Resolver
        token    string
    )

    BeforeEach(func() {
        token = "abcde90792f3898ab464cd3412345"
        store := trackerapi.NewStore()
        store.Set("APIToken", token)
        client, _ = trackerapi.NewClient()
        json = `{
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
        ts = &TestServer{
            apiToken: token,
        }
        ts.Boot()
        ts.SetResponse("/me", json)
        resolver = &trackerapi.Resolver{
            TrackerDomain: ts.URL,
        }
        client.SetResolver(resolver)
    })

    AfterEach(func() {
        client.Cleanup()
        ts.Close()
    })

    It("makes a request, returning the response body", func() {
        strategy := &trackerapi.APITokenStrategy{
            APIToken: token,
        }
        request := trackerapi.NewRequest(resolver.MeRequestURL(), strategy)
        responseBody, _ := request.Execute()
        actual := strings.TrimSpace(string(responseBody))
        Expect(actual).To(Equal(json))
    })
})
