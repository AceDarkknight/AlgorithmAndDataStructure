package skipList

import (
	"math/rand"
	"time"
)

const (
	MAX_LEVEL   = 32
	PROBABILITY = 0.25
)

type node struct {
	Index     int64
	Value     interface{}
	nextNodes []*node
}

func newNode(index int64, value interface{}, level int) *node {
	if level <= 0 || level > MAX_LEVEL {
		return nil
	}

	return &node{
		Index:     index,
		Value:     value,
		nextNodes: make([]*node, level),
	}
}

func (n *node) next(level int) *node {
	if level < 0 || level > len(n.nextNodes) {
		return nil
	}

	return n.nextNodes[level]
}

type SkipList struct {
	level  int
	length int32
	head   *node
	tail   *node
}

func NewSkipList(level int) *SkipList {
	if level <= 0 || level > MAX_LEVEL {
		level = MAX_LEVEL
	}

	head := newNode(0, nil, level)
	var tail *node = nil
	for i := 0; i < len(head.nextNodes); i++ {
		head.nextNodes[i] = tail
	}

	return &SkipList{
		level:  level,
		length: 0,
		head:   head,
		tail:   tail,
	}
}

func (s *SkipList) GetLevel() int {
	return s.level
}

func (s *SkipList) GetLength() int32 {
	return s.length
}

func (s *SkipList) Insert(index int64, value interface{}) {
	previousNode, currentNode := s.doSearch(index)
	if currentNode.Index == index {
		currentNode.Value = value
		return
	}

	pendingNode := newNode(index, value, s.randomLevel())

	// Adjust pointer.
	for i := 0; i < len(pendingNode.nextNodes); i++ {
		pendingNode.nextNodes[i] = previousNode[i].next(i)
		previousNode[i].nextNodes[i] = pendingNode
	}

	s.length++
}

func (s *SkipList) Delete(index int64) bool {
	return false
}

func (s *SkipList) Search(index int64) interface{} {
	if s.length <= 0 {
		return nil
	} else {
		_, result := s.doSearch(index)
		if result.Index != index {
			return nil
		} else {
			return result.Value
		}
	}
}

func (s *SkipList) doSearch(index int64) ([]*node, *node) {
	// Store all previous node whose index is less than index and whose next node's index is larger than index.
	previousNodes := make([]*node, s.level)
	currentNode := s.head

	for l := s.level; l > 0; l-- {
		// Search form top to bottom.
		for i := len(currentNode.nextNodes) - 1; i >= 0; i-- {
			// If next node's index is tail or larger than current node's index, move down.
			if currentNode.next(i) == s.tail || currentNode.next(i).Index > index {
				previousNodes[i] = currentNode
				continue
			}

			// If index does not exist, it will return the closest node whose index is less than index.
			if currentNode.next(i).Index <= index {
				currentNode = currentNode.next(i)
				break
			}
		}

		if currentNode.Index == index {
			break
		}
	}

	return previousNodes, currentNode
}

// This comes from redis's implementation.
func (s *SkipList) randomLevel() int {
	level := 1
	for rand.New(rand.NewSource(time.Now().Unix())).Float64() < PROBABILITY && level < s.level {
		level++
	}

	return level
}
