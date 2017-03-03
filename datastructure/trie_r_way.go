package datastructure

type TrieRWay struct {
	N    int
	Root *TRNode
}

type TRNode struct {
	Key  interface{}
	Next []*TRNode
}

func NewTrieRWay(size int) Trie {
	return &TrieRWay{
		N: size,
	}
}

func (t *TrieRWay) Put(key string, val interface{}) {
	if len(key) == 0 {
		return
	}

	t.Root = t.put(t.Root, key, val, 0)
}

func (t *TrieRWay) put(node *TRNode, key string, val interface{}, position int) *TRNode {
	if node == nil {
		node = &TRNode{
			Next: make([]*TRNode, t.N),
		}
	}
	if position == len(key)-1 {
		node.Key = val
		return node
	}
	index := key[position] - 'a'
	node.Next[index] = t.put(node.Next[index], key, val, position+1)
	return node
}

func (t *TrieRWay) Get(key string) interface{} {
	return t.get(t.Root, key, 0)
}

func (t *TrieRWay) get(node *TRNode, key string, position int) interface{} {
	if node == nil {
		return nil
	}
	
	if position == len(key)-1 {
		return node.Key
	}
	
	index := key[position] - 'a'
	return t.get(node.Next[index], key, position+1)
}
