package number

import (
	"math/big"
)

type Number struct {
	Base *big.Int
	Exp  *big.Int
}

func (num *Number) SetInt(input int64) *Number {
	var base, exp int64 = input, 0
	for input%10 == 0 {
		input /= 10
		base = input
		exp++
	}

	num.Base = big.NewInt(base)
	num.Exp = big.NewInt(exp)
	return num
}

func (num *Number) SetFloat(input float64) *Number {
	var bZero, bTen = big.NewFloat(0), big.NewFloat(10)

	var bInput = big.NewFloat(input)
	var base, exp int64 = int64(input), 0

	var diff = new(big.Float).Sub(bInput, big.NewFloat(float64(base)))
	for diff.Cmp(bZero) != 0 {
		bInput = new(big.Float).Mul(bInput, bTen)
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
	output = num.Base.Int64()
	var exp = num.Exp.Int64()
	for exp != 0 {
		if exp > 0 {
			output *= 10
			exp--
		} else {
			output /= 10
			exp++
		}
	}

	return
}

func (num *Number) Float() (output float64) {
	output = float64(num.Base.Int64())
	var exp = num.Exp.Int64()
	for exp != 0 {
		if exp > 0 {
			output *= 10
			exp--
		} else {
			output /= 10
			exp++
		}
	}

	return
}
