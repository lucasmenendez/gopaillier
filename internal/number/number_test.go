package number

import (
	"math/big"
	"testing"
)

func TestSet(t *testing.T) {
	var value = big.NewInt(123)
	var exp = big.NewInt(3)

	var expected = &Number{value, exp, false}
	var result = new(Number).Set(expected)
	if expected.Value.Cmp(result.Value) != 0 {
		t.Fatalf("expected %d, got %d", expected.Value, result.Value)
	} else if expected.Exp.Cmp(result.Exp) != 0 {
		t.Fatalf("expected %d, got %d", expected.Exp, result.Exp)
	} else if expected.encrypted != result.encrypted {
		t.Fatalf("expected %t, got %t", expected.encrypted, result.encrypted)
	}
}

func TestSetEncrypted(t *testing.T) {
	var num = new(Number)
	if num.encrypted != false || num.IsEncrypted() != false {
		t.Fatalf("expected false, got %t", num.encrypted)
	}

	num.SetEncrypted(true)
	if num.encrypted != true || num.IsEncrypted() != true {
		t.Fatalf("expected true, got %t", num.encrypted)
	}
}

func TestSetInt(t *testing.T) {
	var A, B int64 = 12, -12400
	var aValue, aExp *big.Int = big.NewInt(12), big.NewInt(0)
	var bValue, bExp *big.Int = big.NewInt(-124), big.NewInt(2)

	var resA = new(Number).SetInt(A)
	var resB = new(Number).SetInt(B)

	if resA.Value.Cmp(aValue) != 0 {
		t.Fatalf("expected %d, got %d", aValue, resA.Value)
	}

	if resA.Exp.Cmp(aExp) != 0 {
		t.Fatalf("expected %d, got %d", aExp, resA.Exp)
	}

	if resB.Value.Cmp(bValue) != 0 {
		t.Fatalf("expected %d, got %d", bValue, resB.Value)
	}

	if resB.Exp.Cmp(bExp) != 0 {
		t.Fatalf("expected %d, got %d", bExp, resB.Exp)
	}
}

func TestSetFloat(t *testing.T) {
	var A, B float64 = -800, 12400.36
	var aValue, aExp *big.Int = big.NewInt(-8), big.NewInt(2)
	var bValue, bExp *big.Int = big.NewInt(1240036), big.NewInt(-2)

	var resA = new(Number).SetFloat(A)
	var resB = new(Number).SetFloat(B)

	if resA.Value.Cmp(aValue) != 0 {
		t.Fatalf("expected %d, got %d", aValue, resA.Value)
	}

	if resA.Exp.Cmp(aExp) != 0 {
		t.Fatalf("expected %d, got %d", aExp, resA.Exp)
	}

	if resB.Value.Cmp(bValue) != 0 {
		t.Fatalf("expected %d, got %d", bValue, resB.Value)
	}

	if resB.Exp.Cmp(bExp) != 0 {
		t.Fatalf("expected %d, got %d", bExp, resB.Exp)
	}
}

func TestInt(t *testing.T) {
	var A, B int64 = 12, -12400

	var resA = new(Number).SetInt(A).Int()
	var resB = new(Number).SetInt(B).Int()

	if A != resA {
		t.Fatalf("expected %d, got %d", A, resA)
	}

	if B != resB {
		t.Fatalf("expected %d, got %d", B, resB)
	}

	var C, D float64 = 0.125, -12400.36
	var expC, expD int64 = 0, -12401

	var resC = new(Number).SetFloat(C).Int()
	var resD = new(Number).SetFloat(D).Int()

	if expC != resC {
		t.Fatalf("expected %d, got %d", expC, resC)
	}

	if expD != resD {
		t.Fatalf("expected %d, got %d", expD, resD)
	}
}

func TestFloat(t *testing.T) {
	var A, B int64 = 12, -12400
	var expA, expB float64 = 12, -12400

	var resA = new(Number).SetInt(A).Float()
	var resB = new(Number).SetInt(B).Float()

	if expA != resA {
		t.Fatalf("expected %.5f, got %.5f", expA, resA)
	}

	if expB != resB {
		t.Fatalf("expected %.5f, got %.5f", expB, resB)
	}

	var C, D float64 = 0.125, -12400.36

	var resC = new(Number).SetFloat(C).Float()
	var resD = new(Number).SetFloat(D).Float()

	if C != resC {
		t.Fatalf("expected %.5f, got %.5f", C, resC)
	}

	if D != resD {
		t.Fatalf("expected %.5f, got %.5f", D, resD)
	}
}
