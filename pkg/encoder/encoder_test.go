package encoder

import (
	"math/big"
	"testing"
)

func TestEncodeDecodeInt(t *testing.T) {
	var a, b int64 = 2, 3
	var sum, mul int64 = 5, 6
	var encA, encB *big.Int = EncodeInt(a), EncodeInt(b)

	var decSum int64 = DecodeInt(new(big.Int).Add(encA, encB))
	if sum != decSum {
		t.Errorf("expected %d, got %d", sum, decSum)
	}

	var resMul = new(big.Int).Mul(encA, encB)
	var decMul int64 = DecodeInt(new(big.Int).Div(resMul, Factor))
	if mul != decMul {
		t.Errorf("expected %d, got %d", mul, decMul)
	}
}

func TestEncodeDecodeFloat(t *testing.T) {
	var a, b float64 = 2.1, 3.2
	var sum, mul float64 = 5.3, 6.72
	var encA, encB *big.Int = EncodeFloat(a), EncodeFloat(b)

	var decSum float64 = DecodeFloat(new(big.Int).Add(encA, encB))
	if sum != decSum {
		t.Errorf("expected %.2f, got %.2f", sum, decSum)
	}

	var resMul *big.Int = new(big.Int).Mul(encA, encB)
	var decMul float64 = DecodeFloat(new(big.Int).Div(resMul, Factor))
	if mul != decMul {
		t.Errorf("expected %.2f, got %.2f", mul, decMul)
	}
}
