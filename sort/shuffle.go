package sort

import (
	"math/rand"

	"github.com/maximelamure/algorithms/common"
)

type Shuffle struct {
}

func NewShuffle() *Shuffle {
	return &Shuffle{}
}

func (s *Shuffle) Sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		s := rand.Intn(i)
		common.Swap(arr, i, s)
	}
}
