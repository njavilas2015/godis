package internal

import (
	"fmt"
	"strconv"
	"sync"
)

type JobListStore struct {
	Command  string
	Key      string
	Value    string
	Index    int
	Start    int
	Stop     int
	Response chan interface{}
}

type QueueListStore struct {
	jobs chan *JobListStore
}

type ListStorage struct {
	data  map[string][]string
	mu    sync.RWMutex
	queue *QueueListStore
}

type CommandHandlerListStorage func(job *JobListStore)

func NewQueueListStore(bufferSize int) *QueueListStore {
	return &QueueListStore{
		jobs: make(chan *JobListStore, bufferSize),
	}
}

func (q *QueueListStore) Add(job *JobListStore) {
	q.jobs <- job
}

func NewListStorage(bufferSize int) *ListStorage {
	storage := &ListStorage{
		data:  make(map[string][]string),
		queue: NewQueueListStore(bufferSize),
	}

	go storage.Process()

	return storage
}

func (ls *ListStorage) Process() {

	commandHandlers := map[string]CommandHandlerListStorage{
		"LPUSH":  ls.LeftPush,
		"RPUSH":  ls.RightPush,
		"LINDEX": ls.ListIndex,
		"LRANGE": ls.ListRange,
		"LPOP":   ls.LeftPop,
	}

	for job := range ls.queue.jobs {

		handler, exists := commandHandlers[job.Command]

		if exists {
			handler(job)
		} else {
			job.Response <- "ERROR: Unknown command"
		}
	}

}

func (ls *ListStorage) LeftPush(job *JobListStore) {
	ls.mu.Lock()

	defer ls.mu.Unlock()

	ls.data[job.Key] = append([]string{job.Value}, ls.data[job.Key]...)

	job.Response <- "OK"
}

func (ls *ListStorage) RightPush(job *JobListStore) {

	ls.mu.Lock()

	defer ls.mu.Unlock()

	ls.data[job.Key] = append(ls.data[job.Key], job.Value)

	job.Response <- "OK"
}

func (ls *ListStorage) ListIndex(job *JobListStore) {

	ls.mu.RLock()

	defer ls.mu.RUnlock()

	list, exists := ls.data[job.Key]

	if !exists || job.Index < 0 || job.Index >= len(list) {
		job.Response <- "ERROR: Index out of bounds or key not found"
		return
	}

	job.Response <- list[job.Index]
}

func (ls *ListStorage) ListRange(job *JobListStore) {

	ls.mu.RLock()

	defer ls.mu.RUnlock()

	list, exists := ls.data[job.Key]

	if !exists {
		job.Response <- "ERROR: Key not found"
		return
	}

	start := job.Start

	stop := job.Stop

	if start < 0 {
		start = 0
	}
	if stop < 0 || stop >= len(list) {
		stop = len(list) - 1
	}

	job.Response <- list[start : stop+1]
}

func (ls *ListStorage) LeftPop(job *JobListStore) {

	ls.mu.Lock()

	defer ls.mu.Unlock()

	if len(ls.data[job.Key]) == 0 {
		job.Response <- "ERROR: List is empty"
		return
	}

	val := ls.data[job.Key][0]

	ls.data[job.Key] = ls.data[job.Key][1:]

	job.Response <- val
}

func (ls *ListStorage) AddJobLeftPush(key string, value string) string {

	job := &JobListStore{
		Command:  "LPUSH",
		Key:      key,
		Value:    value,
		Response: make(chan interface{}),
	}

	ls.queue.Add(job)

	return (<-job.Response).(string)
}

func (ls *ListStorage) AddJobRightPush(key, value string) string {

	job := &JobListStore{
		Command:  "RPUSH",
		Key:      key,
		Value:    value,
		Response: make(chan interface{}),
	}

	ls.queue.Add(job)

	return (<-job.Response).(string)
}

func (ls *ListStorage) AddJobListIndex(key string, index int) string {

	job := &JobListStore{
		Command:  "LINDEX",
		Key:      key,
		Index:    index,
		Response: make(chan interface{}),
	}

	ls.queue.Add(job)

	resp := <-job.Response

	if result, ok := resp.(string); ok {
		return result
	}

	return "ERROR"
}

func (ls *ListStorage) AddJobListRange(key string, start, stop int) []string {

	job := &JobListStore{
		Command:  "LRANGE",
		Key:      key,
		Start:    start,
		Stop:     stop,
		Response: make(chan interface{}),
	}

	ls.queue.Add(job)

	resp := <-job.Response

	if result, ok := resp.([]string); ok {
		return result
	}

	return []string{"ERROR"}
}

func (ls *ListStorage) AddJobLeftPop(key string) string {

	job := &JobListStore{
		Command:  "LPOP",
		Key:      key,
		Response: make(chan interface{}),
	}

	ls.queue.Add(job)

	resp := <-job.Response

	if result, ok := resp.(string); ok {
		return result
	}

	return "ERROR"
}

var Ls *ListStorage = NewListStorage(10)

func HandlerListStore(command string, args []string) <-chan string {
	response := make(chan string)

	go func() {
		defer close(response)

		switch command {
		case "LPUSH":

			valid, errMsg := ValidateArgs(command, args, 2)

			if !valid {
				response <- errMsg
				return
			}

			response <- Ls.AddJobLeftPush(args[0], args[1])

		case "RPUSH":

			valid, errMsg := ValidateArgs(command, args, 2)

			if !valid {
				response <- errMsg
				return
			}

			response <- Ls.AddJobRightPush(args[0], args[1])

		case "LINDEX":

			valid, errMsg := ValidateArgs(command, args, 2)

			if !valid {
				response <- errMsg
				return
			}

			index, err := strconv.Atoi(args[1])

			if err != nil {
				response <- "ERROR: Index must be an integer"
				return
			}

			response <- Ls.AddJobListIndex(args[0], index)

		case "LRANGE":

			valid, errMsg := ValidateArgs(command, args, 3)

			if !valid {
				response <- errMsg
				return
			}

			start, err1 := strconv.Atoi(args[1])

			stop, err2 := strconv.Atoi(args[2])

			if err1 != nil || err2 != nil {
				response <- "ERROR: Start and stop must be integers"
				return
			}

			rangeResult := Ls.AddJobListRange(args[0], start, stop)

			response <- fmt.Sprintf("%v", rangeResult)

		case "LPOP":

			valid, errMsg := ValidateArgs(command, args, 1)

			if !valid {
				response <- errMsg
				return
			}

			response <- Ls.AddJobLeftPop(args[0])

		default:
			response <- "ERROR: Unknown command"
		}
	}()

	return response
}
