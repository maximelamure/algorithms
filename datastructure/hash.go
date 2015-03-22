package datastructure

type Hash interface {
	Get(key string) string
	Put(key, value string)
}
