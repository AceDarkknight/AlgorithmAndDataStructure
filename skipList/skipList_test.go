package skipList

import (
	"fmt"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	type args struct {
		level int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{-1}, []int{32, 0}},
		{"test2", args{4}, []int{4, 0}},
		{"test3", args{40}, []int{32, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSkipList(tt.args.level); got.Level() != tt.want[0] || got.Length() != int32(tt.want[1]) {
				t.Errorf("NewSkipList() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestHash(t *testing.T) {
	input := `Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
`

	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"test", args{[]byte(input)}, 0xFFAE31BEBFED7652},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash(tt.args.input); got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkipList_Insert(t *testing.T) {
	skipList := NewSkipList(8)
	type args struct {
		index uint64
	}

	t.Run("test level", func(t *testing.T) {
		if skipList.Level() != 8 {
			t.Errorf("skip list level is not correct")
		}
	})

	t.Run("test Length0", func(t *testing.T) {
		if skipList.Length() != 0 {
			t.Errorf("skip list error length are not correct")
		}
	})

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"test0", args{1}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := skipList.Search(tt.args.index); got != tt.want {
				t.Errorf("Search() = %v, want = %v", got, tt.want)
			}
		})
	}

	// Insert elements.
	for i := 2; i <= 10; i++ {
		skipList.Insert(uint64(i), i)
	}

	skipList.Insert(uint64(1), 1)
	skipList.Insert(uint64(2018), 111)

	t.Run("test Length1", func(t *testing.T) {
		if skipList.Length() != 11 {
			t.Errorf("skip list error length is not correct")
		}
	})

	tests = []struct {
		name string
		args args
		want interface{}
	}{
		{"test1", args{1}, 1},
		{"test2", args{2}, 2},
		{"test3", args{3}, 3},
		{"test4", args{4}, 4},
		{"test5", args{5}, 5},
		{"test6", args{6}, 6},
		{"test7", args{7}, 7},
		{"test8", args{8}, 8},
		{"test9", args{9}, 9},
		{"test10", args{10}, 10},
		{"test11", args{2018}, 111},
		{"test12", args{20}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := skipList.Search(tt.args.index); got != tt.want {
				t.Errorf("Search() = %v, want = %v", got, tt.want)
			}
		})
	}

	skipList.Insert(uint64(1024), nil)
	skipList.Insert(uint64(2018), 2017)
	t.Run("test length2", func(t *testing.T) {
		if got := skipList.Length(); got != 11 {
			t.Errorf("skip list's length = %d is not correct", got)
		}
	})

	tests = []struct {
		name string
		args args
		want interface{}
	}{
		{"test13", args{1}, 1},
		{"test14", args{2}, 2},
		{"test15", args{3}, 3},
		{"test16", args{4}, 4},
		{"test17", args{5}, 5},
		{"test18", args{6}, 6},
		{"test19", args{7}, 7},
		{"test20", args{8}, 8},
		{"test21", args{9}, 9},
		{"test22", args{10}, 10},
		{"test23", args{2018}, 2017},
		{"test24", args{1024}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := skipList.Search(tt.args.index); got != tt.want {
				t.Errorf("Search() = %v, want = %v", got, tt.want)
			}
		})
	}

	var lastIndex uint64 = 0
	skipList.ForEach(func(index uint64, value interface{}) bool {
		t.Run("test sequence", func(t *testing.T) {
			t.Logf("index:%v value:%v", index, value)
			if lastIndex > index {
				t.Errorf("incorrect sequence")
			}

			lastIndex = index
		})

		return true
	})
}

func TestSkipList_Delete(t *testing.T) {
	skipList := NewSkipList(16)
	skipList.Delete(uint64(1))
	type args struct {
		index uint64
	}

	t.Run("test level", func(t *testing.T) {
		if skipList.Level() != 16 {
			t.Errorf("skip list level is not correct")
		}
	})

	t.Run("test Length0", func(t *testing.T) {
		if skipList.Length() != 0 {
			t.Errorf("skip list error length are not correct")
		}
	})

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"test0", args{1}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := skipList.Search(tt.args.index); got != tt.want {
				t.Errorf("Search() = %v, want = %v", got, tt.want)
			}
		})
	}

	// Insert elements.
	for i := 2; i <= 10; i++ {
		skipList.Insert(uint64(i), i)
	}
	skipList.Insert(uint64(1), 1)

	t.Run("test Length1", func(t *testing.T) {
		if skipList.Length() != 10 {
			t.Errorf("skip list error length is not correct")
		}
	})

	// Delete elements.
	skipList.Delete(uint64(5))
	skipList.Delete(uint64(1))
	skipList.Delete(uint64(11))

	t.Run("test Length1", func(t *testing.T) {
		if skipList.Length() != 8 {
			t.Errorf("skip list error length is not correct")
		}
	})

	tests = []struct {
		name string
		args args
		want interface{}
	}{
		{"test1", args{1}, nil},
		{"test2", args{2}, 2},
		{"test3", args{3}, 3},
		{"test4", args{4}, 4},
		{"test5", args{5}, nil},
		{"test6", args{6}, 6},
		{"test7", args{7}, 7},
		{"test8", args{8}, 8},
		{"test9", args{9}, 9},
		{"test10", args{10}, 10},
		{"test11", args{11}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := skipList.Search(tt.args.index); got != tt.want {
				t.Errorf("Search() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestSkipList_Level(t *testing.T) {
	skipList := NewSkipList(32)
	number := 1000000
	for i := number; i > 0; i-- {
		skipList.Insert(uint64(i), i)
	}

	levelCount := make([]int, 33)
	currentNode := skipList.head.nextNodes[0]
	for currentNode != skipList.tail {
		levelCount[len(currentNode.nextNodes)]++
		currentNode = currentNode.nextNodes[0]
	}

	fmt.Printf("%#v\n", levelCount)
}
