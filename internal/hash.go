package internal

import (
	"sync"
)

type JobHashStore struct {
	Command  string
	Key      string
	Field    string
	Value    string
	Response chan string
}

type QueueHashStore struct {
	jobs chan *JobHashStore
}

type HashStore struct {
	data  map[string]map[string]string
	mu    sync.RWMutex
	queue *QueueHashStore
}

type CommandHandlerHashStore func(job *JobHashStore)

func NewQueueHashStore(bufferSize int) *QueueHashStore {
	return &QueueHashStore{
		jobs: make(chan *JobHashStore, bufferSize),
	}
}

func (q *QueueHashStore) Add(job *JobHashStore) {
	q.jobs <- job
}

func (hs *HashStore) HSet(job *JobHashStore) {

	hs.mu.Lock()

	_, exists := hs.data[job.Key]

	if !exists {
		hs.data[job.Key] = make(map[string]string)
	}

	hs.data[job.Key][job.Field] = job.Value

	hs.mu.Unlock()

	job.Response <- "OK"
}

func (hs *HashStore) HGet(job *JobHashStore) {

	hs.mu.RLock()

	fields, exists := hs.data[job.Key]

	if exists {

		value, ok := fields[job.Field]

		if ok {
			job.Response <- value
		} else {
			job.Response <- "NOT FOUND"
		}

	} else {
		job.Response <- "NOT FOUND"
	}

	hs.mu.RUnlock()
}

func NewHashStorage(bufferSize int) *HashStore {

	storage := &HashStore{
		data:  make(map[string]map[string]string),
		queue: NewQueueHashStore(bufferSize),
	}

	go storage.Process()

	return storage
}

func (hs *HashStore) AddJobHSet(key, field, value string) string {
	job := &JobHashStore{
		Command:  "HSET",
		Key:      key,
		Field:    field,
		Value:    value,
		Response: make(chan string),
	}
	hs.queue.Add(job)
	return <-job.Response
}

func (hs *HashStore) AddJobHGet(key, field string) string {
	job := &JobHashStore{
		Command:  "HGET",
		Key:      key,
		Field:    field,
		Response: make(chan string),
	}
	hs.queue.Add(job)
	return <-job.Response
}

var Hs *HashStore = NewHashStorage(10)

func (hs *HashStore) Process() {

	commandHandlers := map[string]CommandHandlerHashStore{
		"HGET": hs.HGet,
		"HSET": hs.HSet,
	}

	for job := range hs.queue.jobs {

		handler, exists := commandHandlers[job.Command]

		if exists {
			handler(job)
		} else {
			job.Response <- "ERROR: Unknown command"
		}
	}
}

func HandlerHashStore(command string, args []string) <-chan string {

	response := make(chan string)

	go func() {
		defer close(response)

		switch command {
		case "HSET":

			valid, errMsg := ValidateArgs(command, args, 3)

			if !valid {
				response <- errMsg
				return
			}

			response <- Hs.AddJobHSet(args[0], args[1], args[2])

		case "HGET":

			valid, errMsg := ValidateArgs(command, args, 2)

			if !valid {
				response <- errMsg
				return
			}

			response <- Hs.AddJobHGet(args[0], args[1])

		default:
			response <- "ERROR: Unknown command"
		}
	}()

	return response
}