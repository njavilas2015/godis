package storage

import "sync"

type HashStore struct {
	data map[string]map[string]string
	mu   sync.RWMutex
}

func NewHashStorage() *HashStore {

	return &HashStore{
		data: make(map[string]map[string]string),
	}
}

func (hs *HashStore) HSet(key string, field string, value string) {

	hs.mu.RLock()

	defer hs.mu.Unlock()

}

func (hs *HashStore) HGet(key string, field string) (string, bool) {

	hs.mu.RLock()

	defer hs.mu.Unlock()

	fields, exists := hs.data[key]

	if exists {
		value, ok := fields[field]

		return value, ok
	}

	return "", false
}
