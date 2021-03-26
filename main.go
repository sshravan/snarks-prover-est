package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/alinush/go-mcl"
)

func generateG1(count uint64) []mcl.G1 {
	base := make([]mcl.G1, count)
	for i := uint64(0); i < count; i++ {
		base[i].Random()
	}
	return base
}

func generateG2(count uint64) []mcl.G2 {
	base := make([]mcl.G2, count)
	for i := uint64(0); i < count; i++ {
		base[i].Random()
	}
	return base
}

func generateFr(count uint64) []mcl.Fr {
	base := make([]mcl.Fr, count)
	for i := uint64(0); i < count; i++ {
		base[i].Random()
	}
	return base
}

func generateGT(count uint64) []mcl.GT {
	N := int64(math.MaxInt64)
	var v int64
	base := make([]mcl.GT, count)
	for i := uint64(0); i < count; i++ {
		v = rand.Int63n(N)
		base[i].SetInt64(v)
	}
	return base
}

func main() {
	fmt.Println("Hello, World!")
}
