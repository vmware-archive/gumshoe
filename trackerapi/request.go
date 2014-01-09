package trackerapi

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

type Request struct {
    url            string
    responseStruct interface{}
    authStrategy   RequestAuthStrategy
}

func (r *Request) Execute() error {
    httpClient := &http.Client{}
    req, err := http.NewRequest("GET", r.url, nil)
    if err != nil {
        return err
    }
    r.authStrategy.Strategize(req)
    resp, err := httpClient.Do(req)
    if err != nil {
        return err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    err = json.Unmarshal(body, r.responseStruct)
    if err != nil {
        return err
    }
    return nil
}
