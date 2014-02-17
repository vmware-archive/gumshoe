package responses

import "github.com/pivotal/gumshoe/trackerapi/domain"

type Activities struct {
    Structure []ActivityStructure
}

type ActivityStructure struct {
    Message string `json:"message"`
}

func (s ActivityStructure) Activity() domain.Activity {
    return domain.Activity{
        Message: s.Message,
    }
}

func (a Activities) Activities() []domain.Activity {
    activities := make([]domain.Activity, len(a.Structure))
    for i, activityStructure := range a.Structure {
        activities[i] = activityStructure.Activity()
    }
    return activities
}
