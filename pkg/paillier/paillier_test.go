package paillier

import (
	"math/big"
	"testing"
)

func TestNewKeys(t *testing.T) {
	if _, err := NewKeys(64); err != nil {
		t.Fatalf("expected nil, got %s", err)
	} else if _, err = NewKeys(8); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestEncryptDecrypt(t *testing.T) {
	var key, _ = NewKeys(64)
	var inputA = new(big.Int).SetInt64(12)

	var err error
	var encryptedA *big.Int
	if encryptedA, err = key.PubKey.Encrypt(inputA); err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	var decryptedA *big.Int
	if decryptedA, err = key.Decrypt(encryptedA); err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	if inputA.Cmp(decryptedA) != 0 {
		t.Fatalf("expected %d, got %d", inputA, decryptedA)
	}

	key, err = NewKeys(1024)
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	inputA = new(big.Int).SetInt64(324234987)
	if encryptedA, err = key.PubKey.Encrypt(inputA); err != nil {
		t.Fatalf("expected nil, got %s", err)
	}
	if decryptedA, err = key.Decrypt(encryptedA); err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	if inputA.Cmp(decryptedA) != 0 {
		t.Fatalf("expected %d, got %d", inputA, decryptedA)
	}

	inputA = key.PubKey.N
	if _, err = key.PubKey.Encrypt(inputA); err == nil {
		t.Fatal("expected error, got nil")
	}

	inputA = new(big.Int).Sub(key.PubKey.N, bOne)
	if _, err = key.PubKey.Encrypt(inputA); err != nil {
		t.Fatalf("expected nil, got %s", err)
	}
}

func TestAddEncrypt(t *testing.T) {
	var key, _ = NewKeys(64)

	var inputA = new(big.Int).SetInt64(12)
	var inputB = new(big.Int).SetInt64(3)
	var expectedRes1 = new(big.Int).SetInt64(15)
	var expectedRes2 = new(big.Int).SetInt64(18)

	var encryptedA, _ = key.PubKey.Encrypt(inputA)
	var encryptedB, _ = key.PubKey.Encrypt(inputB)
	var encryptedSum = key.PubKey.AddEncrypted(encryptedA, encryptedB)

	var result, _ = key.Decrypt(encryptedSum)
	if result.Cmp(expectedRes1) != 0 {
		t.Fatalf("expected %d, got %d", expectedRes1, result)
	}

	encryptedSum = key.PubKey.AddEncrypted(encryptedSum, encryptedB)
	result, _ = key.Decrypt(encryptedSum)
	if result.Cmp(expectedRes2) != 0 {
		t.Fatalf("expected %d, got %d", expectedRes2, result)
	}
}

func TestAdd(t *testing.T) {
	var key, _ = NewKeys(64)

	var inputA = new(big.Int).SetInt64(12)
	var inputB = new(big.Int).SetInt64(3)
	var expectedRes1 = new(big.Int).SetInt64(15)
	var expectedRes2 = new(big.Int).SetInt64(18)

	var encryptedA, _ = key.PubKey.Encrypt(inputA)
	var encryptedSum = key.PubKey.Add(encryptedA, inputB)

	var result, _ = key.Decrypt(encryptedSum)
	if result.Cmp(expectedRes1) != 0 {
		t.Fatalf("expected %d, got %d", expectedRes1, result)
	}

	encryptedSum = key.PubKey.Add(encryptedSum, inputB)
	result, _ = key.Decrypt(encryptedSum)
	if result.Cmp(expectedRes2) != 0 {
		t.Fatalf("expected %d, got %d", expectedRes2, result)
	}
}

func TestMul(t *testing.T) {
	var key, _ = NewKeys(64)

	var inputA = new(big.Int).SetInt64(12)
	var inputB = new(big.Int).SetInt64(3)
	var expectedRes1 = new(big.Int).SetInt64(36)
	var expectedRes2 = new(big.Int).SetInt64(108)

	var encryptedA, _ = key.PubKey.Encrypt(inputA)
	var encryptedSum = key.PubKey.Mul(encryptedA, inputB)

	var result, _ = key.Decrypt(encryptedSum)
	if result.Cmp(expectedRes1) != 0 {
		t.Fatalf("expected %d, got %d", expectedRes1, result)
	}

	encryptedSum = key.PubKey.Mul(encryptedSum, inputB)
	result, _ = key.Decrypt(encryptedSum)
	if result.Cmp(expectedRes2) != 0 {
		t.Fatalf("expected %d, got %d", expectedRes2, result)
	}
}
