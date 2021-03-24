package main

import (
	"fmt"

	"github.com/alinush/go-mcl"
)

// import (
// 	"github.com/alinush/go-mcl"
// )

func groth16_G1(constraintsCount, numVars, statementSize uint64) {

}

func groth16_G2(constraintsCount, numVars, statementSize uint64) {

}

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

func main() {
	fmt.Println("Hello, World!")
	groth16_G1(112, 20, 2)
}
