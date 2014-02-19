package responses

import (
    "encoding/json"
    "github.com/pivotal/gumshoe/trackerapi/domain"
)

type Me struct {
    structure MeStructure
}

func (m *Me) Parse(body []byte) error {
    return json.Unmarshal(body, &m.structure)
}

type MeStructure struct {
    APIToken string `json:"api_token"`
    Username string
    Name     string
    Email    string
    Initials string
    Timezone struct {
        OlsonName string `json:"olson_name"`
    }   `json:"time_zone"`
}

func (r Me) User() domain.User {
    s := r.structure
    return domain.User{
        Name:     s.Name,
        Username: s.Username,
        APIToken: s.APIToken,
        Email:    s.Email,
        Initials: s.Initials,
        Timezone: s.Timezone.OlsonName,
    }
}
