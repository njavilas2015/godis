package internal

import "sync"

type JobSetStorage struct {
	Command  string
	Key      string
	Value    string
	Response chan interface{}
}

type QueueSetStorage struct {
	jobs chan *JobSetStorage
}

type SetStorage struct {
	data  map[string]map[string]struct{}
	mu    sync.RWMutex
	queue *QueueSetStorage
}

type CommandHandlerSetStorage func(job *JobSetStorage)

func NewQueueSetStorage(bufferSize int) *QueueSetStorage {
	return &QueueSetStorage{
		jobs: make(chan *JobSetStorage, bufferSize),
	}
}

func (q *QueueSetStorage) Add(job *JobSetStorage) {
	q.jobs <- job
}

func (ss *SetStorage) SetAdd(job *JobSetStorage) {

	ss.mu.Lock()

	defer ss.mu.Unlock()

	_, exists := ss.data[job.Key]

	if !exists {
		ss.data[job.Key] = make(map[string]struct{})
	}

	ss.data[job.Key][job.Value] = struct{}{}

	job.Response <- "OK"
}

func (ss *SetStorage) SetMembers(job *JobSetStorage) {

	ss.mu.RLock()

	defer ss.mu.RUnlock()

	set, exists := ss.data[job.Key]

	if !exists {
		job.Response <- nil
		return
	}

	members := make([]string, 0, len(set))

	for member := range set {
		members = append(members, member)
	}

	job.Response <- members
}

func NewSetStorage(bufferSize int) *SetStorage {

	storage := &SetStorage{
		data:  make(map[string]map[string]struct{}),
		queue: NewQueueSetStorage(bufferSize),
	}

	go storage.Process()

	return storage
}

func (ss *SetStorage) Process() {

	commandHandlers := map[string]CommandHandlerSetStorage{
		"SETADD":     ss.SetAdd,
		"SETMEMBERS": ss.SetMembers,
	}

	for job := range ss.queue.jobs {

		handler, exists := commandHandlers[job.Command]

		if exists {
			handler(job)
		} else {
			job.Response <- "ERROR: Unknown command"
		}
	}
}

func (ss *SetStorage) AddJobSetAdd(key, value string) string {

	job := &JobSetStorage{
		Command:  "SETADD",
		Key:      key,
		Value:    value,
		Response: make(chan interface{}),
	}

	ss.queue.Add(job)

	return (<-job.Response).(string)
}

func (ss *SetStorage) AddJobSetMembers(key string) ([]string, bool) {

	job := &JobSetStorage{
		Command:  "SETMEMBERS",
		Key:      key,
		Response: make(chan interface{}),
	}

	ss.queue.Add(job)

	result := <-job.Response

	if result == nil {
		return nil, false
	}

	return result.([]string), true
}

var Ss *SetStorage = NewSetStorage(10)

func FormatMembers(members []string) string {
	result := ""
	for i, member := range members {
		if i > 0 {
			result += ", "
		}
		result += member
	}
	return result
}

func HandlerSetStorage(command string, args []string) <-chan string {

	response := make(chan string)

	go func() {
		defer close(response)

		switch command {
		case "SETADD":

			if valid, errMsg := ValidateArgs(command, args, 2); !valid {
				response <- errMsg
				return
			}

			result := Ss.AddJobSetAdd(args[0], args[1])

			response <- result

		case "SETMEMBERS":

			if valid, errMsg := ValidateArgs(command, args, 1); !valid {
				response <- errMsg
				return
			}

			members, exists := Ss.AddJobSetMembers(args[0])

			if !exists {
				response <- "NOT FOUND"
				return
			}

			response <- "Members: " + FormatMembers(members)

		default:
			response <- "ERROR: Unknown command"
		}
	}()

	return response
}
