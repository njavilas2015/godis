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

func NewQueue(bufferSize int) *QueueHashStore {
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

	hs := &HashStore{
		data:  make(map[string]map[string]string),
		queue: NewQueue(bufferSize),
	}

	go hs.Process()

	return hs
}

func (hs *HashStore) Process() {
	for job := range hs.queue.jobs {
		switch job.Command {
		case "HSET":
			hs.HSet(job)
		case "HGET":
			hs.HGet(job)
		default:
			job.Response <- "ERROR: Unknown command"
		}
	}
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

func Handler(command string, args []string) <-chan string {

	response := make(chan string)

	go func() {
		defer close(response)

		switch command {
		case "HSET":

			if len(args) < 3 {
				response <- "ERROR: HSET requires key, field, and value"
				return
			}

			response <- Hs.AddJobHSet(args[0], args[1], args[2])

		case "HGET":

			if len(args) < 2 {
				response <- "ERROR: HGET requires key and field"
				return
			}

			response <- Hs.AddJobHGet(args[0], args[1])

		default:
			response <- "ERROR: Unknown command"
		}
	}()

	return response
}
