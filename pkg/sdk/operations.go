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

// Function Add computes the addition of the encrypted number.Number and plain
// number.Number inputs using the provided paillier.PublicKey. It transform the
// number with the greatest Number.Exp and scale its num.Value to normalize it
// with the input number.Number, and then perform de addition. If the greatest
// exponent is not from encrypted number.Number it scale using Paillier
// multiplication. It returns an error if the encrypted number.Number is not
// encrypted or if the input number.Number is encrypted.
func Add(key *paillier.PublicKey, encrypted, input *number.Number) (*number.Number, error) {
	if err := checkArgs(encrypted, input); err != nil {
		return nil, err
	}

	// Instance the result to store the computed Number.Exp and Number.Value.
	var result = new(number.Number)

	// Compare encrypted.Exp and input.Exp, if both are equals, perform Paillier
	// addition using the provided paillier.PublicKey. If not, transform one of
	// the inputs to ensure that both have the same Number.Exp. If the
	// transformation will be applied over encrypted input it will use Paillier
	// operations.
	if cmp := encrypted.Exp.Cmp(input.Exp); cmp == 0 {
		result.Exp = encrypted.Exp
		result.Value = key.Add(encrypted.Value, input.Value)
	} else {
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
	}

	return new(number.Number).SetEncrypted(result), nil
}

// Function Sub computes the subtraction of the encrypted number.Number and the
// input number.Number provided. To perform the operation, it computes the
// negative version of the provided input first and then calculates the addition
// of between it and the encrypted number.Number. It returns an error if the
// encrypted number.Number is not encrypted or if the input number.Number is
// encrypted.
func Sub(key *paillier.PublicKey, encrypted, input *number.Number) (*number.Number, error) {
	if err := checkArgs(encrypted, input); err != nil {
		return nil, err
	}

	var negInput = new(number.Number)
	negInput.Exp = input.Exp
	negInput.Value = new(big.Int).Neg(input.Value)
	return Add(key, encrypted, negInput)
}

// Function Mul computes the multiplication between the encrypted number.Number
// and the input number.Number provided. To perform the operation calculates the
// Paillier multiplication between encrypted.Value and input.Value, and then
// calculates the plain addition between encrypted.Exp and input.Exp. It returns
// an error if the encrypted number.Number is not encrypted or if the input
// number.Number is encrypted.
func Mul(key *paillier.PublicKey, encrypted, input *number.Number) (*number.Number, error) {
	if err := checkArgs(encrypted, input); err != nil {
		return nil, err
	}

	var result = new(number.Number)
	result.Value = key.Mul(encrypted.Value, input.Value)
	result.Exp = new(big.Int).Add(encrypted.Exp, input.Exp)
	return new(number.Number).SetEncrypted(result), nil
}

// Function Div computes the division of the encrypted number.Number and the
// input number.Number provided. To perform the operation, it computes the
// inverse of the provided input first and then calculates the multiplication
// of between it and the encrypted number.Number. It returns an error if the
// encrypted number.Number is not encrypted or if the input number.Number is
// encrypted.
func Div(key *paillier.PublicKey, encrypted, input *number.Number) (*number.Number, error) {
	if err := checkArgs(encrypted, input); err != nil {
		return nil, err
	}

	var inverse = 1 / input.Float()
	var invInput = new(number.Number).SetFloat(inverse)
	return Mul(key, encrypted, invInput)
}
