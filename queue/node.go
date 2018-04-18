package queue

type Node struct {
	value    interface{}
	previous *Node
	next     *Node
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) Set(value interface{}) {
	n.value = value
}

func (n *Node) Previous() *Node {
	return n.previous
}

func (n *Node) Next() *Node {
	return n.next
}
