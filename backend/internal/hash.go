package internal

import (
	"fmt"
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
			job.Response <- ""
		}

	} else {
		job.Response <- "NOT FOUND"
	}

	hs.mu.RUnlock()
}

func (hs *HashStore) Drop(job *JobHashStore) {

	hs.mu.Lock()

	defer hs.mu.Unlock()

	hs.data = make(map[string]map[string]string)

	job.Response <- "OK"
}

func (hs *HashStore) Has(job *JobHashStore) {

	hs.mu.RLock()

	defer hs.mu.RUnlock()

	if _, exists := hs.data[job.Key]; exists {
		job.Response <- "true"
	} else {
		job.Response <- "false"
	}
}

func (hs *HashStore) All(job *JobHashStore) {

	hs.mu.RLock()

	defer hs.mu.RUnlock()

	keys := make([]string, 0, len(hs.data))

	for key := range hs.data {
		keys = append(keys, key)
	}

	job.Response <- fmt.Sprintf("%v", keys)
}

func NewHashStorage(bufferSize int) *HashStore {

	storage := &HashStore{
		data:  make(map[string]map[string]string),
		queue: NewQueueHashStore(bufferSize),
	}

	go storage.Process()

	return storage
}

func (hs *HashStore) AddJobHSet(key string, field string, value string) string {

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

func (hs *HashStore) AddJobHGet(key string, field string) string {

	job := &JobHashStore{
		Command:  "HGET",
		Key:      key,
		Field:    field,
		Response: make(chan string),
	}

	hs.queue.Add(job)

	return <-job.Response
}

func (hs *HashStore) AddJobDrop() string {

	job := &JobHashStore{
		Command:  "DROP",
		Response: make(chan string),
	}

	hs.queue.Add(job)

	return <-job.Response
}

func (hs *HashStore) AddJobAll() string {

	job := &JobHashStore{
		Command:  "ALL",
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
		"DROP": hs.Drop,
		"ALL":  hs.All,
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

			key := args[0]

			keyValue := args[1:]

			index := len(keyValue)

			if !IsEven(index) {

				response <- "ERROR: HSET incorrectly configured"

				return
			}

			for i := range keyValue {

				if i == 0 {
					continue
				}

				if IsEven(i) {
					continue
				}

				response <- Hs.AddJobHSet(key, args[i], args[i+1])
			}

		case "HGET":

			valid, errMsg := ValidateArgs(command, args, 2)

			if !valid {
				response <- errMsg
				return
			}

			response <- Hs.AddJobHGet(args[0], args[1])

		case "DROP":

			if len(args) > 0 {
				response <- "ERROR: DROPALL does not take arguments"
				return
			}

			response <- Hs.AddJobDrop()

		case "ALL":

			if len(args) > 0 {
				response <- "ERROR: READALL does not take arguments"
				return
			}

			response <- Hs.AddJobAll()

		default:
			response <- "ERROR: Unknown command"
		}
	}()

	return response
}
