package datastructure

// A BST is a binaray tree in symmetric order
// A binary tree is
// - either empty
// - two disjoinct binary tree (left and right)
// Symmetric order: each node as a key and every node's key is:
// - larger than all keys in its left subtree
// - smaller than all keys in its right subtree
// If N distincts Keys are inserted in random order, expected heigh of tree is
// ~ 4.311 ln N

type binarySearchTree struct {
	Root *BNode
}

type BNode struct {
	Key   int
	Value int
	Left  *BNode
	Right *BNode
}

func NewBNode(key, val int) *BNode {
	return &BNode{Key: key, Value: val}
}

func (b *binarySearchTree) Get(key int) *int {

	x := b.Root
	for {
		if x == nil {
			return nil
		}

		if x.Key == key {
			return &x.Value
		} else if x.Key > b.Root.Key {
			x = b.Root.Right
		} else {
			x = b.Root.Left
		}
	}
}

// Number of compares is equal 1 + depth of node.
// If N distinct keys are inserted into a BST in random order, the expceted
// number of compares for a search/insert is ~ 2 ln N (1-1 correspondence with
// quick sort paritioning)

func (b *binarySearchTree) Put(key, val int) {
	b.Root = b.put(b.Root, key, val)
}

func (b *binarySearchTree) put(root *BNode, key, val int) *BNode {
	if root == nil {
		return NewBNode(key, val)
	}
	if key > b.Root.Key {
		b.Root.Right = b.put(b.Root.Right, key, val)
	} else if key < b.Root.Key {
		b.Root.Left = b.put(b.Root.Left, key, val)
	} else {
		b.Root.Value = val
	}
	return root
}
