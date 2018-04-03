package skipList

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/OneOfOne/xxhash"
)

// Comes from redis's implementation.
const (
	MAX_LEVEL   = 32
	PROBABILITY = 0.25
)

type node struct {
	Index     uint64
	Value     interface{}
	nextNodes []*node
}

func newNode(index uint64, value interface{}, level int) *node {
	if level <= 0 || level > MAX_LEVEL {
		return nil
	}

	return &node{
		Index:     index,
		Value:     value,
		nextNodes: make([]*node, level),
	}
}

// getNextNode will get next with given level. If level < 0 or level > next nodes' length, return nil.
func (n *node) getNextNode(level int) *node {
	if level < 0 || level > len(n.nextNodes)-1 {
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

// NewSkipList will create and initialize a skip list with the given level.
// Level must between 1 to 32. If not, the level will set as 32.
// After initialization, the head field's level equal to level parameter and point to tail field.
func NewSkipList(level int) *SkipList {
	if level <= 0 || level > MAX_LEVEL {
		level = MAX_LEVEL
	}

	head := newNode(0, nil, level)
	var tail *node = nil
	for i := 0; i < len(head.nextNodes); i++ {
		head.nextNodes[i] = tail
	}
	fmt.Printf("tail address:%p\n", tail)

	return &SkipList{
		level:  level,
		length: 0,
		head:   head,
		tail:   tail,
	}
}

// Level will return the level of skip list.
func (s *SkipList) Level() int {
	return s.level
}

// Length will return the length of skip list.
func (s *SkipList) Length() int32 {
	return s.length
}

// Insert will insert a node into skip list. If skip has these this index, overwrite the value, otherwise add it.
func (s *SkipList) Insert(index uint64, value interface{}) {
	// Ignore nil value.
	if value == nil {
		return
	}

	fmt.Printf("start insert %v %v\n", index, value)
	previousNodes, currentNode := s.doSearch(index)

	// If skip list contains index, update the value.
	if currentNode.Value != nil && currentNode.Index == index {
		currentNode.Value = value
		return
	}

	pendingNode := newNode(index, value, s.randomLevel())

	// Adjust pointer.
	for i := len(pendingNode.nextNodes) - 1; i >= 0; i-- {
		pendingNode.nextNodes[i] = previousNodes[i].nextNodes[i]
		previousNodes[i].nextNodes[i] = pendingNode
	}

	s.length++

	fmt.Printf("end insert:%+v,address:%p\n", pendingNode, pendingNode)
}

// Delete will find the index is existed or not firstly. If existed, delete it, otherwise do nothing.
func (s *SkipList) Delete(index uint64) {
}

// Search will search the skip list with the given index.
// If the index exists, return the value, otherwise return nil.
func (s *SkipList) Search(index uint64) interface{} {
	_, result := s.doSearch(index)
	if result == nil || result.Index != index {
		return nil
	} else {
		return result.Value
	}
}

// doSearch will search given index in skip list.
func (s *SkipList) doSearch(index uint64) ([]*node, *node) {
	// Store all previous node whose index is less than index and whose getNextNode node's index is larger than index.
	previousNodes := make([]*node, s.level)

	fmt.Printf("start doSearch:%v\n", index)
	currentNode := s.head
	for l := s.level - 1; l >= 0; l-- {
		if currentNode.getNextNode(l) == s.tail || currentNode.getNextNode(l).Index > index {
			previousNodes[l] = currentNode
			continue
		}

		if currentNode.getNextNode(l).Index < index {
			currentNode = currentNode.getNextNode(l)
			previousNodes[l] = currentNode
		}
	}
	fmt.Printf("previous node:\n")
	for _, n := range previousNodes {
		fmt.Printf("%p\t", n)
	}
	fmt.Println()
	fmt.Printf("end doSearch %v\n", index)

	return previousNodes, currentNode
}

// ForEach will iterate the whole skip list and do the function f for each index and value.
// Function f will not change the index and value in skip list.
func (s *SkipList) ForEach(f func(index uint64, value interface{}) bool) {
	currentNode := s.head.getNextNode(0)
	for currentNode != nil {
		i := currentNode.Index
		v := currentNode.Value
		fmt.Printf("%v\n", v)

		if !f(i, v) {
			break
		}

		currentNode = currentNode.getNextNode(0)
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

// Hash will calculate the input's hash value using xxHash algorithm.
// It can be used to calculate the index of skip list.
// See more detail in https://cyan4973.github.io/xxHash/
func Hash(input []byte) uint64 {
	h := xxhash.New64()
	h.Write(input)
	return h.Sum64()
}
