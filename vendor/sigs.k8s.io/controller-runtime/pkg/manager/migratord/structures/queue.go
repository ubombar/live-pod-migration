package structures

type Queue interface {
	Push(obj interface{})
	Pop() interface{}
	Cap() int
	Len() int
	Name() string
}

func NewQueue(name string, size int) *queue {
	q := &queue{
		queue: make(chan interface{}, size),
		name:  name,
	}

	return q
}

type queue struct {
	Queue
	queue chan interface{}
	name  string
}

// Push a new interface to the queue. If full then blocks.
func (q *queue) Push(obj interface{}) {
	q.queue <- obj
}

// Pop interface from the queue. If empty then blocks.
func (q *queue) Pop() interface{} {
	return <-q.queue
}

// Get number of elements in the queue.
func (q *queue) Len() int {
	return len(q.queue)
}

// Get the size of the queue.
func (q *queue) Cap() int {
	return cap(q.queue)
}

// Get the name of the queue.
func (q *queue) Name() string {
	return q.name
}
