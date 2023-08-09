package consumers

import (
	"github.com/sirupsen/logrus"
	"github.com/ubombar/live-pod-migration/pkg/migratord/structures"
)

type Consumer interface {
	Run() error
	Consume(interface{}) (error, bool)
}

type ConsumerCallback func(id string) error

type consumer struct {
	callback ConsumerCallback
	queue    structures.Queue
}

func NewConsumer(queue structures.Queue, callback ConsumerCallback) *consumer {
	c := &consumer{
		callback: callback,
		queue:    queue,
	}

	return c
}

// Always consume from the queue and invoke the callback
func (c *consumer) Run() error {
	go func() {
		var exit bool = false
		var err error
		// Consume in an infinite loop
		for !exit {

			// Blocking
			obj := c.queue.Pop()
			err, exit = c.Consume(obj)

			if err != nil {
				logrus.Errorln(err)
			}
		}
	}()
	return nil
}

// Consumes the object. Returns true if we want to exit.
func (c *consumer) Consume(obj interface{}) (error, bool) {
	if id, ok := obj.(string); ok {
		// Call the callback
		return c.callback(id), false
	} else {
		logrus.Warningln("consumer received a non-string object")
		return nil, false
	}
}
