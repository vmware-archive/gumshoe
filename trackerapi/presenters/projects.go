package presenters

import (
    "bytes"
    "github.com/pivotal/gumshoe/trackerapi/domain"
)

type Projects struct {
    Projects []domain.Project
}

func (p Projects) String() string {
    var buffer bytes.Buffer

    for _, project := range p.Projects {
        projectPresenter := Project{project}
        buffer.WriteString(projectPresenter.String())
    }

    return buffer.String()
}
