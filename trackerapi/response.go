package trackerapi

import "github.com/pivotal/gumshoe/trackerapi/domain"

type MeResponseStructure struct {
    APIToken string `json:"api_token"`
    Username string `json:"username"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Initials string `json:"initials"`
    Timezone struct {
        OlsonName string `json:"olson_name"`
    }   `json:"time_zone"`
}

func (s MeResponseStructure) User() domain.User {
    return domain.User{
        Name:     s.Name,
        Username: s.Username,
        APIToken: s.APIToken,
        Email:    s.Email,
        Initials: s.Initials,
        Timezone: s.Timezone.OlsonName,
    }
}

type ProjectResponseStructure struct {
    ID               int    `json:"id"`
    Name             string `json:"name"`
    Description      string `json:"description"`
    CurrentIteration int    `json:"current_iteration_number"`
}

type ActivityResponseStructure struct {
    Message string `json:"message"`
}
