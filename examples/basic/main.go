package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/lucasmenendez/gopaillier/pkg/encoder"
	"github.com/lucasmenendez/gopaillier/pkg/paillier"
)

func main() {
	var a, b float64 = 2.4, 3.12

	// Creating a new private key with 64 bits of length. It contains its public
	// key with the same size.
	key, err := paillier.NewKeys(64)
	if err != nil {
		log.Fatalln(err)
	}

	// Setting random A & B inputs
	var inputA = encoder.EncodeFloat(a)
	var inputB = encoder.EncodeFloat(b)

	// Encrypting inputs A' = E(A) & B' = E(B)
	var encryptedA, _ = key.PubKey.Encrypt(inputA)
	var encryptedB, _ = key.PubKey.Encrypt(inputB)

	// Calculating some operations:
	//     - A' + B
	//     - A' + B'
	//     - A' + B + B
	//     - A' * B
	var sub = key.PubKey.Sub(encryptedA, inputB)
	var sum = key.PubKey.Add(encryptedA, inputB)
	var sumEnc = key.PubKey.AddEncrypted(encryptedA, encryptedB)
	var mul = key.PubKey.Mul(encryptedA, inputB)

	// Decrypting results
	var decryptedSub, _ = key.Decrypt(sub)
	var decryptedSum, _ = key.Decrypt(sum)
	var decryptedSumEnc, _ = key.Decrypt(sumEnc)
	var decryptedMul, _ = key.Decrypt(mul)
	decryptedMul = new(big.Int).Div(decryptedMul, encoder.Factor)

	var decodedSub = encoder.DecodeFloat(decryptedSub)
	var decodedSum = encoder.DecodeFloat(decryptedSum)
	var decodedSumEnc = encoder.DecodeFloat(decryptedSumEnc)
	var decodedMul = encoder.DecodeFloat(decryptedMul)

	// Printing results
	fmt.Printf("Original A: %.2f\n", a)
	fmt.Printf("Constant B: %.2f\n", b)
	fmt.Printf("Encrypted A (A'): %d\n", encryptedA)
	fmt.Printf("Encrypted B (B'): %d\n\n", encryptedB)
	fmt.Printf("Encrypted Sum (A' + B): %d\n", sum)
	fmt.Printf("Decrypted Sum: %.2f + %.2f = %.2f\n\n", a, b, decodedSum)
	fmt.Printf("Encrypted SumEnc (A' + B'): %d\n", sumEnc)
	fmt.Printf("Decrypted SumEnc: %.2f + %.2f = %.2f\n\n", a, b, decodedSumEnc)
	fmt.Printf("Encrypted Sub (A' - B): %d\n", sub)
	fmt.Printf("Decrypted Sub: %.2f - %.2f = %.2f\n\n", a, b, decodedSub)
	fmt.Printf("Encrypted Mul (A' * B): %d\n", mul)
	fmt.Printf("Decrypted Mul: %.2f * %.2f = %.2f\n\n", a, b, decodedMul)

	// Calc median
	var numbers = []int64{4, 27, 2, 39, 25, 37, 85, 17, 15, 21, 58, 27, 77, 4, 91, 64, 90, 78, 48, 43, 40, 55, 56, 57, 92, 50, 78, 6, 42, 64, 19, 14, 7, 61, 87, 86, 73, 82, 72, 48, 28, 76, 49, 65, 34, 81, 40, 10, 83, 70, 30, 55, 35, 85, 45, 6, 41, 24, 42, 61, 34, 54, 88, 14, 99, 23, 9, 69, 36, 18, 59, 49, 48, 14, 13, 11, 42, 80, 91, 50, 35, 26, 90, 60, 41, 26, 85, 84, 9, 79, 30, 81, 51, 90, 16, 21, 13, 69, 57, 71}
	var sumatory = float64(numbers[0])

	var first = encoder.EncodeInt(numbers[0])
	var encryptedSumatory, _ = key.PubKey.Encrypt(first)
	for _, num := range numbers[1:] {
		sumatory += float64(num)

		var encoded = encoder.EncodeInt(num)
		encryptedSumatory = key.PubKey.Add(encryptedSumatory, encoded)
	}

	var factor = encoder.EncodeFloat(1 / float64(len(numbers)))
	var encryptedMedian = key.PubKey.Mul(encryptedSumatory, factor)
	var decryptedMedian, _ = key.Decrypt(encryptedMedian)
	decryptedMedian = new(big.Int).Div(decryptedMedian, encoder.Factor)
	var decodedMedian = encoder.DecodeFloat(decryptedMedian)

	var median = sumatory / float64(len(numbers))
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Raw median: %.2f, Encrypted median: %.2f\n", median, decodedMedian)
}
