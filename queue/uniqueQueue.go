package queue

import "errors"

type UniqueQueue struct {
	queue    *NormalQueue
	valueMap map[interface{}]bool
}

func NewUniqueQueue(capacity int) (*UniqueQueue, error) {
	if capacity <= 0 {
		return nil, errors.New("capacity is less than 0")
	}

	front := &Node{
		value:    nil,
		previous: nil,
	}

	rear := &Node{
		value:    nil,
		previous: front,
	}

	front.next = rear
}
