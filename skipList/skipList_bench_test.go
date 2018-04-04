package skipList

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func BenchmarkSkipList_Insert_Ordered(b *testing.B) {
	skipList := NewSkipList(32)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100000)
		_ = Hash([]byte(strconv.Itoa(t)))
		skipList.Insert(uint64(i), i)
	}
}

func BenchmarkSkipList_Insert_Randomly(b *testing.B) {
	skipList := NewSkipList(32)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100000)
		index := Hash([]byte(strconv.Itoa(t)))
		skipList.Insert(index, i)
	}
}

func BenchmarkSkipList_Search_10000Elements(b *testing.B) {
	skipList := NewSkipList(32)
	for i := 0; i < 10000; i++ {
		skipList.Insert(uint64(i), i)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		skipList.Search(uint64(i))
	}
}

func BenchmarkSkipList_Search_100000Elements(b *testing.B) {
	skipList := NewSkipList(32)
	for i := 0; i < 100000; i++ {
		skipList.Insert(uint64(i), i)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		skipList.Search(uint64(i))
	}
}

func BenchmarkSkipList_Search_200000Elements(b *testing.B) {
	skipList := NewSkipList(32)
	for i := 0; i < 200000; i++ {
		skipList.Insert(uint64(i), i)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		skipList.Search(uint64(i))
	}
}

func BenchmarkSkipList_Search_500000Elements(b *testing.B) {
	skipList := NewSkipList(32)
	for i := 0; i < 500000; i++ {
		skipList.Insert(uint64(i), i)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		skipList.Search(uint64(i))
	}
}
