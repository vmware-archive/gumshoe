package trackerapi_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"

    "encoding/base64"
    "errors"
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
            auth, err := parseBasicAuth(r)
            if err != nil {
                fmt.Fprintln(w, errorJson)
            } else {
                if auth[0] == username && auth[1] == password {
                    fmt.Fprintln(w, json)
                } else {
                    fmt.Fprintln(w, errorJson)
                }
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

func parseBasicAuth(r *http.Request) ([]string, error) {
    auth := r.Header["Authorization"]
    if auth != nil {
        var header string = auth[0]
        parts := strings.SplitN(header, " ", 2)
        b, _ := base64.StdEncoding.DecodeString(parts[1])
        return strings.Split(string(b), ":"), nil
    } else {
        return []string{}, errors.New("")
    }
}
