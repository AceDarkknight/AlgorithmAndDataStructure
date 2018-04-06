package skipList

import "sync/atomic"

type Node struct {
	index        uint64
	value        interface{}
	previousNode *Node
	nextNodes    []*Node
}

func NewNode(index uint64, value interface{}, level int) *Node {
	if level <= 0 || level > MAX_LEVEL {
		level = MAX_LEVEL
	}

	return &Node{
		index:     index,
		value:     value,
		nextNodes: make([]*Node, level),
	}
}

func (n *Node) Index() uint64 {
	return n.index
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) Next() *Node {
	return n.nextNodes[0]
}

func (n *Node) Previous() *Node {
	return n.previousNode
}

type ConcurrentSkipList struct {
	level  int
	length int32
	head   *Node
	tail   *Node
}

func NewConcurrentSkipList(level int) *ConcurrentSkipList {
	if level <= 0 || level > MAX_LEVEL {
		level = MAX_LEVEL
	}

	head := NewNode(0, nil, level)
	tail := NewNode(0, nil, level)
	for i := 0; i < len(head.nextNodes); i++ {
		head.nextNodes[i] = tail
	}

	tail.previousNode = head
	head.previousNode = nil

	return &ConcurrentSkipList{
		level:  level,
		length: 0,
		head:   head,
		tail:   tail,
	}
}

// Level will return the level of skip list.
func (s *ConcurrentSkipList) Level() int {
	return s.level
}

// Length will return the length of skip list.
func (s *ConcurrentSkipList) Length() int32 {
	return s.length
}

func (s *ConcurrentSkipList) Search(index uint64) *Node {
	if atomic.LoadInt32(&s.length) == 0 {
		return nil
	}

	currentNode := s.head

	// Iterate from top level to bottom level.
	for l := s.level - 1; l >= 0; l-- {
		// Iterate node util node's index is >= given index.
		// The max iterate count is skip list's length. So the worst O(n) is N.
		for currentNode.nextNodes[l] != s.tail && currentNode.nextNodes[l].index < index {
			currentNode = currentNode.nextNodes[l]
		}

		if currentNode.nextNodes[l].index == index {
			break
		}
	}

	return currentNode
}
