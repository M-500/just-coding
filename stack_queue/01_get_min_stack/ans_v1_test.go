//@Author: wulinlin
//@Description:
//@File:  ans_v1_test.go
//@Version: 1.0.0
//@Date: 2024/04/30 00:38

package _1_get_min_stack

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	s := NewMyStack()
	s.Push(3)
	s.Push(8)
	s.Push(2)
	s.Push(5)
	t.Log(s.GetMin())
	t.Log(s.Pop())
	t.Log(s.GetMin())
	t.Log(s.Pop())
	t.Log(s.GetMin())
}

func TestMyStack_GetMin(t *testing.T) {
	type fields struct {
		baseData []int
		minData  []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "没有数据",
			fields: fields{
				baseData: make([]int, 0),
				minData:  make([]int, 0),
			},
			want: -1,
		},
		{
			name: "没有数据",
			fields: fields{
				baseData: []int{5, 2, 3, 7, 1, 9, 0},
				minData:  []int{5, 2, 2, 2, 1, 1, 0}, // 构建一个单调栈
			},
			want: 0,
		},
		{
			name: "没有数据",
			fields: fields{
				baseData: []int{5, 2, 3, 7, 1, 9},
				minData:  []int{5, 2, 2, 2, 1, 1}, // 构建一个单调栈
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MyStack{
				baseData: tt.fields.baseData,
				minData:  tt.fields.minData,
			}
			if got := s.GetMin(); got != tt.want {
				t.Errorf("GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyStack_Pop(t *testing.T) {
	type fields struct {
		baseData []int
		minData  []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MyStack{
				baseData: tt.fields.baseData,
				minData:  tt.fields.minData,
			}
			if got := s.Pop(); got != tt.want {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyStack_Push(t *testing.T) {
	type fields struct {
		baseData []int
		minData  []int
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "空栈",
			fields: fields{
				baseData: nil,
				minData:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MyStack{
				baseData: tt.fields.baseData,
				minData:  tt.fields.minData,
			}
			s.Push(tt.args.data)
		})
	}
}

func TestNewMyStack(t *testing.T) {
	tests := []struct {
		name string
		want *MyStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMyStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMyStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
