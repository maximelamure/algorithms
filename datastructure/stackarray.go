package datastructure

type stackArray struct {
	Items []interface{}
	Index int
}

func NewStackArray() Stack {
	obj := &stackArray{}
	obj.Items = make([]interface{}, 0)
	return obj
}

func (s *stackArray) Push(item interface{}) {
	s.Resize()
	s.Items[s.Index] = item
	s.Index++
}

func (s *stackArray) Pop() interface{} {
	if !s.IsEmpty() {
		item := s.Items[s.Index-1]
		s.Index--
		s.Resize()
		return item
	}

	return 0 //FIXME
}

func (s *stackArray) Resize() {
	if cap(s.Items)/4 > s.Index {
		s.Items = s.Items[0 : cap(s.Items)/2]
	} else if cap(s.Items) == s.Index {
		c := make([]interface{}, 1+s.Index*2)
		copy(c, s.Items)
		s.Items = c
	}
}

func (s *stackArray) IsEmpty() bool {
	return s.Size() == 0
}

func (s *stackArray) Size() int {
	return s.Index
}

func (s *stackArray) Iterate() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for {
			if s.IsEmpty() {
				break
			}
			ch <- s.Pop()
		}
		close(ch)

	}()
	return ch
}
