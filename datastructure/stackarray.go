package datastructure

type stackArray struct {
	Items []string
	Index int
}

func NewStackArray() Stack {
	obj := &stackArray{}
	obj.Items = make([]string, 0)
	return obj
}

func (s *stackArray) Push(item string) {
	//Here we can use append. append is obtimized to set the capactity
	// by 2 x the length when the slice is full
	s.Items = append(s.Items, item)
	s.Index++
}

func (s *stackArray) Pop() string {
	if !s.IsEmpty() {
		item := s.Items[s.Index-1]
		s.Index--
		s.Resize()
		return item
	}

	return ""
}

func (s *stackArray) Resize() {
	if cap(s.Items)/4 > s.Index {
		s.Items = s.Items[0 : cap(s.Items)/2]
	}
}

func (s *stackArray) IsEmpty() bool {
	return s.Size() == 0
}

func (s *stackArray) Size() int {
	return s.Index
}
