/*
Package skipList provide an implementation of skip list. But is not thread-safe in concurrency.
*/
package skipList

import (
	"math/rand"

	"github.com/OneOfOne/xxhash"
)

// Comes from redis's implementation.
// Also you can see more detail in William Pugh's paper <Skip Lists: A Probabilistic Alternative to Balanced Trees>.
// The paper is in ftp://ftp.cs.umd.edu/pub/skipLists/skiplists.pdf
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
		level = MAX_LEVEL
	}

	return &node{
		Index:     index,
		Value:     value,
		nextNodes: make([]*node, level),
	}
}

type SkipList struct {
	level  int
	length int32
	head   *node
	tail   *node
}

// NewSkipList will create and initialize a skip list with the given level.
// Level must between 1 to 32. If not, the level will set as 32.
// To determine the level, you can see the paper ftp://ftp.cs.umd.edu/pub/skipLists/skiplists.pdf.
// A simple way to determine the level is L(N) = log(1/PROBABILITY)(N).
// N is the count of the skip list which you can estimate. PROBABILITY is 0.25 in this case.
// For example, if you expect the skip list contains 10000000 elements, then N = 10000000, L(N) ≈ 12.
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

	previousNodes, currentNode := s.doSearch(index)

	// If skip list contains index, update the value.
	// Avoid to update the head.
	if currentNode != s.head && currentNode.Index == index {
		currentNode.Value = value
		return
	}

	// Make a new node.
	pendingNode := newNode(index, value, s.randomLevel())

	// Adjust pointer. Similar to update linked list.
	for i := len(pendingNode.nextNodes) - 1; i >= 0; i-- {
		// Firstly, new node point to next node.
		pendingNode.nextNodes[i] = previousNodes[i].nextNodes[i]

		// Secondly, previous nodes point to new node.
		previousNodes[i].nextNodes[i] = pendingNode
	}

	s.length++
}

// Delete will find the index is existed or not firstly. If existed, delete it, otherwise do nothing.
func (s *SkipList) Delete(index uint64) {
	previousNodes, currentNode := s.doSearch(index)

	// If skip list length is 0 or could not find node with the given index.
	if currentNode != s.head && currentNode.Index == index {
		// Adjust pointer. Similar to update linked list.
		for i := 0; i < len(currentNode.nextNodes); i++ {
			previousNodes[i].nextNodes[i] = currentNode.nextNodes[i]
			currentNode.nextNodes[i] = nil
		}

		s.length--
	}
}

// Search will search the skip list with the given index.
// If the index exists, return the value, otherwise return nil.
func (s *SkipList) Search(index uint64) interface{} {
	_, result := s.doSearch(index)
	if result == s.head || result.Index != index {
		return nil
	} else {
		return result.Value
	}
}

// doSearch will search given index in skip list.
// The first return value represents the previous nodes need to update when call Insert function.
// The second return value represents the node with given index or the closet node whose index is larger than given index.
func (s *SkipList) doSearch(index uint64) ([]*node, *node) {
	// Store all previous node whose index is less than index and whose next node's index is larger than index.
	previousNodes := make([]*node, s.level)

	// fmt.Printf("start doSearch:%v\n", index)
	currentNode := s.head

	// Iterate from top level to bottom level.
	for l := s.level - 1; l >= 0; l-- {
		// Iterate node util node's index is >= given index.
		// The max iterate count is skip list's length. So the worst O(n) is N.
		for currentNode.nextNodes[l] != s.tail && currentNode.nextNodes[l].Index < index {
			currentNode = currentNode.nextNodes[l]
		}

		// When next node's index is >= given index, add current node whose index < given index.
		previousNodes[l] = currentNode
	}

	// Avoid point to tail which will occur panic in Insert and Delete function.
	// When the next node is tail.
	// The index is larger than the maximum index in the skip list or skip list's length is 0. Don't point to tail.
	// When the next node isn't tail.
	// Next node's index must >= given index. Point to it.
	if currentNode.nextNodes[0] != s.tail {
		currentNode = currentNode.nextNodes[0]
	}
	// fmt.Printf("previous node:\n")
	// for _, n := range previousNodes {
	// 	fmt.Printf("%p\t", n)
	// }
	// fmt.Println()
	// fmt.Printf("end doSearch %v\n", index)

	return previousNodes, currentNode
}

// ForEach will iterate the whole skip list and do the function f for each index and value.
// Function f will not modify the index and value in skip list.
// Don't Insert or Delete element in ForEach function.
func (s *SkipList) ForEach(f func(index uint64, value interface{}) bool) {
	currentNode := s.head.nextNodes[0]
	for currentNode != s.tail {
		i := currentNode.Index
		v := currentNode.Value

		if !f(i, v) {
			break
		}

		currentNode = currentNode.nextNodes[0]
	}
}

// randomLevel will generate and random level that level > 0 and level < skip list's level
// This comes from redis's implementation.
func (s *SkipList) randomLevel() int {
	level := 1
	for rand.Float64() < PROBABILITY && level < s.level {
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
