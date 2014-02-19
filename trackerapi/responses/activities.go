package responses

import (
    "encoding/json"
    "github.com/pivotal/gumshoe/trackerapi/domain"
)

type Activities struct {
    structure []ActivityStructure
}

func (a *Activities) Parse(body []byte) error {
    return json.Unmarshal(body, &a.structure)
}

func (a Activities) Activities() []domain.Activity {
    activities := make([]domain.Activity, len(a.structure))
    for i, activityStructure := range a.structure {
        activities[i] = activityStructure.Activity()
    }
    return activities
}

type ActivityStructure struct {
    Message string
}

func (s ActivityStructure) Activity() domain.Activity {
    return domain.Activity{
        Message: s.Message,
    }
}
