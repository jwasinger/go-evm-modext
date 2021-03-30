package evmmax_arith

import (
	"testing"
)

// test converting 3 to montgomery and back using MulModMont
func TestMulModMontUnrolled_8limbs(t *testing.T) {
	test_val := make([]uint64, 8, 8)
	test_val[0] = 3

	intermediate_val := make([]uint64, 8, 8)

	mod := MaxModulus(8)
	mont_const := MontConstant_Interleaved(mod)
	r_squared := RSquared(mod)

	_MulModMont_8Limbs(intermediate_val, test_val, r_squared, mod, mont_const)

	if LimbsEq(test_val, intermediate_val) {
		t.Fatalf("conversion to montgomery should modify test val")
	}

	_MulModMont_8Limbs(intermediate_val, intermediate_val, One(8), mod, mont_const)

	if !LimbsEq(test_val, intermediate_val) {
		t.Fatalf("conversion back from montgomery failed")
	}
}

func TestAddMod_8limbs(t *testing.T) {

}

func TestSubMod_8limbs(t *testing.T) {

}

// test converting 3 to montgomery and back using MulModMont
func BenchmarkMulModMontUnrolled_8limbs(b *testing.B) {
	test_val := make([]uint64, 8)
	test_val[0] = 3

	intermediate_val := [8]uint64{0, 0, 0, 0, 0, 0, 0, 0}

	mod := MaxModulus(8)
	mont_const := MontConstant_Interleaved(mod)
	r_squared := RSquared(mod)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_MulModMont_8Limbs(intermediate_val[:], test_val, r_squared[:], mod[:], mont_const)
	}
}
