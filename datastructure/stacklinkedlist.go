package datastructure

type stackLinkedList struct {
	Current *node
	Lenght  int
}

type node struct {
	Item int
	Next *node
}

func NewStackLinkedList() Stack {
	return &stackLinkedList{}
}

func (s *stackLinkedList) Push(obj int) {
	newNode := &node{}
	newNode.Item = obj
	newNode.Next = s.Current
	s.Current = newNode
	s.Lenght++
}

func (s *stackLinkedList) Pop() int {
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

func (s *stackLinkedList) Iterate() <-chan int {
	ch := make(chan int)
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
