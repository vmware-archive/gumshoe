package responses

import (
    "encoding/json"
    "github.com/pivotal/gumshoe/trackerapi/domain"
)

type Projects struct {
    structure []ProjectStructure
}

func (p *Projects) Parse(body []byte) error {
    return json.Unmarshal(body, &p.structure)
}

type ProjectStructure struct {
    ID               int    `json:"id"`
    Name             string `json:"name"`
    Description      string `json:"description"`
    CurrentIteration int    `json:"current_iteration_number"`
}

func (s ProjectStructure) Project() domain.Project {
    return domain.Project{
        ID:               s.ID,
        Name:             s.Name,
        Description:      s.Description,
        CurrentIteration: s.CurrentIteration,
    }
}

func (r Projects) Projects() []domain.Project {
    projects := make([]domain.Project, len(r.structure))
    for i, projectStructure := range r.structure {
        projects[i] = projectStructure.Project()
    }
    return projects
}
