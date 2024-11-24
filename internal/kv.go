package internal

import (
	"sync"
)

type JobKvStorage struct {
	Command  string
	Key      string
	Value    string
	Response chan string
}

type QueueKvStorage struct {
	jobs chan *JobKvStorage
}

type KvStorage struct {
	data  map[string]string
	mu    sync.RWMutex
	queue *QueueKvStorage
}

type CommandHandlerKvStorage func(job *JobKvStorage)

func NewKvQueue(bufferSize int) *QueueKvStorage {
	return &QueueKvStorage{
		jobs: make(chan *JobKvStorage, bufferSize),
	}
}

func (q *QueueKvStorage) Add(job *JobKvStorage) {
	q.jobs <- job
}

func (s *KvStorage) Set(job *JobKvStorage) {

	s.mu.Lock()

	defer s.mu.Unlock()

	s.data[job.Key] = job.Value

	job.Response <- "OK"
}

func (s *KvStorage) Get(job *JobKvStorage) {

	s.mu.RLock()

	defer s.mu.RUnlock()

	value, exists := s.data[job.Key]

	if exists {
		job.Response <- value
	} else {
		job.Response <- "NOT FOUND"
	}
}

func (s *KvStorage) Delete(job *JobKvStorage) {

	s.mu.Lock()

	defer s.mu.Unlock()

	delete(s.data, job.Key)

	job.Response <- "OK"
}

func NewKvStorage(bufferSize int) *KvStorage {
	storage := &KvStorage{
		data:  make(map[string]string),
		queue: NewKvQueue(bufferSize),
	}

	go storage.Process()

	return storage
}

func (s *KvStorage) Process() {

	commandHandlers := map[string]CommandHandlerKvStorage{
		"SET":    s.Set,
		"GET":    s.Get,
		"DELETE": s.Delete,
	}

	for job := range s.queue.jobs {

		handler, exists := commandHandlers[job.Command]

		if exists {
			handler(job)
		} else {
			job.Response <- "ERROR: Unknown command"
		}
	}
}

func (s *KvStorage) AddJobSet(key string, value string) string {
	job := &JobKvStorage{
		Command:  "SET",
		Key:      key,
		Value:    value,
		Response: make(chan string),
	}
	s.queue.Add(job)
	return <-job.Response
}

func (s *KvStorage) AddJobGet(key string) string {
	job := &JobKvStorage{
		Command:  "GET",
		Key:      key,
		Response: make(chan string),
	}
	s.queue.Add(job)
	return <-job.Response
}

func (s *KvStorage) AddJobDelete(key string) string {
	job := &JobKvStorage{
		Command:  "DELETE",
		Key:      key,
		Response: make(chan string),
	}
	s.queue.Add(job)
	return <-job.Response
}

var Kv *KvStorage = NewKvStorage(10)

func HandlerKvStore(command string, args []string) <-chan string {

	response := make(chan string)

	go func() {
		defer close(response)

		switch command {
		case "SET":

			valid, errMsg := ValidateArgs(command, args, 2)

			if !valid {
				response <- errMsg
				return
			}

			response <- Kv.AddJobSet(args[0], args[1])

		case "GET":

			valid, errMsg := ValidateArgs(command, args, 1)

			if !valid {
				response <- errMsg
				return
			}

			response <- Kv.AddJobGet(args[0])

		case "DELETE":

			valid, errMsg := ValidateArgs(command, args, 1)

			if !valid {
				response <- errMsg
				return
			}

			response <- Kv.AddJobDelete(args[0])

		default:
			response <- "ERROR: Unknown command"
		}
	}()

	return response
}
