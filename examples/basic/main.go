package main

import (
	"fmt"
	"log"
	"math/big"
	"math/rand"

	"github.com/lucasmenendez/gopaillier/pkg/paillier"
)

func main() {
	// Creating a new private key with 64 bits of length. It contains its public
	// key with the same size.
	key, err := paillier.NewKeys(64)
	if err != nil {
		log.Fatalln(err)
	}

	// Setting random A & B inputs
	var inputA = new(big.Int).SetInt64(rand.Int63n(100))
	var inputB = new(big.Int).SetInt64(rand.Int63n(100))

	// Encrypting inputs A' = E(A) & B' = E(B)
	var encryptedA, _ = key.PubKey.Encrypt(inputA)
	var encryptedB, _ = key.PubKey.Encrypt(inputB)

	// Calculating some operations:
	//     - A' + B
	//     - A' + B'
	//     - A' + B + B
	//     - A' * B
	var sum = key.PubKey.Add(encryptedA, inputB)
	var sum2 = key.PubKey.Add(sum, inputB)
	var sumEnc = key.PubKey.AddEncrypted(encryptedA, encryptedB)
	var mul = key.PubKey.Mul(encryptedA, inputB)

	// Decrypting results
	var decryptedSum, _ = key.Decrypt(sum)
	var decryptedSum2, _ = key.Decrypt(sum2)
	var decryptedSumEnc, _ = key.Decrypt(sumEnc)
	var decryptedMul, _ = key.Decrypt(mul)

	// Printing results
	fmt.Printf("Original A: %d\n", inputA)
	fmt.Printf("Constant B: %d\n", inputB)
	fmt.Printf("Encrypted A (A'): %d\n", encryptedA)
	fmt.Printf("Encrypted B (B'): %d\n\n", encryptedB)
	fmt.Printf("Encrypted Sum (A' + B): %d\n", sum)
	fmt.Printf("Decrypted Sum: %d + %d = %d\n\n", inputA, inputB, decryptedSum)
	fmt.Printf("Encrypted SumEnc (A' + B'): %d\n", sumEnc)
	fmt.Printf("Decrypted SumEnc: %d + %d = %d\n\n", inputA, inputB, decryptedSumEnc)
	fmt.Printf("Encrypted Sum (A' + B + B): %d\n", sum2)
	fmt.Printf("Decrypted Sum2: %d + %d + %d = %d\n\n", inputA, inputB, inputB, decryptedSum2)
	fmt.Printf("Encrypted Mul (A' * B): %d\n", mul)
	fmt.Printf("Decrypted Mul: %d * %d = %d\n\n", inputA, inputB, decryptedMul)
}
