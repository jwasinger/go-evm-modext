package evmmax_arith

import (
	"fmt"
	"math/rand"
	"testing"
)

func benchmarkMulMod(b *testing.B, limbCount int) {
	mod := MaxModulus(limbCount)
	modState, err := NewFieldContext(LimbsToBytes(mod), 256)
	if err != nil {
		panic(err)
	}
	xIdxs := make([]int, 256)
	yIdxs := make([]int, 256)
	outIdxs := make([]int, 256)
	for i := 0; i < 256; i++ {
		outIdxs[i] = rand.Intn(255)
		xIdxs[i] = rand.Intn(255)
		yIdxs[i] = rand.Intn(255)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		modState.MulMod(
			modState.scratchSpace[outIdxs[i%256]:outIdxs[i%256]+limbCount],
			modState.scratchSpace[xIdxs[i%256]:xIdxs[i%256]+limbCount],
			modState.scratchSpace[yIdxs[i%256]:yIdxs[i%256]+limbCount],
			modState.Modulus,
			modState.modInv)
	}
}

type opFn func(*testing.B, uint)

func BenchmarkOps(b *testing.B) {
	for i := 1; i <= 12; i++ {
		b.Run(fmt.Sprintf("%d-bit", i*64), func(b *testing.B) {
			benchmarkMulMod(b, i)
		})
	}
}
