package sdk

import (
	"github.com/lucasmenendez/gopaillier/internal/number"
	"github.com/lucasmenendez/gopaillier/pkg/paillier"
)

type Client struct {
	Key *paillier.PrivateKey
}

func InitClient() (*Client, error) {
	var err error
	var client = new(Client)

	client.Key, err = paillier.NewKeys(128)
	return client, err
}

func (client *Client) Encrypt(num *number.Number) *number.Number {
	num.Base, _ = client.Key.PubKey.Encrypt(num.Base)
	return num
}

func (client *Client) Decrypt(num *number.Number) *number.Number {
	num.Base, _ = client.Key.Decrypt(num.Base)
	return num
}
