package queue

import (
	"fmt"
	"reflect"
	"testing"
)

func NewQueue(name string, capacity int) Queue {
	var q Queue
	switch name {
	case "normal":
		q, _ = NewNormalQueue(capacity)
	case "unique":
		q, _ = NewUniqueQueue(capacity)
	case "cyclic":
		q, _ = NewCyclicQueue(capacity)
	}

	return q
}

func TestQueue_Length(t *testing.T) {
	type args struct {
		name     string
		capacity int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test normal", args{"normal", -1}},
		{"test unique", args{"unique", -1}},
		{"test cyclic", args{"cyclic", -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue(tt.args.name, tt.args.capacity); !reflect.ValueOf(got).IsNil() {
				t.Errorf("got not nil")
			}
		})
	}
}

func TestQueue_Front(t *testing.T) {

}

func TestQueue_Rear(t *testing.T) {

}

func TestQueue_Dequeue(t *testing.T) {
	type args struct {
		name     string
		capacity int
		value    []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// Normal queue.
		{"test normal", args{"normal", 10, []interface{}{0}}, []interface{}{0}},
		{"test normal", args{"normal", 10, []interface{}{666, 55, 4, -3}}, []interface{}{666, 55, 4, -3, nil}},
		{"test normal", args{"normal", 10, []interface{}{nil, nil}}, []interface{}{nil, nil, nil}},
		{"test normal", args{"normal", 3, []interface{}{0, 1, 2, 3, 4}}, []interface{}{0, 1, 2, nil, nil}},
		// Unique queue.
		{"test unique", args{"unique", 10, []interface{}{0}}, []interface{}{0}},
		{"test unique", args{"unique", 10, []interface{}{666, 55, 4, -3}}, []interface{}{666, 55, 4, -3, nil}},
		{"test unique", args{"unique", 10, []interface{}{nil, nil}}, []interface{}{nil, nil, nil}},
		{"test unique", args{"unique", 3, []interface{}{0, 1, 2, 3, 4}}, []interface{}{0, 1, 2, nil, nil}},
		{"test unique", args{"unique", 3, []interface{}{0, 0, 0, 0, 0}}, []interface{}{0, nil, nil, nil, nil}},
		{"test unique", args{"unique", 4, []interface{}{int32(0), int64(0), int8(0), 0, 0}}, []interface{}{int32(0), int64(0), int8(0), 0, nil}},
		{"test unique", args{"unique", 3, []interface{}{[]int{0, 0}, []int{0, 0}, 0, 1, "1"}}, []interface{}{0, 1, "1", nil, nil}},
		{"test unique", args{"unique", 3, []interface{}{0, 1, 0, 2, 0}}, []interface{}{0, 1, 2, nil, nil}},
		// Cyclic queue.
		{"test cyclic", args{"cyclic", 10, []interface{}{0}}, []interface{}{0}},
		{"test cyclic", args{"cyclic", 10, []interface{}{666, 55, 4, -3}}, []interface{}{666, 55, 4, -3, nil}},
		{"test cyclic", args{"cyclic", 10, []interface{}{nil, nil}}, []interface{}{nil, nil, nil}},
		{"test cyclic", args{"cyclic", 3, []interface{}{0, 1, 2, 3, 4}}, []interface{}{0, 1, 2, nil, nil}},
	}
	for _, tt := range tests {
		q := NewQueue(tt.args.name, tt.args.capacity)
		for _, v := range tt.args.value {
			q.Enqueue(v)
		}

		for i, w := range tt.want {
			t.Run(fmt.Sprintf(tt.name+"%d", i+1), func(t *testing.T) {
				if got := q.Dequeue(); got != w {
					t.Errorf("Dequeue() got %v, want %v", got, w)
				}
			})
		}
	}
}

func TestQueue_Enqueue(t *testing.T) {
	type args struct {
		name     string
		capacity int
		value    []interface{}
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// Normal queue.
		{"test normal", args{"normal", 10, []interface{}{0}}, []bool{true}},
		{"test normal", args{"normal", 10, []interface{}{666, 55, 4, -3}}, []bool{true, true, true, true}},
		{"test normal", args{"normal", 10, []interface{}{nil, nil}}, []bool{false, false}},
		{"test normal", args{"normal", 3, []interface{}{0, 1, 2, 3, 4}}, []bool{true, true, true, false, false}},
		// Unique queue.
		{"test unique", args{"unique", 10, []interface{}{0}}, []bool{true}},
		{"test unique", args{"unique", 10, []interface{}{666, 55, 4, -3}}, []bool{true, true, true, true}},
		{"test unique", args{"unique", 10, []interface{}{nil, nil}}, []bool{false, false}},
		{"test unique", args{"unique", 3, []interface{}{0, 1, 2, 3, 4}}, []bool{true, true, true, false, false}},
		{"test unique", args{"unique", 3, []interface{}{0, 0, 0, 0, 0}}, []bool{true, false, false, false, false}},
		{"test unique", args{"unique", 4, []interface{}{int32(0), int64(0), int8(0), 0, 0}}, []bool{true, true, true, true, false}},
		{"test unique", args{"unique", 3, []interface{}{[]int{0, 0}, []int{0, 0}, 0, 1, "1"}}, []bool{false, false, true, true, true}},
	}
	for _, tt := range tests {
		q := NewQueue(tt.args.name, tt.args.capacity)
		for i := 0; i < len(tt.args.value); i++ {
			t.Run(fmt.Sprintf(tt.name+"%d", i+1), func(t *testing.T) {
				if got := q.Enqueue(tt.args.value[i]); got != tt.want[i] {
					t.Errorf("Eequeue() got %v, want %v", got, tt.want[i])
				}
			})
		}
	}
}
