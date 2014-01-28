package trackerapi

import (
    "io/ioutil"
    "net/http"
)

type Request struct {
    url          string
    authStrategy RequestAuthStrategy
}

func NewRequest(url string, authStrategy RequestAuthStrategy) Request {
    return Request{
        url:          url,
        authStrategy: authStrategy,
    }
}

func (r *Request) Execute() ([]byte, error) {
    httpClient := &http.Client{}
    req, err := http.NewRequest("GET", r.url, nil)
    if err != nil {
        return []byte{}, err
    }
    r.authStrategy.Strategize(req)
    resp, err := httpClient.Do(req)
    if err != nil {
        return []byte{}, err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return []byte{}, err
    }
    return body, nil
}
