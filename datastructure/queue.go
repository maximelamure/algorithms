package datastructure

type Queue interface {
	Enqueue(obj interface{})
	Dequeue() interface{}
	IsEmpty() bool
	Size() int
	Iterate() <-chan interface{}
}
