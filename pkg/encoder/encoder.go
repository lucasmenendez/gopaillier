package encoder

import "math/big"

var Precision int64 = 5
var Factor = new(big.Int).Exp(big.NewInt(10), big.NewInt(Precision), nil)

var (
	intTransformer   = Factor
	floatTransformer = new(big.Float).SetInt(Factor)
)

func EncodeInt(input int64) *big.Int {
	var bigInput = big.NewInt(input)
	return new(big.Int).Mul(bigInput, intTransformer)
}

func DecodeInt(encoded *big.Int) int64 {
	var bigOutput = new(big.Int).Div(encoded, intTransformer)
	return bigOutput.Int64()
}

func EncodeFloat(input float64) *big.Int {
	var bigInput = big.NewFloat(input)
	var floatOutput = new(big.Float).Mul(bigInput, floatTransformer)

	var output = new(big.Int)
	floatOutput.Int(output)

	return output
}

func DecodeFloat(encoded *big.Int) float64 {
	var floatInput = new(big.Float).SetInt(encoded)
	var bigOutput = new(big.Float).Quo(floatInput, floatTransformer)

	var output, _ = bigOutput.Float64()
	return output
}
