package queue

type Job struct {
	ID       string
	Request  string
	Response chan string
}

type Queue struct {
	jobs chan *Job
}

func NewQueue(bufferSize int) *Queue {
	return &Queue{
		jobs: make(chan *Job, bufferSize),
	}
}

func (q *Queue) Add(job *Job) {
	q.jobs <- job
}

func (q *Queue) Process(worker func(*Job)) {
	for job := range q.jobs {
		worker(job)
	}
}
