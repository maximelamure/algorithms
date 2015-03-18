package datastructure

//SimpleIterator is a simple struct to illustrate an implementation of an iterator
type SimpleIterator struct {
	fullNames []string
	Length    int
}

// NewSimpleIterator initialize a new SimpleIterator
func NewSimpleIterator() *SimpleIterator {
	si := &SimpleIterator{}
	si.fullNames = []string{"Maxime", "Florie", "Louise", "Victoire"}
	si.Length = 4
	return si
}

// Names return an Iterator by using a chanel
func (s *SimpleIterator) Names() <-chan string {
	ch := make(chan string)
	go func() {
		for _, v := range s.fullNames {
			ch <- v
		}
		close(ch)
	}()
	return ch
}
