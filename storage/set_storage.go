package storage

import "sync"

type SetStorage struct {
	data map[string]map[string]struct{}
	mu   sync.RWMutex
}

func NewSetStorage() *SetStorage {

	return &SetStorage{
		data: make(map[string]map[string]struct{}),
	}
}

func (ss *SetStorage) SetAdd(key string, value string) {

	ss.mu.Lock()

	defer ss.mu.Unlock()

	_, exists := ss.data[key]

	if !exists {
		ss.data[key] = make(map[string]struct{})
	}

	ss.data[key][value] = struct{}{}
}

func (ss *SetStorage) SetMembers(key string) ([]string, bool) {

	ss.mu.Lock()

	defer ss.mu.Unlock()

	set, exists := ss.data[key]

	if !exists {
		return nil, false
	}

	members := make([]string, 0, len(set))

	for member := range set {
		members = append(members, member)
	}

	return members, true
}
