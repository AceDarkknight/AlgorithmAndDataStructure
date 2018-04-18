package queue

type Queue interface {
	Length() int
	Capacity() int
	Front() *Node
	Rear() *Node
	Enqueue(value interface{}) bool
	Dequeue() interface{}
}
