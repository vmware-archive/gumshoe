package store

import (
    "encoding/json"
    "os"
    "os/user"
)

func handleError(err error) {
    if err != nil {
        panic(err)
    }
}

type Store interface {
    Set(key, value string) error
    Get(key string) (string, error)
    Clear() error
}

type FileStore struct {
    cache    map[string]string
    filePath string
}

func NewFileStore() *FileStore {
    u, err := user.Current()
    handleError(err)
    return &FileStore{
        cache:    make(map[string]string),
        filePath: u.HomeDir + "/.tracker",
    }
}

func (s *FileStore) Set(key, value string) error {
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

func (s *FileStore) Get(key string) (string, error) {
    file, err := os.OpenFile(s.filePath, os.O_CREATE|os.O_RDONLY, 0666)
    if err != nil {
        return "", err
    }
    defer file.Close()
    decoder := json.NewDecoder(file)
    decoder.Decode(&s.cache)
    return s.cache[key], nil
}

func (s *FileStore) Clear() error {
    s.cache = make(map[string]string)
    return os.Remove(s.filePath)
}
