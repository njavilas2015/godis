package storage

import (
	"fmt"
	"sync"
)

type ListStorage struct {
	data map[string][]string
	mu   sync.RWMutex
}

func NewListStorage() *ListStorage {

	return &ListStorage{
		data: make(map[string][]string),
	}
}

func (ls *ListStorage) LeftPush(key string, value string) {

	ls.mu.Lock()

	defer ls.mu.Unlock()

	ls.data[key] = append([]string{value}, ls.data[key]...)
}

func (ls *ListStorage) RightPush(key string, value string) {

	ls.mu.Lock()

	defer ls.mu.Unlock()

	ls.data[key] = append(ls.data[key], value)
}

func (ls *ListStorage) ListIndex(key string, index int) (string, error) {

	ls.mu.RLock()

	defer ls.mu.RUnlock()

	list, exists := ls.data[key]

	if !exists {
		return "", fmt.Errorf("key not found")
	}

	if index < 0 || index >= len(list) {
		return "", fmt.Errorf("index out of bounds")
	}

	return list[index], nil
}

func (ls *ListStorage) ListRange(key string, start, stop int) ([]string, error) {

	ls.mu.RLock()

	defer ls.mu.RUnlock()

	list, exists := ls.data[key]

	if !exists {
		return nil, fmt.Errorf("key not found")
	}

	if start < 0 {
		start = 0
	}
	if stop < 0 || stop >= len(list) {
		stop = len(list) - 1
	}

	return list[start : stop+1], nil
}

func (ls *ListStorage) LeftPop(key string) (string, bool) {

	ls.mu.Lock()

	defer ls.mu.Unlock()

	if len(ls.data[key]) == 0 {
		return "", false
	}

	val := ls.data[key][0]

	ls.data[key] = ls.data[key][1:]

	return val, true
}

func (ls *ListStorage) Get(key string) ([]string, bool) {

	ls.mu.Lock()

	defer ls.mu.Unlock()

	values, exists := ls.data[key]

	return values, exists
}
