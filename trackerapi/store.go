package trackerapi

import (
    "encoding/json"
    "os"
)

type Store struct {
    cache    map[string]string
    filePath string
}

func NewStore() *Store {
    return &Store{
        cache:    make(map[string]string),
        filePath: homeDir() + "/.tracker",
    }
}

func (s *Store) Set(key, value string) error {
    file, err := os.OpenFile(s.filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
    if err != nil {
        return err
    }
    defer file.Close()
    encoder := json.NewEncoder(file)
    s.cache[key] = value
    encoder.Encode(s.cache)

    return nil
}

func (s *Store) Get(key string) (string, error) {
    file, err := os.OpenFile(s.filePath, os.O_CREATE|os.O_RDONLY, 0666)
    if err != nil {
        return "", err
    }
    defer file.Close()
    decoder := json.NewDecoder(file)
    decoder.Decode(&s.cache)
    return s.cache[key], nil
}

func (s *Store) Clear() error {
    s.cache = make(map[string]string)
    return os.Remove(s.filePath)
}
