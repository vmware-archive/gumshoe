package responses

import "github.com/pivotal/gumshoe/trackerapi/domain"

type Projects struct {
    Structure []ProjectStructure
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
    projects := make([]domain.Project, len(r.Structure))
    for i, projectStructure := range r.Structure {
        projects[i] = projectStructure.Project()
    }
    return projects
}
