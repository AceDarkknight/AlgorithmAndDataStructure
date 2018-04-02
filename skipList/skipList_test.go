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
			if got := NewSkipList(tt.args.level); got.GetLevel() != tt.want[0] || got.GetLength() != int32(tt.want[1]) {
				t.Errorf("NewSkipList() = %v, want %v", got, tt.want)
			}
		})
	}

}
