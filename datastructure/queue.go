package datastructure

type Queue interface {
	Enqueue(obj string)
	Dequeue() string
	IsEmpty() bool
	Size() int
}
