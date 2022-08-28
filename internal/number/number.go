package number

import "math/big"

var iZero = big.NewInt(0)
var iTen = big.NewInt(10)
var fZero = big.NewFloat(0)
var fTen = big.NewFloat(10)

type Number struct {
	Base *big.Int
	Exp  *big.Int
}

func (num *Number) SetInt(input int64) *Number {
	var bInput = big.NewInt(input)

	var exp int64
	for new(big.Int).Mod(bInput, iTen).Cmp(iZero) == 0 {
		bInput.Div(bInput, iTen)
		exp++
	}

	num.Base = new(big.Int).Set(bInput)
	num.Exp = big.NewInt(exp)
	return num
}

func (num *Number) SetFloat(input float64) *Number {
	var bInput = big.NewFloat(input)
	var base, exp int64 = int64(input), 0

	var diff = new(big.Float).Sub(bInput, big.NewFloat(float64(base)))
	for diff.Cmp(fZero) != 0 {
		bInput = new(big.Float).Mul(bInput, fTen)
		base, _ = bInput.Int64()
		exp--

		diff = new(big.Float).Sub(bInput, big.NewFloat(float64(base)))
	}

	var numInt = new(Number).SetInt(base)
	num.Base = numInt.Base
	num.Exp = new(big.Int).Add(numInt.Exp, big.NewInt(exp))
	return num
}

func (num *Number) Int() (output int64) {
	var bOutput = new(big.Int).Set(num.Base)
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

func (num *Number) Float() (output float64) {
	var bOutput = new(big.Float).SetInt(num.Base)
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
