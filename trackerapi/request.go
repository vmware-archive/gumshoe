package trackerapi

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

type Request struct {
    url       string
    APIToken  string
    structure interface{}
}

func (r *Request) Execute() error {
    httpClient := &http.Client{}
    req, err := http.NewRequest("GET", r.url, nil)
    if err != nil {
        return err
    }
    req.Header.Set("X-TrackerToken", r.APIToken)
    resp, err := httpClient.Do(req)
    if err != nil {
        return err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    err = json.Unmarshal(body, r.structure)
    if err != nil {
        return err
    }
    return nil
}
