package store

type MemoryStore struct {
    cache map[string]string
}

func NewMemoryStore() *MemoryStore {
    return &MemoryStore{
        cache: make(map[string]string),
    }
}

func (s *MemoryStore) Set(key, value string) error {
    s.cache[key] = value
    return nil
}

func (s *MemoryStore) Get(key string) (string, error) {
    return s.cache[key], nil
}

func (s *MemoryStore) Clear() error {
    s.cache = make(map[string]string)
    return nil
}
