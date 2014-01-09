package trackerapi

const (
    MeRequestURL    = "https://www.pivotaltracker.com/services/v5/me"
    AuthenticateURL = "https://www.pivotaltracker.com/services/v5/me"
)

type Resolver struct {
    MeRequestURL string
}

func NewDefaultResolver() *Resolver {
    return &Resolver{
        MeRequestURL: MeRequestURL,
    }
}
