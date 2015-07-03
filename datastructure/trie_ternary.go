package datastructure

type Trie_Ternary struct {
	Root *TTNode
}

type TTNode struct {
	Key    string
	Value  interface{}
	Left   *TTNode
	Right  *TTNode
	Middle *TTNode
}

func NewTrieTernary() Trie {
	return &Trie_Ternary{}
}

func (t *Trie_Ternary) Put(key string, val interface{}) {
	t.Root = t.put(t.Root, key, val, 0)
}

func (t *Trie_Ternary) put(node *TTNode, key string, val interface{}, position int) *TTNode {
	if node == nil {
		node = &TTNode{
			Key: string(key[position]),
		}
	}
	if position == len(key) {
		node.Value = val
		return node
	}

	cmp := (key[position] - 'a') - (node.Key[0] - 'a')
	if cmp == 0 {
		node.Middle = t.put(node.Middle, key, val, position+1)
	} else if cmp < 0 {
		node.Left = t.put(node.Left, key, val, position+1)
	} else {
		node.Right = t.put(node.Right, key, val, position+1)
	}
	return node
}

func (t *Trie_Ternary) Get(key string) interface{} {
	return t.get(t.Root, key, 0)
}

func (t *Trie_Ternary) get(node *TTNode, key string, position int) interface{} {
	if node == nil {
		return nil
	}
	if position >= len(key) {
		return node.Value
	}
	cmp := (key[position] - 'a') - (node.Key[0] - 'a')
	if cmp == 0 {
		return t.get(node.Middle, key, position+1)
	} else if cmp < 0 {
		return t.get(node.Left, key, position)
	} else {
		return t.get(node.Right, key, position)
	}
}
