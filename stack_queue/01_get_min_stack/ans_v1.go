//@Author: wulinlin
//@Description:
//@File:  ans_v1
//@Version: 1.0.0
//@Date: 2024/04/30 00:25

package _1_get_min_stack

type MyStack struct {
	baseData []int
	minData  []int
}

func NewMyStack() *MyStack {
	return &MyStack{
		baseData: make([]int, 0),
		minData:  make([]int, 0),
	}
}

func (s *MyStack) Push(data int) {
	// 如果都没有数据 直接添加没有什么好说的
	if len(s.minData) == 0 {
		s.baseData = append(s.baseData, data)
		s.minData = append(s.minData, data)
		return
	}
	// 如果不为空
	tmp := s.minData[len(s.minData)-1]
	if tmp < data {
		s.minData = append(s.minData, tmp)
	} else {
		s.minData = append(s.minData, data)
	}
	// 这一行是无论如何都不能少的
	s.baseData = append(s.baseData, data)
}

func (s *MyStack) Pop() int {
	if len(s.baseData) == 0 {
		return -1
	}
	tnp := s.baseData[len(s.baseData)-1]
	s.baseData = s.baseData[:len(s.baseData)-1]
	s.minData = s.minData[:len(s.minData)-1]
	return tnp
}
func (s *MyStack) GetMin() int {
	if len(s.minData) == 0 {
		return -1
	}
	return s.minData[len(s.minData)-1]
}
