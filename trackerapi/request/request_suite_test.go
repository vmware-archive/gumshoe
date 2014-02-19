package request_test

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

func TestRequest(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Request Suite")
}

var errorJson string = `{
    "possible_fix": "Recheck your name (email address or Tracker username) and password.",
    "error":        "Invalid authentication credentials were presented.",
    "kind":         "error",
    "code":         "invalid_authentication"
}`

type TestServer struct {
    server    *httptest.Server
    URL       string
    username  string
    password  string
    apiToken  string
    responses map[string]string
}

func (s *TestServer) Boot() {
    s.server = httptest.NewServer(s.buildHandler())
    s.URL = s.server.URL
    s.responses = make(map[string]string)
}

func (s *TestServer) Close() {
    s.server.Close()
}

func (s *TestServer) SetResponse(path, json string) {
    s.responses[path] = json
}

func (s *TestServer) GetResponse(path string) string {
    response := s.responses[path]
    if response == "" {
        panic(path)
    }
    return response
}

func (s *TestServer) buildHandler() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        if s.username != "" && s.password != "" {
            s.handleBasicAuthRequest(w, r)
        } else {
            s.handleAPITokenRequest(w, r)
        }
    })
}

func (s *TestServer) handleBasicAuthRequest(w http.ResponseWriter, r *http.Request) {
    auth, err := s.parseBasicAuth(r)
    if err != nil {
        fmt.Fprintln(w, errorJson)
    } else {
        if auth[0] == s.username && auth[1] == s.password {
            fmt.Fprintln(w, s.GetResponse(r.URL.Path))
        } else {
            fmt.Fprintln(w, errorJson)
        }
    }
}

func (s *TestServer) handleAPITokenRequest(w http.ResponseWriter, r *http.Request) {
    headerToken := r.Header.Get("X-TrackerToken")
    if s.apiToken == headerToken {
        fmt.Fprintln(w, s.GetResponse(r.URL.Path))
    } else {
        fmt.Fprintln(w, errorJson)
    }
}

func (s *TestServer) parseBasicAuth(r *http.Request) ([]string, error) {
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
