package main

import (
	"finki/gol"
	"fmt"
	"math/rand"
	"time"
)

func TestMedianRandom() {
	f := gol.NewSortMedian(11)
	m := gol.NewHeapMedian(11)

	rng := rand.New(rand.NewSource(42))
	for i := 0; i < 1*3600*250; i++ {
		v := i + rng.Intn(1000)
		a := f.Filter(v)
		b := m.Filter(v)
		if m.Ready && a != b {
			println(">>> MISMATCH", i, v, a, b)
			panic("Median mismatch")
		}
	}
	fmt.Println("Median random test passed")
}

func TestMedianPerf() {
	WIN := (600 / 2) | 1
	LEN := 10 * 3600 * 250

	DATA := make([]int, LEN)
	{
		rng := rand.New(rand.NewSource(42))
		for i := 0; i < LEN; i++ {
			DATA[i] = rng.Intn(1000)
		}
	}

	var srt time.Duration
	{
		med := gol.NewSortMedian(WIN)
		start := time.Now()
		for i := 0; i < LEN; i++ {
			v := DATA[i]
			med.Filter(v)
		}
		srt = time.Since(start)
		fmt.Println("SortMedian took:", srt.Microseconds())
	}

	var hep time.Duration
	{
		med := gol.NewHeapMedian(WIN)
		start := time.Now()
		for i := 0; i < LEN; i++ {
			v := DATA[i]
			med.Filter(v)
		}
		hep = time.Since(start)
		fmt.Println("HeapMedian took:", hep.Microseconds())
	}

	fmt.Printf("Heap/Sort = %.3f%%\n", float64(hep)*100/float64(srt)-100)
}

func main() {
	TestMedianRandom()
	TestMedianPerf()
}
