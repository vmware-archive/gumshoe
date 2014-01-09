package trackerapi_test

import (
    "io/ioutil"
    "os"

    . "github.com/pivotal/gumshoe/ginkgo"
    . "github.com/pivotal/gumshoe/gomega"
    "github.com/pivotal/gumshoe/trackerapi"
)

var _ = Describe("Client #Me", func() {
    var (
        json   string
        client *trackerapi.Client
    )

    BeforeEach(func() {
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

        client = trackerapi.NewClient()
        ts := testServer("", "", "abcde90792f3898ab464cd3412345", json)
        client.SetResolver(&trackerapi.Resolver{
            MeRequestURL: ts.URL,
        })
        store := trackerapi.NewStore()
        store.Set("APIToken", "abcde90792f3898ab464cd3412345")
        client.SetLogger(trackerapi.NewFileLogger("/tmp/stdout"))
    })

    AfterEach(func() {
        client.Cleanup()
        os.Remove("/tmp/stdout")
    })

    It("prints the user representation to the output file", func() {
        client.Me()

        fileContents, _ := ioutil.ReadFile("/tmp/stdout")
        Expect(string(fileContents)).To(ContainSubstring("Username:  mister_tee"))
        Expect(string(fileContents)).To(ContainSubstring("Name:      Mister Tee"))
        Expect(string(fileContents)).To(ContainSubstring("Email:     mister_tee@pivotallabs.com"))
        Expect(string(fileContents)).To(ContainSubstring("API Token: abcde90792f3898ab464cd3412345"))
        Expect(string(fileContents)).To(ContainSubstring("Initials:  MT"))
        Expect(string(fileContents)).To(ContainSubstring("Timezone:  America/Los_Angeles"))
    })
})
