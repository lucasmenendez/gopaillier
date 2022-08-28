// Package SDK allows to interact with paillier package easily, allowing to use
// floating point and integer numbers with it, and supporting more operations
// such as substraction and division over encrypted numbers.
package sdk

import (
	"errors"

	"github.com/lucasmenendez/gopaillier/internal/number"
	"github.com/lucasmenendez/gopaillier/pkg/paillier"
)

// Struct Client contains a Paillier key pair allowing to encrypt and decrypt
// number.Number instances. Sharing Client.Key.PubKey (a paillier.PubKey
// instance) with an external actor, it could compute operations over a
// number.Number encrypted with the same paillier.PubKey.
type Client struct {
	Key *paillier.PrivateKey
}

// Function InitClient returns a new client with a generated paillier.PrivKey
// and paillier.PubKet pair with the size provided.
func InitClient(keySize int) (*Client, error) {
	var err error
	var client = &Client{}

	client.Key, err = paillier.NewKeys(keySize)
	return client, err
}

// Function Encrypt returns the encrypted version of the provided number.Number.
// It returns an error if the provided input is already encrypted or if some
// error occurs during the input encryption process.
func (client *Client) Encrypt(num *number.Number) (*number.Number, error) {
	if num.IsEncrypted() {
		return nil, errors.New("provided number is already encrypted")
	}

	var err error
	var result = new(number.Number).SetEncrypted(num)
	result.Value, err = client.Key.PubKey.Encrypt(num.Value)
	return result, err
}

// Function Decrypt returns the decrypted version of the provided number.Number.
// It returns an error if the provided input is not encrypted or if some error
// occurs during the input decryption process.
func (client *Client) Decrypt(num *number.Number) (*number.Number, error) {
	if !num.IsEncrypted() {
		return nil, errors.New("provided number is not encrypted")
	}

	var err error
	var result = new(number.Number).Set(num)
	result.Value, err = client.Key.Decrypt(num.Value)
	return result, err
}
