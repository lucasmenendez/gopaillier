package sdk

import (
	"fmt"
	"testing"

	"github.com/lucasmenendez/gopaillier/internal/number"
)

var a, b float64 = 1890.05213, -0.00125
var c, d int64 = 2340, 12200023

var client, _ = InitClient(128)

var encodedA = new(number.Number).SetFloat(a)
var encryptedA, _ = client.Encrypt(encodedA)
var encodedB = new(number.Number).SetFloat(b)
var encryptedB, _ = client.Encrypt(encodedB)
var encodedC = new(number.Number).SetInt(c)
var encryptedC, _ = client.Encrypt(encodedC)
var encodedD = new(number.Number).SetInt(d)

func TestAdd(t *testing.T) {
	var encryptedSumAB = Add(client.Key.PubKey, encryptedA, encodedB)
	var decryptedSumAB, _ = client.Decrypt(encryptedSumAB)
	var rawSumAB = fmt.Sprintf("%f", a+b)
	if sResult := fmt.Sprintf("%f", decryptedSumAB.Float()); rawSumAB != sResult {
		t.Fatalf("expected %s, got %s", rawSumAB, sResult)
	}

	var encryptedSumCD = Add(client.Key.PubKey, encryptedC, encodedD)
	var decryptedSumCD, _ = client.Decrypt(encryptedSumCD)
	var rawSumCD = fmt.Sprintf("%d", c+d)
	if sResult := fmt.Sprintf("%d", decryptedSumCD.Int()); rawSumCD != sResult {
		t.Fatalf("expected %s, got %s", rawSumCD, sResult)
	}

	var encryptedSumAC = Add(client.Key.PubKey, encryptedA, encodedC)
	var decryptedSumAC, _ = client.Decrypt(encryptedSumAC)
	var rawSumAC = fmt.Sprintf("%f", a+float64(c))
	if sResult := fmt.Sprintf("%f", decryptedSumAC.Float()); rawSumAC != sResult {
		t.Fatalf("expected %s, got %s", rawSumAC, sResult)
	}

	var encryptedSumBD = Add(client.Key.PubKey, encryptedB, encodedD)
	var decryptedSumBD, _ = client.Decrypt(encryptedSumBD)
	var rawSumBD = fmt.Sprintf("%f", b+float64(d))
	if sResult := fmt.Sprintf("%f", decryptedSumBD.Float()); rawSumBD != sResult {
		t.Fatalf("expected %s, got %s", rawSumBD, sResult)
	}
}

func TestSub(t *testing.T) {
	var encryptedDiffAB = Sub(client.Key.PubKey, encryptedA, encodedB)
	var decryptedDiffAB, _ = client.Decrypt(encryptedDiffAB)
	var rawDiffAB = fmt.Sprintf("%f", a-b)
	if sResult := fmt.Sprintf("%f", decryptedDiffAB.Float()); rawDiffAB != sResult {
		t.Fatalf("expected %s, got %s", rawDiffAB, sResult)
	}

	var encryptedDiffCD = Sub(client.Key.PubKey, encryptedC, encodedD)
	var decryptedDiffCD, _ = client.Decrypt(encryptedDiffCD)
	var rawDiffCD = fmt.Sprintf("%d", c-d)
	if sResult := fmt.Sprintf("%d", decryptedDiffCD.Int()); rawDiffCD != sResult {
		t.Fatalf("expected %s, got %s", rawDiffCD, sResult)
	}

	var encryptedDiffAC = Sub(client.Key.PubKey, encryptedA, encodedC)
	var decryptedDiffAC, _ = client.Decrypt(encryptedDiffAC)
	var rawDiffAC = fmt.Sprintf("%f", a-float64(c))
	if sResult := fmt.Sprintf("%f", decryptedDiffAC.Float()); rawDiffAC != sResult {
		t.Fatalf("expected %s, got %s", rawDiffAC, sResult)
	}

	var encryptedDiffBD = Sub(client.Key.PubKey, encryptedB, encodedD)
	var decryptedDiffBD, _ = client.Decrypt(encryptedDiffBD)
	var rawDiffBD = fmt.Sprintf("%f", b-float64(d))
	if sResult := fmt.Sprintf("%f", decryptedDiffBD.Float()); rawDiffBD != sResult {
		t.Fatalf("expected %s, got %s", rawDiffBD, sResult)
	}
}

