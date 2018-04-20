package queue

type CyclicQueue struct {
	front    *Node
	rear     *Node
	length   int
	capacity int
}

func NewCyclicQueue(capacity int) (*CyclicQueue, error) {

}

func (q *CyclicQueue) Length() int {
	return q.length
}

func (q *CyclicQueue) Capacity() int {
	return q.capacity
}

func (q *CyclicQueue) Front() *Node {
	return q.front.next
}

func (q *CyclicQueue) Rear() *Node {
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
