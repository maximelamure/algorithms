package datastructure

type Stack interface {
	Push(item string)
	Pop() string
	IsEmpty() bool
	Size() int
}
