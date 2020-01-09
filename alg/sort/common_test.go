package sort

import (
	"math/rand"
	"testing"
)

func randomSlice(b *testing.B) (ret []int) {
	b.StopTimer()
	defer b.StartTimer()
	ret = make([]int, 1000)
	for i := range ret {
		ret[i] = rand.Intn(100)
	}
	return
}
