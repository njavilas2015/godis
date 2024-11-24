package queue

import "sync"

var Tracker = NewProcessedTracker()

type ProcessedTracker struct {
	mu        sync.Mutex
	processed map[string]bool
}

func NewProcessedTracker() *ProcessedTracker {
	return &ProcessedTracker{
		processed: make(map[string]bool),
	}
}

func (pt *ProcessedTracker) IsProcessed(id string) bool {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	return pt.processed[id]
}

func (pt *ProcessedTracker) MarkProcessed(id string) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.processed[id] = true
}
