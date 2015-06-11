package datastructure

type Stack interface {
	Push(item interface{})
	Pop() interface{}
	IsEmpty() bool
	Size() int
	Iterate() <-chan interface{}
}
