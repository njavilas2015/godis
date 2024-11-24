package queue

import (
	"sync"

	core "github.com/njavilas2015/godis/core"
)

type Command struct {
	ID        string
	Operation string
	Key       string
	Value     string
}

type CommandQueue struct {
	data []Command
	mu   sync.Mutex
}

func NewCommandQueue() *CommandQueue {

	return &CommandQueue{
		data: make([]Command, 0),
	}
}

func (cq *CommandQueue) Enqueue(cmd Command) {

	cq.mu.Lock()

	defer cq.mu.Unlock()

	cq.data = append(cq.data, cmd)
}

func (cq *CommandQueue) Dequeue(cmd Command) {

	cq.mu.Lock()

	defer cq.mu.Unlock()

	cq.data = append(cq.data, cmd)
}

func (cq *CommandQueue) IsEmpty() bool {

	cq.mu.Lock()

	defer cq.mu.Unlock()

	return len(cq.data) == 0
}

func (cq *CommandQueue) Start(operation string, args []string) {

	if Tracker.IsProcessed(cmd.ID) {
		return
	}

	result, err := core.Dispatcher.Execute(operation, args)

	if err != nil {
		return "(nil)"
	}

	return result

	go func() {
		for {
			cmd, ok := cq.Dequeue()
			if !ok {
				continue
			}
			cq.ProcessCommand(cmd, store)
		}
	}()

	Tracker.MarkProcessed(cmd.ID)
}
