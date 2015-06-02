package sort

import (
	"math/rand"
	"time"

	"github.com/maximelamure/algorithms/common"
)

type Shuffle struct {
}

func NewShuffle() *Shuffle {
	return &Shuffle{}
}

func (s *Shuffle) Sort(arr []int) {
	rand.Seed(time.Now().Unix())
	for i := 1; i < len(arr); i++ {
		s := rand.Intn(i)
		common.Swap(arr, i, s)
	}
}
