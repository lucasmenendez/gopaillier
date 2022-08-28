// Package paillier is the Go implementation of the Paillier Cryptosysmem
// described by Pascal Paillier in 1999. It is a probabilistic asymmetric
// algorithm for public key cryptography, that support homomorphic addition and
// multiplication of plaintexts to a ciphertext.
// The current package extends the original Paillier Cryptosystem to support
// same homomorphic operation over non-positive integers with an optimized
// decryption procedure (read more here: https://eprint.iacr.org/2010/520).
// Read more: https://en.wikipedia.org/wiki/Paillier_cryptosystem
package paillier

import (
	"crypto/rand"
	"errors"
	"math/big"
)

var bOne *big.Int = new(big.Int).SetInt64(1)

// Struct PublicKey includes the required parameters n and g, the
// precomputed n^2 value (nsq) and the length of the key.
type PublicKey struct {
	N, Nsq, G *big.Int
	Len       int64
}

// Struct PrivateKey includes the required parameters λ (d) and μ (u), with the
// associated PublicKey and the length of the key.
type PrivateKey struct {
	d, u   *big.Int
	Len    int64
	PubKey *PublicKey
}

// Function NewKeys computes the required parameters of a paillier.PrivateKey,
// including the parameters of its paillier.PublicKey, following the key
// generation algorithm. Read more:
// https://en.wikipedia.org/wiki/Paillier_cryptosystem#Key_generation
func NewKeys(size int) (*PrivateKey, error) {
	var err error
	if size < 16 {
		return nil, errors.New("size must be greater than 16")
	}

	// Calc p and q large prime numbers with equivalent lenght
	var p, q *big.Int
	if p, err = rand.Prime(rand.Reader, size); err != nil {
		return nil, err
	} else if q, err = rand.Prime(rand.Reader, size); err != nil {
		return nil, err
	}

	// Compute public key parameters n (n), nsq (nsq) and g (g), where:
	//		n = p * q => n
	//		nsq = n^2 => nsq
	//		g = n - 1 => g
	// Also compute private key parameters λ (d) and μ (u), where:
	// 		λ = φ(n) = (p - 1)(q - 1) => d
	//		μ = φ(n)^-1 mod n => u
	var (
		pl  = new(big.Int).Sub(p, bOne)
		ql  = new(big.Int).Sub(q, bOne)
		n   = new(big.Int).Mul(p, q)
		nsq = new(big.Int).Mul(n, n)
		g   = new(big.Int).Add(n, bOne)
		d   = new(big.Int).Mul(pl, ql)
		u   = new(big.Int).ModInverse(d, n)
	)

	return &PrivateKey{
		d, u, int64(size),
		&PublicKey{n, nsq, g, int64(size)},
	}, nil
}

// Function Encrypt convert the received input big.Int into its encrypted
// version using the current paillier.PublicKey. Returns an error if the
// provided input its too big for the current key paillier.PublicKey size or
// if the random number generation fails. Read more:
// https://en.wikipedia.org/wiki/Paillier_cryptosystem#Encryption
func (key *PublicKey) Encrypt(input *big.Int) (*big.Int, error) {
	if input.Cmp(key.N) != -1 {
		return nil, errors.New("input too long on encrypt")
	}

	// Calc a large random number (r) that satisfies the condition of
	// gdc(key.N, r) == 1
	var (
		err  error
		r    *big.Int
		size = new(big.Int).SetInt64(key.Len)
	)
	for {
		if r, err = rand.Int(rand.Reader, size); err != nil {
			return nil, err
		}

		if gdc := new(big.Int).GCD(nil, nil, r, key.N); gdc.Cmp(bOne) == 0 {
			break
		}
	}

	// Compute encrypted message (C) of input (m), where:
	//		C = g^m * r^n mod nsq
	var (
		gm     = new(big.Int).Exp(key.G, input, key.Nsq)
		rn     = new(big.Int).Exp(r, key.N, key.Nsq)
		gmrn   = new(big.Int).Mul(gm, rn)
		output = new(big.Int).Mod(gmrn, key.Nsq)
	)

	return output, nil
}

// Function Decrypt convert the received encrypted input big.Int into its
// decrypted version using the current paillier.PrivateKey. Returns an error if
// the provided input its too big for the current key paillier.PrivateKey size.
// Read more: https://en.wikipedia.org/wiki/Paillier_cryptosystem#Decryption
func (key *PrivateKey) Decrypt(input *big.Int) (*big.Int, error) {
	if input.Cmp(key.PubKey.Nsq) != -1 {
		return nil, errors.New("input too long on decrypt")
	}

	// Compute decrypted message (D) of input (c), where:
	//		L(x) = (x - 1) / n
	//		D = L(c^λ mod nsq) * μ mod n
	var (
		cd = new(big.Int).Exp(input, key.d, key.PubKey.Nsq)
		l  = new(big.Int).Div(new(big.Int).Sub(cd, bOne), key.PubKey.N)
		d  = new(big.Int).Mod(new(big.Int).Mul(l, key.u), key.PubKey.N)
	)

	// Parse sign appliying: D'(c) = [D(c)]_n, where:
	// 		[x]_n = ((x + ⌊n/2⌋) mod n) - ⌊n/2⌋
	// Read more here: https://tinyurl.com/paillier-subtraction-negatives
	var (
		n2 = new(big.Int).Div(key.PubKey.N, big.NewInt(2))
		xn = new(big.Int).Mod(new(big.Int).Add(d, n2), key.PubKey.N)
	)

	return new(big.Int).Sub(xn, n2), nil
}

// Function AddEncrypted returns the result of adding both encrypted big.Int's
// provided as input (a and b). Read more:
// https://en.wikipedia.org/wiki/Paillier_cryptosystem#Homomorphic_properties
func (key *PublicKey) AddEncrypted(a, b *big.Int) *big.Int {
	// Compute a + b, where:
	//		a = E(m1) & b = E(m2)
	//		a + b = a * b mod nsq
	return new(big.Int).Mod(new(big.Int).Mul(a, b), key.Nsq)
}

// Function Add returns the result of adding the the encrypted big.Int
// provided as a input to the plain big.Int provided as b input. Read more:
// https://en.wikipedia.org/wiki/Paillier_cryptosystem#Homomorphic_properties
func (key *PublicKey) Add(a, b *big.Int) *big.Int {
	// Compute a + b, where:
	//		a = E(m1) & b = m2
	//		a + b = a * g^b mod nsq
	var gb = new(big.Int).Exp(key.G, b, key.Nsq)
	return new(big.Int).Mod(new(big.Int).Mul(a, gb), key.Nsq)
}

// Function Add returns the result of to multiplying the the encrypted big.Int
// provided as a input to the plain big.Int provided as b input. Read more:
// https://en.wikipedia.org/wiki/Paillier_cryptosystem#Homomorphic_properties
func (key *PublicKey) Mul(a, b *big.Int) *big.Int {
	// Compute a * b, where:
	//		a = E(m1) & b = m2
	//		a + b = a^b mod n^2
	return new(big.Int).Exp(a, b, key.Nsq)
}
