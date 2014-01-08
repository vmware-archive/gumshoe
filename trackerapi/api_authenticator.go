package trackerapi

import (
    "encoding/json"
    "errors"
    "io/ioutil"
    "net/http"
)

const (
    AuthenticateURL = "https://www.pivotaltracker.com/services/v5/me"
)

type APIAuthenticator struct {
    URL  string
    user *User
}

func NewAPIAuthenticator() *APIAuthenticator {
    return &APIAuthenticator{
        URL: AuthenticateURL,
    }
}

func (a *APIAuthenticator) Authenticate(u *User) (string, error) {
    a.user = u
    if !a.user.HasCredentials() {
        return "", errors.New("Given trackerapi.User does not have Username and Password")
    }
    body, err := a.makeRequest()
    if err != nil {
        return "", err
    }

    token, err := a.parse(body)
    if err != nil {
        return "", err
    }
    return token, nil
}

func (a *APIAuthenticator) makeRequest() ([]byte, error) {
    httpClient := &http.Client{}
    req, err := http.NewRequest("GET", a.URL, nil)
    if err != nil {
        return nil, err
    }
    req.SetBasicAuth(a.user.Username, a.user.Password)
    resp, err := httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    return body, nil
}

func (a *APIAuthenticator) parse(body []byte) (string, error) {
    var resp = response{}
    err := json.Unmarshal(body, &resp)
    if err != nil {
        return "", err
    }

    return resp.APIToken, nil
}
