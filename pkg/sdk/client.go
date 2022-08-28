// Package SDK
package sdk

import (
	"errors"

	"github.com/lucasmenendez/gopaillier/internal/number"
	"github.com/lucasmenendez/gopaillier/pkg/paillier"
)

// Struct Client
type Client struct {
	Key *paillier.PrivateKey
}

// Function InitClient
func InitClient(keySize int) (*Client, error) {
	var err error
	var client = &Client{}

	client.Key, err = paillier.NewKeys(keySize)
	return client, err
}

// Function Encrypt
func (client *Client) Encrypt(num *number.Number) (*number.Number, error) {
	if num.IsEncrypted() {
		return nil, errors.New("provided number is already encrypted")
	}

	var err error
	var result = new(number.Number).Set(num)
	result.Value, err = client.Key.PubKey.Encrypt(num.Value)
	result.SetEncrypted(true)
	return result, err
}

// Function Decrypt
func (client *Client) Decrypt(num *number.Number) (*number.Number, error) {
	if !num.IsEncrypted() {
		return nil, errors.New("provided number is not encrypted")
	}

	var err error
	var result = new(number.Number).Set(num)
	result.Value, err = client.Key.Decrypt(num.Value)
	result.SetEncrypted(false)
	return result, err
}
