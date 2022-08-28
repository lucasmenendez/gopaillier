package sdk

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/lucasmenendez/gopaillier/internal/number"
)

func TestInitClient(t *testing.T) {
	var keySize = 2
	if _, err := InitClient(keySize); err == nil {
		t.Fatal("expected error, got nil")
	}

	keySize = -16
	if _, err := InitClient(keySize); err == nil {
		t.Fatal("expected error, got nil")
	}

	keySize = 16
	if _, err := InitClient(keySize); err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	keySize = 512
	if _, err := InitClient(keySize); err != nil {
		t.Fatalf("expected nil, got %s", err)
	}
}

func TestEncryptDecrypt(t *testing.T) {
	var a float64 = -1223.1056
	var encodedA = new(number.Number).SetFloat(a)

	var b int64 = 1209345
	var encodedB = new(number.Number).SetInt(b)

	var err error
	var client, _ = InitClient(512)

	var wrongNum, _ = rand.Prime(rand.Reader, 2048)
	var encodedWrong = &number.Number{Value: wrongNum, Exp: big.NewInt(0)}
	if _, err = client.Encrypt(encodedWrong); err == nil {
		t.Fatal("expected error, got nil")
	}

	var fakeClient, _ = InitClient(1024)
	var encryptedWrong, _ = fakeClient.Encrypt(new(number.Number).SetInt(10))
	if _, err = client.Decrypt(encryptedWrong); err == nil {
		t.Fatal("expected error, got nil")
	}

	var encryptedA, decryptedA = new(number.Number), new(number.Number)
	if encryptedA, err = client.Encrypt(encodedA); err != nil {
		t.Fatalf("expected nil, got %s", err)
	}  else if decryptedA, err = client.Decrypt(encryptedA); err != nil {
		t.Fatalf("expected nil, got %s", err)
	} else if a != decryptedA.Float() {
		t.Fatalf("expected nil, got %s", err)
	} else if _, err = client.Encrypt(encryptedA); err == nil {
		t.Fatal("expected error, got nil")
	} else if _, err = client.Decrypt(decryptedA); err == nil {
		t.Fatal("expected error, got nil")
	}

	var encryptedB, decryptedB = new(number.Number), new(number.Number)
	if encryptedB, err = client.Encrypt(encodedB); err != nil {
		t.Fatalf("expected nil, got %s", err)
	} else if decryptedB, err = client.Decrypt(encryptedB); err != nil {
		t.Fatalf("expected nil, got %s", err)
	} else if b != decryptedB.Int() {
		t.Fatalf("expected nil, got %s", err)
	}
}
