package trackerapi_test

import (
    "fmt"
    "io/ioutil"
    "os"

    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/cmdutil"
    "github.com/pivotal/gumshoe/trackerapi"
)

var _ = Describe("APIAuthenticator #Authenticate", func() {

    var (
        user *trackerapi.User
        auth *trackerapi.APIAuthenticator
        json string
    )

    BeforeEach(func() {
        user = &trackerapi.User{
            Username: "mister_tee",
            Password: "sekret",
        }

        auth = &trackerapi.APIAuthenticator{}
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

        ts := testServer("mister_tee", "sekret", "", json)
        auth.URL = ts.URL
        cmdutil.InputFile, _ = os.Create("/tmp/stdin")
        cmdutil.InputBuffer = nil
        ioutil.WriteFile("/tmp/stdin", []byte("mister_tee\nsekret\n"), 0644)
    })

    AfterEach(func() {
        os.Remove("/tmp/stdin")
    })

    It("makes a request to the tracker api, prompting the user for their creds, and returning the api token", func() {
        token, _ := auth.Authenticate(user)
        Expect(token).To(Equal("abcde90792f3898ab464cd3412345"))
    })

    It("requires a user with a username and password", func() {
        _, err := auth.Authenticate(&trackerapi.User{})
        Expect(fmt.Sprint(err)).To(Equal("Given trackerapi.User does not have Username and Password"))
    })
})
