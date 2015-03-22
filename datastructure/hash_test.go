package datastructure

import (
	"testing"

	"github.com/maximelamure/api/common"
)

func TestHashLinear(t *testing.T) {
	h := NewHashLinearProbing()
	testHash(h, t)
}

func TestHashSeparate(t *testing.T) {
	h := NewHashSeperateChaining()
	testHash(h, t)
}

func testHash(h Hash, t *testing.T) {
	helper := common.Test{}

	h.Put("maxime", "lamure")
	h.Put("florie", "perouze")

	v1 := h.Get("maxime")
	helper.Assert(t, v1 == "lamure", "The value should be lamure. Here: "+v1)
	h.Put("florie", "lamure")
	v1 = h.Get("florie")
	helper.Assert(t, v1 == "lamure", "The value should be lamure. Here: "+v1)
}
