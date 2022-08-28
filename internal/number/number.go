// Package number
package number

import "math/big"

var iZero = big.NewInt(0)
var iTen = big.NewInt(10)
var fZero = big.NewFloat(0)
var fTen = big.NewFloat(10)

// Struct Number
type Number struct {
	Value     *big.Int
	Exp       *big.Int
	encrypted bool
}

// Function Set
func (num *Number) Set(original *Number) *Number {
	num.Value = original.Value
	num.Exp = original.Exp
	num.encrypted = original.encrypted

	return num
}

// Function SetInt
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

// Function SetFloat
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

// Function SetEncrypted
func (num *Number) SetEncrypted(encrypted bool) *Number {
	num.encrypted = encrypted
	return num
}

// Function IsEncrypted
func (num *Number) IsEncrypted() bool {
	return num.encrypted
}

// Function Int
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

// Function Float
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
