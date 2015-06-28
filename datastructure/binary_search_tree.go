package datastructure

import "fmt"

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
	Count int
	Level int
}

func NewBNode(key, val, count int) *BNode {
	return &BNode{Key: key, Value: val, Count: count}
}

func (b *binarySearchTree) SizeNode(node *BNode) int {
	if node == nil {
		return 0
	}
	return node.Count
}

func (b *binarySearchTree) Get(key int) int {

	x := b.Root
	for {
		if x == nil {
			return -1 //FIXME
		}

		if x.Key == key {
			return x.Value
		} else if x.Key > key {
			x = x.Left
		} else {
			x = x.Right
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
		return NewBNode(key, val, 1)
	}
	if key > root.Key {
		root.Right = b.put(root.Right, key, val)
	} else if key < root.Key {
		root.Left = b.put(root.Left, key, val)
	} else {
		root.Value = val
	}
	root.Count = 1 + b.SizeNode(root.Left) + b.SizeNode(root.Right)
	return root
}

func (b *binarySearchTree) Min() int {
	min := b.min(b.Root)
	if min != nil {
		return min.Value
	}
	return 0 //FIXME
}
func (b *binarySearchTree) min(node *BNode) *BNode {
	if node == nil {
		return nil
	}

	for {
		if node.Left != nil {
			node = node.Left
		} else {
			break
		}
	}
	return node
}

func (b *binarySearchTree) Max() int {
	if b.Root == nil {
		return -1 // FIXME: manage error
	}

	node := b.Root
	for {
		if node.Right == nil {
			break
		}
		node = node.Right
	}

	return node.Value
}

// Floor returns the largest key <= to a given key
func (b *binarySearchTree) Floor(key int) int {
	node := b.floor(b.Root, key)
	if node == nil {
		return -1 // FIXME: manage error
	}
	return node.Key
}

func (b *binarySearchTree) floor(node *BNode, key int) *BNode {
	if node == nil || node.Key == key {
		return node
	}

	if key < node.Key {
		return b.floor(node.Left, key)
	} else {
		if t := node.Right; t != nil {
			return t
		}
		return node
	}
}

func (b *binarySearchTree) Size() int {
	return b.SizeNode(b.Root)
}

// Rank returns the number of key < k
func (b *binarySearchTree) Rank(key int) int {
	return b.rank(b.Root, key)
}

func (b *binarySearchTree) rank(node *BNode, key int) int {
	if node == nil {
		return 0
	}

	if key == node.Key {
		return b.SizeNode(node.Left)
	} else if key < node.Key {
		return b.rank(node.Left, key)
	} else {
		return 1 + b.SizeNode(node.Left) + b.rank(node.Right, key)
	}
}

//In order traversal
func (b *binarySearchTree) InOrder() <-chan int {
	q := NewQueueLinkedList()
	b.inOrder(b.Root, q)
	ch := make(chan int)
	go func() {
		for {
			if q.IsEmpty() {
				break
			}
			ch <- q.Dequeue().(int)
		}

		close(ch)

	}()
	return ch
}

func (b *binarySearchTree) inOrder(node *BNode, q Queue) {
	if node != nil {
		b.inOrder(node.Left, q)
		q.Enqueue(node.Value)
		b.inOrder(node.Right, q)
	}
}

func (b *binarySearchTree) DeleteMin() {
	b.Root = b.deleteMin(b.Root)
}

func (b *binarySearchTree) deleteMin(node *BNode) *BNode {
	if node == nil {
		return nil
	}
	if node.Left == nil {
		return node.Right
	}
	node.Left = b.deleteMin(node.Left)
	node.Count = 1 + b.SizeNode(node.Left) + b.SizeNode(node.Right)
	return node
}

func (b *binarySearchTree) HibbardDeletion(key int) {
	b.Root = b.hibbardDeletion(b.Root, key)
}

func (b *binarySearchTree) hibbardDeletion(node *BNode, key int) *BNode {
	if node == nil {
		return nil
	}

	//BST to find the key
	if key < node.Key {
		node.Left = b.hibbardDeletion(node.Left, key)
	} else if key > node.Key {
		node.Right = b.hibbardDeletion(node.Right, key)
	} else { // equality
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		t := node
		node = b.min(t.Right)
		node.Right = b.deleteMin(t.Right)
		node.Left = t.Left
	}
	node.Count = 1 + b.SizeNode(node.Left) + b.SizeNode(node.Right)
	return node
}

func (b *binarySearchTree) String() {
	for x := range b.InOrder() {
		fmt.Print(x)
	}
	fmt.Println("")
}
