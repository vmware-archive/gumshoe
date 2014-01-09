package trackerapi_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    "encoding/base64"
    "fmt"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestTrackerApi(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Tracker API Suite")
}

var errorJson string = `{
    "possible_fix": "Recheck your name (email address or Tracker username) and password.",
    "error":        "Invalid authentication credentials were presented.",
    "kind":         "error",
    "code":         "invalid_authentication"
}`

func testServer(username, password, token, json string) (ts *httptest.Server) {
    ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        headerToken := r.Header.Get("X-TrackerToken")
        if username != "" && password != "" {
            auth := parseBasicAuth(r)
            if auth[0] == username && auth[1] == password {
                fmt.Fprintln(w, json)
            } else {
                fmt.Fprintln(w, errorJson)
            }
        } else {
            if token == headerToken {
                fmt.Fprintln(w, json)
            } else {
                fmt.Fprintln(w, errorJson)
            }
        }
    }))
    return
}

func parseBasicAuth(r *http.Request) []string {
    var header string = r.Header["Authorization"][0]
    parts := strings.SplitN(header, " ", 2)
    b, _ := base64.StdEncoding.DecodeString(parts[1])
    return strings.Split(string(b), ":")
}
