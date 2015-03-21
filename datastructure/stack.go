package datastructure

type Stack interface {
	Push(item int)
	Pop() int
	IsEmpty() bool
	Size() int
}
