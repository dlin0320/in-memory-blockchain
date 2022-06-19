package test

import (
	"sort"
	"testing"
	"time"

	"golang.org/x/exp/rand"
	dist "gonum.org/v1/gonum/stat/distuv"
)

func poisson() *[]float64 {
	seed := time.Now().UnixNano()
	src := rand.NewSource(uint64(seed))
	p := dist.Poisson{Lambda: 50, Src: src}
	timeList := []float64{}
	for i := 0; i < 100; i++ {
		timeList = append(timeList, p.Rand())
	}
	sort.Slice(timeList, func(i, j int) bool {
		return timeList[i] < timeList[j]
	})
	return &timeList
}

func TestPoisson(t *testing.T) {

}
