package sdk

import (
	"errors"
	"math/big"

	"github.com/lucasmenendez/gopaillier/internal/number"
	"github.com/lucasmenendez/gopaillier/pkg/paillier"
)

func checkArgs(encrypted, plain *number.Number) error {
	if !encrypted.IsEncrypted() {
		return errors.New("first Number provided must be encrypted")
	} else if plain.IsEncrypted() {
		return errors.New("second Number provided must not be encrypted")
	}

	return nil
}

func Add(key *paillier.PublicKey, encrypted, input *number.Number) (*number.Number, error) {
	if err := checkArgs(encrypted, input); err != nil {
		return nil, err
	}

	var result = new(number.Number)
	result.SetEncrypted(true)

	var cmp = encrypted.Exp.Cmp(input.Exp)
	if cmp == 0 {
		result.Exp = encrypted.Exp
		result.Value = key.Add(encrypted.Value, input.Value)
		return result, nil
	}

	var expDiff = new(big.Int).Abs(new(big.Int).Sub(encrypted.Exp, input.Exp))
	var factor = new(big.Int).Exp(big.NewInt(10), expDiff, nil)
	if cmp > 0 {
		result.Exp = input.Exp
		var normalized = key.Mul(encrypted.Value, factor)
		result.Value = key.Add(normalized, input.Value)
	} else {
		result.Exp = encrypted.Exp
		var normalized = new(big.Int).Mul(input.Value, factor)
		result.Value = key.Add(encrypted.Value, normalized)
	}

	return result, nil
}

func Sub(key *paillier.PublicKey, encrypted, input *number.Number) (*number.Number, error) {
	if err := checkArgs(encrypted, input); err != nil {
		return nil, err
	}

	var negInput = new(number.Number)
	negInput.Exp = input.Exp
	negInput.Value = new(big.Int).Neg(input.Value)
	return Add(key, encrypted, negInput)
}

func Mul(key *paillier.PublicKey, encrypted, input *number.Number) (*number.Number, error) {
	if err := checkArgs(encrypted, input); err != nil {
		return nil, err
	}

	var result = new(number.Number)
	result.Value = key.Mul(encrypted.Value, input.Value)
	result.Exp = new(big.Int).Add(encrypted.Exp, input.Exp)
	result.SetEncrypted(true)
	return result, nil
}

func Div(key *paillier.PublicKey, encrypted, input *number.Number) (*number.Number, error) {
	if err := checkArgs(encrypted, input); err != nil {
		return nil, err
	}

	var inverse = 1 / input.Float()
	var invInput = new(number.Number).SetFloat(inverse)
	return Mul(key, encrypted, invInput)
}
