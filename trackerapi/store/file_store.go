package store

import (
    "encoding/json"
    "os"
    "os/user"
)

type FileStore struct {
    memStore *MemoryStore
    filePath string
}

func NewFileStore() *FileStore {
    u, err := user.Current()
    handleError(err)
    return &FileStore{
        memStore: NewMemoryStore(),
        filePath: u.HomeDir + "/.tracker",
    }
}

func (s *FileStore) Set(key, value string) error {
    file, err := os.OpenFile(s.filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
    handleError(err)
    defer file.Close()
    encoder := json.NewEncoder(file)
    s.memStore.Set(key, value)
    encoder.Encode(s.memStore.cache)

    return nil
}

func (s *FileStore) Get(key string) (string, error) {
    file, err := os.OpenFile(s.filePath, os.O_CREATE|os.O_RDONLY, 0666)
    handleError(err)
    defer file.Close()
    decoder := json.NewDecoder(file)
    decoder.Decode(&s.memStore.cache)
    return s.memStore.Get(key)
}

func (s *FileStore) Clear() error {
    s.memStore.Clear()
    return os.Remove(s.filePath)
}
