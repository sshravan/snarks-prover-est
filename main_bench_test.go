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
		{100_000, 100_000, 2},
		{500_000, 500_000, 2},
		{1_000_000, 1_000_000, 2},
		// {7_648_002, 7_648_002, 2},
		// {14_018_002, 14_018_002, 2},
		// {20_388_002, 20_388_002, 2},
		// {21_964_002, 21_964_002, 2},
	}
	return data
}

func BenchmarkExponentiation(b *testing.B) {

	var size []uint64
	size = []uint64{100_000}
	for i := 0; i < len(size); i++ {
		baseG1 := generateG1(size[i])
		baseG2 := generateG2(size[i])
		expoFr := generateFr(size[i])
		fmt.Println("Done generating the data")

		b.Run(fmt.Sprintf("%d/G1Add;", size[i]), func(t *testing.B) {
			var result mcl.G1
			result.SetString("1", 10)
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
				for j := 0; j < len(baseG1); j++ {
					mcl.G1Add(&result, &result, &baseG1[j])
				}
			}
		})
		b.Run(fmt.Sprintf("%d/G1Mul;", size[i]), func(t *testing.B) {
			var result mcl.G1
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
				for j := 0; j < len(expoFr); j++ {
					mcl.G1Mul(&result, &baseG1[j], &expoFr[j])
				}
			}
		})
		b.Run(fmt.Sprintf("%d/G1MulVec;", size[i]), func(t *testing.B) {
			var result mcl.G1
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
				mcl.G1MulVec(&result, baseG1, expoFr)
			}
		})

		b.Run(fmt.Sprintf("%d/G2Add;", size[i]), func(t *testing.B) {
			var result mcl.G2
			result.SetString("1", 10)
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
				for j := 0; j < len(baseG1); j++ {
					mcl.G2Add(&result, &result, &baseG2[j])
				}
			}
		})
		b.Run(fmt.Sprintf("%d/G2Mul;", size[i]), func(t *testing.B) {
			var result mcl.G2
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
				for j := 0; j < len(expoFr); j++ {
					mcl.G2Mul(&result, &baseG2[j], &expoFr[j])
				}
			}
		})
		b.Run(fmt.Sprintf("%d/G2MulVec;", size[i]), func(t *testing.B) {
			var result mcl.G2
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
				mcl.G2MulVec(&result, baseG2, expoFr)
			}
		})
	}
}

func BenchmarkPairing(b *testing.B) {

	var size []uint64
	size = []uint64{10_000}
	for i := 0; i < len(size); i++ {
		baseG1 := generateG1(size[i])
		baseG2 := generateG2(size[i])
		baseGT := generateGT(size[i])
		fmt.Println("Done generating the data")

		b.Run(fmt.Sprintf("%d/GTMul;", size[i]), func(t *testing.B) {
			var result mcl.GT
			result.SetString("1", 10)
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
				for j := 0; j < len(baseG1); j++ {
					mcl.GTMul(&result, &result, &baseGT[j])
				}
			}
		})

		b.Run(fmt.Sprintf("%d/MillerLoop;", size[i]), func(t *testing.B) {
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
				for j := 0; j < len(baseG1); j++ {
					mcl.MillerLoop(&baseGT[j], &baseG1[j], &baseG2[j])
				}
			}
		})

		b.Run(fmt.Sprintf("%d/FinalExp;", size[i]), func(t *testing.B) {
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
				for j := 0; j < len(baseG1); j++ {
					mcl.FinalExp(&baseGT[j], &baseGT[j])
				}
			}
		})

		b.Run(fmt.Sprintf("%d/NaivePairing;", size[i]), func(t *testing.B) {
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
				for j := 0; j < len(baseG1); j++ {
					mcl.Pairing(&baseGT[j], &baseG1[j], &baseG2[j])
				}
			}
		})

		b.Run(fmt.Sprintf("%d/MillerLoopVec;", size[i]), func(t *testing.B) {
			var result mcl.GT
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
				mcl.MillerLoopVec(&result, baseG1, baseG2)
			}
		})
	}
}

func BenchmarkGroth16(b *testing.B) {

	testcases := inputs()
	for i := 0; i < len(testcases); i++ {
		aCount := testcases[i].numVars
		bCount := testcases[i].constraintsCount + testcases[i].statementSize
		// cCount := testcases[i].numVars - testcases[i].statementSize
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
