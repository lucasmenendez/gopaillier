# Paillier cryprosystem (a Go implementation)

Package paillier is the Go implementation of the Paillier Cryptosysmem described by Pascal Paillier in 1999. It is a probabilistic asymmetric algorithm for public key cryptography, that support homomorphic addition and multiplication of plaintexts to a ciphertext.

The current package extends the original Paillier Cryptosystem to support same homomorphic operation over non-positive integers with an optimized decryption procedure (read more [here](https://eprint.iacr.org/2010/520)).

Read more about Paillier Cryptosystems on its [WikiPedia page](https://en.wikipedia.org/wiki/Paillier_cryptosystem).

## Installation and use

### Installation
Get package via `go` CLI:

```sh
go get github.com/lucasmenendez/gopaillier/pkg/paillier
```

### Encrypt and decrypt inputs

```go
package main

import (
    "log"
    "math/big"

    "github.com/lucasmenendez/gopaillier/pkg/paillier"
)

func main() {
    var key, _ = paillier.NewKeys(128)

    // Create inputs
    var A, B = big.NewInt(10), big.NewInt(5)
    log.Printf("[Original] A: %d, B: %d\n", A, B)

    // Calc C(a) & C(b)
    var encryptedA, _ = key.PubKey.Encrypt(A)
    var encryptedB, _ = key.PubKey.Encrypt(B)
    
    // Decrypte C(a) & C(b)
    var decryptedA, _ = key.Decrypt(encryptedA)
    var decryptedB, _ = key.Decrypt(encryptedB)
    log.Printf("[Decrypted] A: %d, B: %d\n", decryptedA, decryptedB)
}
```

### Perform and addition and multiplication over encrypted data

```go
package main

import (
    "log"
    "math/big"

    "github.com/lucasmenendez/gopaillier/pkg/paillier"
)

func main() {
    var key, _ = paillier.NewKeys(128)

    // Create inputs
    var A, B = big.NewInt(10), big.NewInt(5)
    log.Printf("A: %d, B: %d\n", A, B)

    // Calc C(a) & C(b)
    var encryptedA, _ = key.PubKey.Encrypt(A)
    var encryptedB, _ = key.PubKey.Encrypt(B)
    
    // Add operation 1: C(a) + C(b)
    var encryptedAdd1 = key.PubKey.AddEncrypted(encryptedA, encryptedB)
    var decryptedAdd1, _ = key.Decrypt(encryptedAdd1)
    log.Printf("C(a) + C(b): %d", decryptedAdd1)
    
    // Add operation 2: C(a) + b
    var encryptedAdd2 = key.PubKey.Add(encryptedA, B)
    var decryptedAdd2, _ = key.Decrypt(encryptedAdd2)
    log.Printf("C(a) + b: %d", decryptedAdd2)

    // Mul operation 2: C(a) * b
    var encryptedMul = key.PubKey.Mul(encryptedA, B)
    var decryptedMul, _ = key.Decrypt(encryptedMul)
    log.Printf("C(a) * b: %d", decryptedMul)
}
```