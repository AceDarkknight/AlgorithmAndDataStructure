package queue

import "errors"

type NormalQueue struct {
	front    *Node
	rear     *Node
	length   int
	capacity int
}

func NewNormalQueue(capacity int) (*NormalQueue, error) {
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
	return &NormalQueue{
		front:    front,
		rear:     rear,
		capacity: capacity,
	}, nil
}

func (q *NormalQueue) Length() int {
	return q.length
}

func (q *NormalQueue) Capacity() int {
	return q.capacity
}

func (q *NormalQueue) Front() *Node {
	if q.length == 0 {
		return nil
	}

	return q.front.next
}

func (q *NormalQueue) Rear() *Node {
	if q.length == 0 {
		return nil
	}

	return q.rear.previous
}

func (q *NormalQueue) Enqueue(value interface{}) bool {
	if q.length == q.capacity || value == nil {
		return false
	}

	node := &Node{
		value: value,
	}

	if q.length == 0 {
		q.front.next = node
	}

	node.previous = q.rear.previous
	node.next = q.rear
	q.rear.previous.next = node
	q.rear.previous = node
	q.length++

	return true
}

func (q *NormalQueue) Dequeue() interface{} {
	if q.length == 0 {
		return nil
	}

	result := q.front.next
	q.front.next = result.next
	result.next = nil
	result.previous = nil
	q.length--

	return result.value
}
