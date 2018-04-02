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

// Insert will insert a node into skip list. If skip has these this index, overwrite the value, otherwise add it.
func (s *SkipList) Insert(index int64, value interface{}) {
	if value == nil {
		return
	}

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

// Delete will find the index is existed or not firstly. If existed, delete it, otherwise do nothing.
func (s *SkipList) Delete(index int64) {
}

// Search will search the skip list with the given index.
// If the index exists, return the value, otherwise return nil.
func (s *SkipList) Search(index int64) interface{} {
	_, result := s.doSearch(index)
	if result.Index != index {
		return nil
	} else {
		return result.Value
	}
}

// doSearch will search given index in skip list.
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

// ForEach will iterate the whole skip list and do the function f for each index and value.
// Function f will not change the index and value in skip list.
func (s *SkipList) ForEach(f func(index int64, value interface{}) bool) {
	currentNode := s.head.next(0)
	for currentNode != nil {
		i := currentNode.Index
		v := currentNode.Value
		if !f(i, v) {
			break
		}

		currentNode = currentNode.next(0)
	}
}

// randomLevel will generate and random level that level > 0 and level < skip list's level
// This comes from redis's implementation.
func (s *SkipList) randomLevel() int {
	level := 1
	for rand.New(rand.NewSource(time.Now().Unix())).Float64() < PROBABILITY && level < s.level {
		level++
	}

	return level
}
