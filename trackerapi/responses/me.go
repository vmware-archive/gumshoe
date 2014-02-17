package responses

import "github.com/pivotal/gumshoe/trackerapi/domain"

type Me struct {
    Structure MeStructure
}

type MeStructure struct {
    APIToken string `json:"api_token"`
    Username string `json:"username"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Initials string `json:"initials"`
    Timezone struct {
        OlsonName string `json:"olson_name"`
    }   `json:"time_zone"`
}

func (r Me) User() domain.User {
    return domain.User{
        Name:     r.Structure.Name,
        Username: r.Structure.Username,
        APIToken: r.Structure.APIToken,
        Email:    r.Structure.Email,
        Initials: r.Structure.Initials,
        Timezone: r.Structure.Timezone.OlsonName,
    }
}
