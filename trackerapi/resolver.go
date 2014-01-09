package trackerapi

const (
    MeRequestURL       = "https://www.pivotaltracker.com/services/v5/me"
    AuthenticateURL    = "https://www.pivotaltracker.com/services/v5/me"
    ProjectsRequestURL = "https://www.pivotaltracker.com/services/v5/projects"
)

type Resolver struct {
    MeRequestURL       string
    ProjectsRequestURL string
}

func NewDefaultResolver() *Resolver {
    return &Resolver{
        MeRequestURL:       MeRequestURL,
        ProjectsRequestURL: ProjectsRequestURL,
    }
}
