package trackerapi_test

import (
    "fmt"
    "io/ioutil"
    "os"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "github.com/pivotal/gumshoe/cmdutil"
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
        ts := testServer("mister_tee", "sekret", json)
        client.URL = ts.URL
        client.FileLocation = "/tmp/temp_tracker"
        cmdutil.InputFile, _ = os.Create("/tmp/stdin")
        cmdutil.InputBuffer = nil
        client.SetLogger(trackerapi.NewFileLogger("/tmp/stdout"))
        ioutil.WriteFile("/tmp/stdin", []byte("mister_tee\nsekret\n"), 0644)
    })

    AfterEach(func() {
        os.Remove("/tmp/stdin")
        os.Remove("/tmp/stdout")
        os.Remove("/tmp/temp_tracker")
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

func ExampleClient_SetLogger() {
    client := trackerapi.NewClient()
    fmt.Sprintf("%v", client.Logger)

    // Output
    // foo
}
