package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/lucasmenendez/gopaillier/pkg/paillier"
)

func main() {
	var a, b = big.NewInt(2), big.NewInt(5)

	// Creating a new private key with 64 bits of length. It contains its public
	// key with the same size.
	key, err := paillier.NewKeys(128)
	if err != nil {
		log.Fatalln(err)
	}

	// Encrypting inputs A' = E(A) & B' = E(B)
	var encryptedA, _ = key.PubKey.Encrypt(a)
	var encryptedB, _ = key.PubKey.Encrypt(b)

	// Calculating some operations:
	//     - A' + B
	//     - A' + B'
	//     - A' - B
	//     - A' * B
	var sum = key.PubKey.Add(encryptedA, b)
	var sumEnc = key.PubKey.AddEncrypted(encryptedA, encryptedB)
	var sub = key.PubKey.Sub(encryptedA, b)
	var mul = key.PubKey.Mul(encryptedA, b)

	// Decrypting results

	// Sub
	var decryptedSum, _ = key.Decrypt(sum)
	var decryptedSumEnc, _ = key.Decrypt(sumEnc)
	var decryptedSub, _ = key.Decrypt(sub)
	var decryptedMul, _ = key.Decrypt(mul)

	// Printing results
	fmt.Printf("Original A: %d\n", a)
	fmt.Printf("Constant B: %d\n", b)
	fmt.Printf("Encrypted A (A'): %d\n", encryptedA)
	fmt.Printf("Encrypted B (B'): %d\n\n", encryptedB)
	fmt.Printf("Encrypted Sum (A' + B): %d\n", sum)
	fmt.Printf("Decrypted Sum: %d + %d = %d\n\n", a, b, decryptedSum)
	fmt.Printf("Encrypted SumEnc (A' + B'): %d\n", sumEnc)
	fmt.Printf("Decrypted SumEnc: %d + %d = %d\n\n", a, b, decryptedSumEnc)
	fmt.Printf("Encrypted Sub (A' - B): %d\n", sub)
	fmt.Printf("Decrypted Sub: %d - %d = %d\n\n", a, b, decryptedSub)
	fmt.Printf("Encrypted Mul (A' * B): %d\n", mul)
	fmt.Printf("Decrypted Mul: %d * %d = %d\n\n", a, b, decryptedMul)
}
