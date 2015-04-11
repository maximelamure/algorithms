package datastructure

type Queue interface {
	Enqueue(obj int)
	Dequeue() int
	IsEmpty() bool
	Size() int
	Iterate() <-chan int
}
