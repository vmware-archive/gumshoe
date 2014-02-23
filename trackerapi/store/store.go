package store

type Store interface {
    Set(key, value string) error
    Get(key string) (string, error)
    Clear() error
}

func handleError(err error) {
    if err != nil {
        panic(err)
    }
}
