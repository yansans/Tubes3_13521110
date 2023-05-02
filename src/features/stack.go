package main

type Stack []interface{}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (value interface{}, ok bool) {
	if len(*s) == 0 {
		return nil, false
	}
	value = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value, true
}

func (s *Stack) Peek() (value interface{}, ok bool) {
	if len(*s) == 0 {
		return nil, false
	}
	return (*s)[len(*s)-1], true
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
