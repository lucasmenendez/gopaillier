package sdk

import (
	"math/big"

	"github.com/lucasmenendez/gopaillier/internal/number"
	"github.com/lucasmenendez/gopaillier/pkg/paillier"
)

func Add(key *paillier.PublicKey, encrypted, input *number.Number) *number.Number {
	var result = new(number.Number)
	var cmp = encrypted.Exp.Cmp(input.Exp)
	if cmp == 0 {
		result.Exp = encrypted.Exp
		result.Base = key.Add(encrypted.Base, input.Base)
		return result
	}

	var expDiff = new(big.Int).Abs(new(big.Int).Sub(encrypted.Exp, input.Exp))
	var factor = new(big.Int).Exp(big.NewInt(10), expDiff, nil)
	if cmp > 0 {
		result.Exp = input.Exp
		var normalized = key.Mul(encrypted.Base, factor)
		result.Base = key.Add(normalized, input.Base)
	} else {
		result.Exp = encrypted.Exp
		var normalized = new(big.Int).Mul(input.Base, factor)
		result.Base = key.Add(encrypted.Base, normalized)
	}

	return result
}

func Sub(key *paillier.PublicKey, encrypted, input *number.Number) *number.Number {
	var negInput = new(number.Number)
	negInput.Exp = input.Exp
	negInput.Base = new(big.Int).Neg(input.Base)
	return Add(key, encrypted, negInput)
}

func Mul(key *paillier.PublicKey, encrypted, input *number.Number) *number.Number {
	var result = new(number.Number)
	result.Base = key.Mul(encrypted.Base, input.Base)
	result.Exp = new(big.Int).Add(encrypted.Exp, input.Exp)
	return result
}

func Div(key *paillier.PublicKey, encrypted, input *number.Number) *number.Number {
	var inverse = 1 / input.Float()
	var invInput = new(number.Number).SetFloat(inverse)

	return Mul(key, encrypted, invInput)
}
