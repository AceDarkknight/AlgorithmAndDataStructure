package skipList

import (
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
	skipList := NewSkipList(4)
	for i := 1; i <= 10; i++ {
		skipList.Insert(uint64(i), i)
	}

	t.Run("test", func(t *testing.T) {
		if skipList.Length() != 10 || skipList.Level() != 4 {
			t.Errorf("skip list error length or level are not correct")
		}
	})

	type args struct {
		index uint64
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"test1", args{1}, 1},
		{"test2", args{5}, 5},
		{"test1", args{20}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := skipList.Search(tt.args.index); got != tt.want {
				t.Errorf("Search() = %v, want = %v", got, tt.want)
			}
		})
	}

}
