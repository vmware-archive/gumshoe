package request

import (
    "io/ioutil"
    "net/http"
)

type Requester struct {
    url          string
    authStrategy RequestAuthStrategy
}

func New(url string, authStrategy RequestAuthStrategy) Requester {
    return Requester{
        url:          url,
        authStrategy: authStrategy,
    }
}

func (r *Requester) Execute() ([]byte, error) {
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
