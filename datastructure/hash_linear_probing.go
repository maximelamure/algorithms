package datastructure

import "hash/fnv"

// Under uniform hashing assumption, the average # of probes in a linear
// probing hash table of size M that contains N = alpha M keys
// - M too large => too many empty array entries
// - M too small => search time blow up
// - Typical choice alpha ~ N / M ~ 1/2
// Worst case : ln N, Average case : constant (3-5)

type hashLinearProbing struct {
	m      int
	Keys   []interface{}
	Values []interface{}
}

func NewHashLinearProbing() Hash {
	h := &hashLinearProbing{}
	h.m = 30001
	h.Keys = make([]interface{}, h.m)
	h.Values = make([]interface{}, h.m)
	return h
}

func (h *hashLinearProbing) hash(key string) int {
	ha := fnv.New32a()
	ha.Write([]byte(key))
	return int(ha.Sum32()&0x7fffffff) % h.m
}

func (h *hashLinearProbing) Get(key string) string {
	i := h.hash(key)
	for ; h.Keys[i] != nil; i = (i + 1) % h.m {
		if h.Keys[i] == key {
			return h.Values[i].(string)
		}
	}
	return ""
}

func (h *hashLinearProbing) Put(key, value string) {
	i := h.hash(key)
	for ; h.Keys[i] != nil; i = (i + 1) % h.m {
		if h.Keys[i] == key {
			h.Values[i] = value
			break
		}
	}
	h.Keys[i] = key
	h.Values[i] = value

}
