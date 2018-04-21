package queue

import "errors"

type CyclicQueue struct {
	front    *Node
	rear     *Node
	length   int
	capacity int
	nodes    []*Node
}

func NewCyclicQueue(capacity int) (*CyclicQueue, error) {
	if capacity <= 0 {
		return nil, errors.New("capacity is less than 0")
	}

	front := &Node{
		value: nil,
	}

	rear := &Node{
		value: nil,
	}

	nodes := make([]*Node, 0, capacity)
	return &CyclicQueue{
		front:    front,
		rear:     rear,
		capacity: capacity,
		nodes:    nodes,
	}, nil
}

func (q *CyclicQueue) Length() int {
	return q.length
}

func (q *CyclicQueue) Capacity() int {
	return q.capacity
}

func (q *CyclicQueue) Front() *Node {
	if q.length == 0 {
		return nil
	}

	return q.front.next
}

func (q *CyclicQueue) Rear() *Node {
	if q.length == 0 {
		return nil
	}

	return q.rear.previous
}

func (q *CyclicQueue) Enqueue(value interface{}) bool {
	if q.length == q.capacity || value == nil {
		return false
	}

	node := &Node{
		value: value,
	}

	return true
}

func (q *CyclicQueue) Dequeue() interface{} {
	if q.length == 0 {
		return nil
	}
}
