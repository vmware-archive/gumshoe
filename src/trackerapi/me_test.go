package trackerapi_test

import (
    "cmdutil"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "io/ioutil"
    "os"
    "trackerapi"
)

var _ = Describe("Me", func() {
    var json string

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
    })

    It("writes the API token to a file", func() {
        ts := testServer("mister_tee", "sekret", json)
        trackerapi.URL = ts.URL
        trackerapi.FileLocation = "/tmp/temp_tracker"
        trackerapi.Stdout, _ = os.Open(os.DevNull)

        ioutil.WriteFile("/tmp/stdin", []byte("mister_tee\nsekret\n"), 0644)

        file, _ := os.Open("/tmp/stdin")
        cmdutil.InputFile = file
        trackerapi.Me()
        contents, _ := ioutil.ReadFile("/tmp/temp_tracker")
        Expect(string(contents)).To(Equal("abcde90792f3898ab464cd3412345"))
    })
})
