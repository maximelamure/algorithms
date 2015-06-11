package datastructure

type stackLinkedList struct {
	Current *node
	Lenght  int
}

type node struct {
	Item interface{}
	Next *node
}

func NewStackLinkedList() Stack {
	return &stackLinkedList{}
}

func (s *stackLinkedList) Push(obj interface{}) {
	newNode := &node{}
	newNode.Item = obj
	newNode.Next = s.Current
	s.Current = newNode
	s.Lenght++
}

func (s *stackLinkedList) Pop() interface{} {
	if !s.IsEmpty() {
		item := s.Current.Item
		s.Current = s.Current.Next
		s.Lenght--
		return item
	}
	return 0 //FIXME
}

func (s *stackLinkedList) IsEmpty() bool {
	return s.Lenght == 0
}

func (s *stackLinkedList) Size() int {
	return s.Lenght
}

func (s *stackLinkedList) Iterate() <-chan interface{} {
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
