package features

type Stack []interface{}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (value interface{}) {
	if len(*s) == 0 {
		return nil
	}
	value = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}

func (s *Stack) Peek() (value interface{}) {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
