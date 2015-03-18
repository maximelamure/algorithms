package unionfind

import (
	"testing"

	"github.com/maximelamure/algorithms/common"
)

func TestWeightedQU(t *testing.T) {
	qf := NewWeightedQU2(10)
	qf.Union(1, 4)
	qf.Union(0, 3)
	helper := common.Test{}
	helper.Assert(t, qf.Connected(1, 4), "4 and 1 should be connected")
	helper.Assert(t, !qf.Connected(1, 3), "3 and 1 should not be connected")
	qf.Union(1, 3)
	helper.Assert(t, qf.Connected(1, 3), "3 and 1 should not be connected")
}
