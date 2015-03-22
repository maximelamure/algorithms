package datastructure

import "hash/fnv"

// Under uniform hashing assumption, prob. that the number of key in a list is
// within a constant factor of N/M is extremely close to 1.
// Number of probes for search/insert is proportional to N/M
// - M too large => too many empty chains
// - M too small => chains too long
// - Typical choice: M ~ N/5 constant times ops
// Worst case : ln N, Average case : constant (3-5)

type hashSeperateChaining struct {
	m  int
	st []*HNode
}

type HNode struct {
	Key   string
	Value string
	Next  *HNode
}

func NewHashSeperateChaining() Hash {
	ht := &hashSeperateChaining{}
	ht.m = 97
	ht.st = make([]*HNode, ht.m)
	return ht
}

func (h *hashSeperateChaining) Get(key string) string {
	hash := h.hash(key)

	for x := h.st[hash]; x != nil; x = x.Next {
		if x.Key == key {
			return x.Value
		}
	}
	return ""
}

func (h *hashSeperateChaining) hash(key string) int {
	ha := fnv.New32a()
	ha.Write([]byte(key))
	return int(ha.Sum32()&0x7fffffff) % h.m
}

func (h *hashSeperateChaining) Put(key, value string) {
	hash := h.hash(key)

	for x := h.st[hash]; x != nil; x = x.Next {
		if x.Key == key {
			x.Value = value
			return
		}
	}

	n := &HNode{}
	n.Value = value
	n.Key = key
	n.Next = h.st[hash]
	h.st[hash] = n
}
