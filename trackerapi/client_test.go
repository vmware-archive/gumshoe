package trackerapi_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi"
    "github.com/pivotal/gumshoe/trackerapi/request"
    "github.com/pivotal/gumshoe/trackerapi/store"
)

var _ = Describe("Client", func() {
    var (
        json   string
        client *trackerapi.Client
        ts     *TestServer
    )

    BeforeEach(func() {
        config := trackerapi.NewConfiguration()
        config.Store = store.NewMemoryStore()
        config.Store.Set("APIToken", "abcde90792f3898ab464cd3412345")
        client, _ = trackerapi.NewClient(config)
    })

    AfterEach(func() {
        client.Cleanup()
        ts.Close()
    })

    Describe("Me", func() {
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

            ts = &TestServer{
                apiToken: "abcde90792f3898ab464cd3412345",
            }
            ts.Boot()
            ts.SetResponse("/me", json)
            client.SetResolver(request.Resolver{
                TrackerDomain: ts.URL,
            })
        })

        It("prints the user representation to the output file", func() {
            output := client.Me()

            printedOutput := output.String()
            Expect(printedOutput).To(ContainSubstring("Mister Tee (mister_tee)"))
            Expect(printedOutput).To(ContainSubstring("  Email     : mister_tee@pivotallabs.com"))
            Expect(printedOutput).To(ContainSubstring("  API Token : abcde90792f3898ab464cd3412345"))
            Expect(printedOutput).To(ContainSubstring("  Initials  : MT"))
            Expect(printedOutput).To(ContainSubstring("  Timezone  : America/Los_Angeles"))
        })
    })

    Describe("Projects", func() {
        BeforeEach(func() {
            json = `[
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

            ts = &TestServer{
                apiToken: "abcde90792f3898ab464cd3412345",
            }
            ts.Boot()
            ts.SetResponse("/projects", json)
            client.SetResolver(request.Resolver{
                TrackerDomain: ts.URL,
            })
        })

        It("prints a representation of the user's projects to the screen", func() {
            output := client.Projects()

            printedOutput := output.String()
            Expect(printedOutput).To(ContainSubstring("Learn About the Force (98)"))
            Expect(printedOutput).To(ContainSubstring("  Current Iteration : 1"))

            Expect(printedOutput).To(ContainSubstring("Death Star (99)"))
            Expect(printedOutput).To(ContainSubstring("  Description       : Expeditionary Battle Planetoid"))
            Expect(printedOutput).To(ContainSubstring("  Current Iteration : 15"))
        })
    })

    Describe("Activity", func() {
        BeforeEach(func() {
            json = `[
                {"message": "Wilhuff Tarkin changed iteration 1's length from default to 2 weeks"},
                {"message": "Hey Girl, I'm just like coffee. I'll keep you up all night."}
            ]`
            ts = &TestServer{
                apiToken: "abcde90792f3898ab464cd3412345",
            }
            ts.Boot()
            ts.SetResponse("/projects/124/activity", json)
            client.SetResolver(request.Resolver{
                TrackerDomain: ts.URL,
            })
        })

        It("prints a representation of the user's projects to the screen", func() {
            output := client.Activity(124)

            printedOutput := output.String()
            Expect(printedOutput).To(ContainSubstring("Activity:\n"))
            Expect(printedOutput).To(ContainSubstring("  Wilhuff Tarkin changed iteration 1's length from default to 2 weeks"))
            Expect(printedOutput).To(ContainSubstring("  Hey Girl, I'm just like coffee. I'll keep you up all night."))
        })
    })
})
