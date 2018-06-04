package stack

import "github.com/lapitsky/spreadsheet/cells"

type Item = cells.CellValue

type Stack struct {
	values []Item
}

func NewStack() *Stack {
	return &Stack{[]Item{}}
}

func (s *Stack) Push(value Item) {
	s.values = append(s.values, value)
}

func (s *Stack) Pop() (value Item, ok bool) {
	if len(s.values) == 0 {
		return cells.ErrorValue{}, false
	}
	value, s.values = s.values[len(s.values)-1], s.values[:len(s.values)-1]
	return value, true
}

func (s *Stack) Len() int {
	return len(s.values)
}
