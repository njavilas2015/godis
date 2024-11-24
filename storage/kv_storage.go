package storage

import "sync"

type Storage struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewStorage() *Storage {

	return &Storage{
		data: make(map[string]string),
	}
}

func (s *Storage) Set(key, value string) {

	s.mu.Lock()

	defer s.mu.Unlock()

	s.data[key] = value
}

func (s *Storage) Get(key string) (string, bool) {

	s.mu.Lock()

	defer s.mu.Unlock()

	value, exists := s.data[key]

	return value, exists
}

func (s *Storage) Delete(key string) {

	s.mu.Lock()

	defer s.mu.Unlock()

	delete(s.data, key)
}
