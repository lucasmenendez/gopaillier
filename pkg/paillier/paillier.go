package paillier

import (
	"crypto/rand"
	"errors"
	"math/big"
)

var bOne *big.Int = new(big.Int).SetInt64(1)

type PublicKey struct {
	N, Nsq, G *big.Int
	Len       int64
}

type PrivateKey struct {
	d, u   *big.Int
	Len    int64
	PubKey *PublicKey
}

func NewKeys(size int) (*PrivateKey, error) {
	var err error
	if size < 16 {
		return nil, errors.New("size must be greater than 16")
	}

	var p, q *big.Int
	if p, err = rand.Prime(rand.Reader, size); err != nil {
		return nil, err
	} else if q, err = rand.Prime(rand.Reader, size); err != nil {
		return nil, err
	}

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

func (key *PublicKey) Encrypt(input *big.Int) (*big.Int, error) {
	if input.Cmp(key.N) != -1 {
		return nil, errors.New("input too long on encrypt")
	}

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

	var (
		gm     = new(big.Int).Exp(key.G, input, key.Nsq)
		rn     = new(big.Int).Exp(r, key.N, key.Nsq)
		gmrn   = new(big.Int).Mul(gm, rn)
		output = new(big.Int).Mod(gmrn, key.Nsq)
	)

	return output, nil
}

func (key *PrivateKey) Decrypt(input *big.Int) (*big.Int, error) {
	if input.Cmp(key.PubKey.Nsq) != -1 {
		return nil, errors.New("input too long on decrypt")
	}

	var (
		cd = new(big.Int).Exp(input, key.d, key.PubKey.Nsq)
		l  = new(big.Int).Div(new(big.Int).Sub(cd, bOne), key.PubKey.N)
		d  = new(big.Int).Mod(new(big.Int).Mul(l, key.u), key.PubKey.N)
	)

	// Parse sign appliying: D'(c) = [D(c)]_n.
	// Where [x]_n = ((x + ⌊n/2⌋) mod n) - ⌊n/2⌋
	// https://tinyurl.com/paillier-subtraction-negatives
	var (
		n2 = new(big.Int).Div(key.PubKey.N, big.NewInt(2))
		xn = new(big.Int).Mod(new(big.Int).Add(d, n2), key.PubKey.N)
	)

	return new(big.Int).Sub(xn, n2), nil
}

func (key *PublicKey) AddEncrypted(a, b *big.Int) *big.Int {
	return new(big.Int).Mod(new(big.Int).Mul(a, b), key.Nsq)
}

func (key *PublicKey) Add(a, b *big.Int) *big.Int {
	// x * y mod n^2
	var gb = new(big.Int).Exp(key.G, b, key.Nsq)
	return new(big.Int).Mod(new(big.Int).Mul(a, gb), key.Nsq)
}

func (key *PublicKey) Mul(a, b *big.Int) *big.Int {
	return new(big.Int).Exp(a, b, key.Nsq)
}
