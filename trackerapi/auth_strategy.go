package trackerapi

import "net/http"

type RequestAuthStrategy interface {
    Strategize(*http.Request)
}

type APITokenStrategy struct {
    APIToken string
}

type BasicAuthStrategy struct {
    Username string
    Password string
}

func (s *BasicAuthStrategy) Strategize(r *http.Request) {
    r.SetBasicAuth(s.Username, s.Password)
}

func (s *APITokenStrategy) Strategize(r *http.Request) {
    r.Header.Set("X-TrackerToken", s.APIToken)
}
