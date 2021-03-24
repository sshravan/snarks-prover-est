package main

import (
	"fmt"
	"testing"

	"github.com/alinush/go-mcl"
)

type Input struct {
	// id               uint8
	constraintsCount uint64
	numVars          uint64
	statementSize    uint64
}

func inputs() []Input {
	data := []Input{
		// {100_000, 100_000, 2},
		// {500_000, 500_000, 2},
		// {1_000_000, 1_000_000, 2},
		{7_648_002, 7_648_002, 2},
		{14_018_002, 14_018_002, 2},
		{20_388_002, 20_388_002, 2},
		{21_964_002, 21_964_002, 2},
	}
	return data
}

func BenchmarkGroth16(b *testing.B) {

	testcases := inputs()
	for i := 0; i < len(testcases); i++ {
		aCount := testcases[i].constraintsCount
		bCount := testcases[i].numVars + testcases[i].statementSize
		// cCount := testcases[i].constraintsCount - testcases[i].statementSize
		maxCount := max(aCount, bCount)

		baseG1 := generateG1(maxCount)
		baseG2 := generateG2(maxCount)
		expoFr := generateFr(maxCount)
		fmt.Println("Done generating the data")

		b.Run(fmt.Sprintf("%d/N_G1;", i), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				var result mcl.G1
				mcl.G1MulVec(&result, baseG1[:aCount], expoFr[:aCount])
			}
		})

		// b.Run(fmt.Sprintf("%d/M+n_G1;", i), func(t *testing.B) {
		// 	for i := 0; i < t.N; i++ {
		// 		var result mcl.G1
		// 		mcl.G1MulVec(&result, baseG1[:bCount], expoFr[:bCount])
		// 	}
		// })

		// b.Run(fmt.Sprintf("%d/N-n_G1;", i), func(t *testing.B) {
		// 	for i := 0; i < t.N; i++ {
		// 		var result mcl.G1
		// 		mcl.G1MulVec(&result, baseG1[:cCount], expoFr[:cCount])
		// 	}
		// })

		b.Run(fmt.Sprintf("%d/N_G2;", i), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				var result mcl.G2
				mcl.G2MulVec(&result, baseG2[:aCount], expoFr[:aCount])
			}
		})
	}
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	} else {
		return b
	}
}