func TestMul(t *testing.T) {
	var encryptedMulAB = Mul(client.Key.PubKey, encryptedA, encodedB)
	var decryptedMulAB, _ = client.Decrypt(encryptedMulAB)
	var rawMullAB = fmt.Sprintf("%f", a*b)
	if sResult := fmt.Sprintf("%f", decryptedMulAB.Float()); rawMullAB != sResult {
		t.Fatalf("expected %s, got %s", rawMullAB, sResult)
	}

	var encryptedMulCD = Mul(client.Key.PubKey, encryptedC, encodedD)
	var decryptedMulCD, _ = client.Decrypt(encryptedMulCD)
	var rawMullCD = fmt.Sprintf("%d", c*d)
	if sResult := fmt.Sprintf("%d", decryptedMulCD.Int()); rawMullCD != sResult {
		t.Fatalf("expected %s, got %s", rawMullCD, sResult)
	}

	var encryptedMulAC = Mul(client.Key.PubKey, encryptedA, encodedC)
	var decryptedMulAC, _ = client.Decrypt(encryptedMulAC)
	var rawMullAC = fmt.Sprintf("%f", a*float64(c))
	if sResult := fmt.Sprintf("%f", decryptedMulAC.Float()); rawMullAC != sResult {
		t.Fatalf("expected %s, got %s", rawMullAC, sResult)
	}

	var encryptedMulBD = Mul(client.Key.PubKey, encryptedB, encodedD)
	var decryptedMulBD, _ = client.Decrypt(encryptedMulBD)
	var rawMullBD = fmt.Sprintf("%f", b*float64(d))
	if sResult := fmt.Sprintf("%f", decryptedMulBD.Float()); rawMullBD != sResult {
		t.Fatalf("expected %s, got %s", rawMullBD, sResult)
	}
}

func TestDiv(t *testing.T) {
	var encryptedDivAB = Div(client.Key.PubKey, encryptedA, encodedB)
	var decryptedDivAB, _ = client.Decrypt(encryptedDivAB)
	var rawDivlAB = fmt.Sprintf("%f", a/b)
	if sResult := fmt.Sprintf("%f", decryptedDivAB.Float()); rawDivlAB != sResult {
		t.Fatalf("expected %s, got %s", rawDivlAB, sResult)
	}

	var encryptedDivCD = Div(client.Key.PubKey, encryptedC, encodedD)
	var decryptedDivCD, _ = client.Decrypt(encryptedDivCD)
	var rawDivlCD = fmt.Sprintf("%d", c/d)
	if sResult := fmt.Sprintf("%d", decryptedDivCD.Int()); rawDivlCD != sResult {
		t.Fatalf("expected %s, got %s", rawDivlCD, sResult)
	}

	var encryptedDivAC = Div(client.Key.PubKey, encryptedA, encodedC)
	var decryptedDivAC, _ = client.Decrypt(encryptedDivAC)
	var rawDivlAC = fmt.Sprintf("%f", a/float64(c))
	if sResult := fmt.Sprintf("%f", decryptedDivAC.Float()); rawDivlAC != sResult {
		t.Fatalf("expected %s, got %s", rawDivlAC, sResult)
	}

	var encryptedDivBD = Div(client.Key.PubKey, encryptedB, encodedD)
	var decryptedDivBD, _ = client.Decrypt(encryptedDivBD)
	var rawDivlBD = fmt.Sprintf("%f", b/float64(d))
	if sResult := fmt.Sprintf("%f", decryptedDivBD.Float()); rawDivlBD != sResult {
		t.Fatalf("expected %s, got %s", rawDivlBD, sResult)
	}
}
