package queue

import "errors"

type CyclicQueue struct {
	front    int
	rear     int
	length   int
	capacity int
	nodes    []*Node
}

func NewCyclicQueue(capacity int) (*CyclicQueue, error) {
	if capacity <= 0 {
		return nil, errors.New("capacity is less than 0")
	}

	nodes := make([]*Node, capacity, capacity)

	return &CyclicQueue{
		front:    -1,
		rear:     -1,
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

	return q.nodes[q.front]
}

func (q *CyclicQueue) Rear() *Node {
	if q.length == 0 {
		return nil
	}

	return q.nodes[q.rear]
}

func (q *CyclicQueue) Enqueue(value interface{}) bool {
	if q.length == q.capacity || value == nil {
		return false
	}

	node := &Node{
		value: value,
	}

	index := (q.rear + 1) % cap(q.nodes)
	q.nodes[index] = node
	q.rear = index
	q.length++

	if q.length == 1 {
		q.front = index
	}

	return true
}

func (q *CyclicQueue) Dequeue() interface{} {
	if q.length == 0 {
		return nil
	}

	result := q.nodes[q.front].value
	q.nodes[q.front] = nil
	index := (q.front + 1) % cap(q.nodes)
	q.front = index
	q.length--

	return result
}
