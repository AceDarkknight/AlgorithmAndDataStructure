package skipList

import (
	"math"
	"sync"
)

const SHARDS = 32

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
	skipLists []*concurrentSkipList
	length    int32
}

type concurrentSkipList struct {
	level  int
	length int32
	head   *Node
	tail   *Node
	max    *Node
	mutex  sync.RWMutex
}

func NewConcurrentSkipList(level int) *ConcurrentSkipList {
	if level <= 0 || level > MAX_LEVEL {
		level = MAX_LEVEL
	}

	skipLists := make([]*concurrentSkipList, SHARDS)
	for i := 0; i < SHARDS; i++ {
		head := NewNode(0, nil, level)
		tail := NewNode(0, nil, level)
		for i := 0; i < len(head.nextNodes); i++ {
			head.nextNodes[i] = tail
		}

		tail.previousNode = head
		head.previousNode = nil

		skipLists[i] = &concurrentSkipList{
			level:  level,
			length: 0,
			head:   head,
			tail:   tail,
		}
	}

	return &ConcurrentSkipList{
		skipLists: skipLists,
	}
}

func (s *ConcurrentSkipList) getShardIndex(index uint64) int {
	result := -1
	for i, t := range shardIndexs {
		if t > index {
			result = i
			break
		}
	}

	return result
}

// Level will return the level of skip list.
func (s *ConcurrentSkipList) Level() int {
	return s.skipLists[0].level
}

// Length will return the length of skip list.
func (s *ConcurrentSkipList) Length() int32 {
	return s.length
}

func (s *ConcurrentSkipList) Search(index uint64) *Node {

	return nil
}

func init() {
	var step uint64 = 1 << 59 // 2^64/32
	for i := SHARDS - 1; i >= 0; i-- {
		var t uint64 = math.MaxUint64
		for j := SHARDS - 1; j < i; i++ {
			t = math.MaxUint64 - step
		}

		shardIndexs[i] = t
	}
}

var shardIndexs = make([]uint64, SHARDS)
