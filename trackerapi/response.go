package trackerapi

import "github.com/pivotal/gumshoe/trackerapi/domain"

type MeResponse struct {
    Structure MeResponseStructure
}

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

func (r MeResponse) User() domain.User {
    return domain.User{
        Name:     r.Structure.Name,
        Username: r.Structure.Username,
        APIToken: r.Structure.APIToken,
        Email:    r.Structure.Email,
        Initials: r.Structure.Initials,
        Timezone: r.Structure.Timezone.OlsonName,
    }
}

type ProjectsResponse struct {
    Structure []ProjectResponseStructure
}

type ProjectResponseStructure struct {
    ID               int    `json:"id"`
    Name             string `json:"name"`
    Description      string `json:"description"`
    CurrentIteration int    `json:"current_iteration_number"`
}

func (s ProjectResponseStructure) Project() domain.Project {
    return domain.Project{
        ID:               s.ID,
        Name:             s.Name,
        Description:      s.Description,
        CurrentIteration: s.CurrentIteration,
    }
}

func (r ProjectsResponse) Projects() []domain.Project {
    projects := make([]domain.Project, len(r.Structure))
    for i, projectStructure := range r.Structure {
        projects[i] = projectStructure.Project()
    }
    return projects
}

type ActivityResponseStructure struct {
    Message string `json:"message"`
}
