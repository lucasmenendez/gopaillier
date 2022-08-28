// Package number provide a common representation of integers and floating point
// numbers uting big.Int's. It transform any number into its integer based
// standard form representation, storing its integer value and the correct
// exponent to keep the original value:
//
//	X = value * 10^exp --> 1.032 = 1032 * 10^-3
//
// This package allows to use any integer based cryptosystem over floating point
// numbers too.
package number

import "math/big"

var iZero = big.NewInt(0)
var iTen = big.NewInt(10)
var fZero = big.NewFloat(0)
var fTen = big.NewFloat(10)

// Struct Number includes the integers value of the original number with the
// original power of ten exponent, allowing to encrypt and decrypt the value and
// operate over it.
type Number struct {
	Value     *big.Int
	Exp       *big.Int
	encrypted bool
}

// Function IsEncrypted return if the current number representation is encrypted
// or not.
func (num *Number) IsEncrypted() bool {
	return num.encrypted
}

// Function Set copy the values of the original Number into the current Number
// num and return it as result. By default, the resulting Number will be
// created as decrypted, to create as encrypted use number.SetEncrypted()
// function.
func (num *Number) Set(original *Number) *Number {
	num.Value = original.Value
	num.Exp = original.Exp
	num.encrypted = false

	return num
}

// Function SetEncrypted copy the values of the original Number into the current
// Number num and return it as result. By default, the resulting Number will be
// created as encrypted, to create as decrypted use number.Set() function.
func (num *Number) SetEncrypted(original *Number) *Number {
	num.Set(original)
	num.encrypted = true
	return num
}

// Function SetInt compute and stores into the current Number num the correct
// integer value and exponent of the provided int input and return it as result.
func (num *Number) SetInt(input int64) *Number {
	var bInput = big.NewInt(input)

	var exp int64
	for new(big.Int).Mod(bInput, iTen).Cmp(iZero) == 0 {
		bInput.Div(bInput, iTen)
		exp++
	}

	num.Value = new(big.Int).Set(bInput)
	num.Exp = big.NewInt(exp)
	return num
}

// Function SetFloat compute and stores into the current Number num the correct
// integer value and exponent of the provided float input and return it as
// result.
func (num *Number) SetFloat(input float64) *Number {
	var bInput = big.NewFloat(input)
	var value, exp int64 = int64(input), 0

	var diff = new(big.Float).Sub(bInput, big.NewFloat(float64(value)))
	for diff.Cmp(fZero) != 0 {
		bInput = new(big.Float).Mul(bInput, fTen)
		value, _ = bInput.Int64()
		exp--

		diff = new(big.Float).Sub(bInput, big.NewFloat(float64(value)))
	}

	var numInt = new(Number).SetInt(value)
	num.Value = numInt.Value
	num.Exp = new(big.Int).Add(numInt.Exp, big.NewInt(exp))
	return num
}

// Function Int returns the original int value of the current Number num
// computing the value of num.Value * 10^num.Exp.
func (num *Number) Int() (output int64) {
	var bOutput = new(big.Int).Set(num.Value)
	var exp = num.Exp.Int64()

	for exp != 0 {
		if exp > 0 {
			bOutput.Mul(bOutput, iTen)
			exp--
		} else {
			bOutput.Div(bOutput, iTen)
			exp++
		}
	}

	output = bOutput.Int64()
	return
}

// Function Float returns the original float value of the current Number num
// computing the value of num.Value * 10^num.Exp.
func (num *Number) Float() (output float64) {
	var bOutput = new(big.Float).SetInt(num.Value)
	var exp = num.Exp.Int64()

	for exp != 0 {
		if exp > 0 {
			bOutput.Mul(bOutput, fTen)
			exp--
		} else {
			bOutput.Quo(bOutput, fTen)
			exp++
		}
	}

	output, _ = bOutput.Float64()
	return
}
