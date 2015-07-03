package datastructure

type Trie interface {
	Put(key string, val interface{})
	Get(key string) interface{}
}
