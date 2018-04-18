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
	case "cyclic":
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

func TestNormalQueue_Enqueue(t *testing.T) {
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
		{"test normal", args{"normal", 10, []interface{}{0}}, []interface{}{0}},
		{"test normal", args{"normal", 10, []interface{}{666, 55, 4, -3}}, []interface{}{666, 55, 4, -3, nil}},
		{"test normal", args{"normal", 10, []interface{}{nil, nil}}, []interface{}{nil, nil, nil}},
		{"test normal", args{"normal", 3, []interface{}{0, 1, 2, 3, 4}}, []interface{}{0, 1, 2, nil, nil}},
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

func TestQueue_Dequeue(t *testing.T) {

}
