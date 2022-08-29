package main

import (
	"fmt"

	"github.com/lucasmenendez/gopaillier/pkg/number"
	"github.com/lucasmenendez/gopaillier/pkg/sdk"
)

func main() {
	// Set A plain data and convert to Number
	var aData float64 = 12.2
	var aNum *number.Number = new(number.Number).SetFloat(aData)

	// Instance client and encrypt data and send encrypted data to a third
	// party with the public key.
	var aClient, _ = sdk.InitClient(512)
	var aEncrypted, _ = aClient.Encrypt(aNum)

	// Set B plain data and convert to Number
	var bData float64 = -0.00005
	var bNum *number.Number = new(number.Number).SetFloat(bData)

	// Perform the multiplication between the encrypted received Number and the
	// B Number using the received public key.
	var sumEncrypted, _ = sdk.Add(aClient.Key.PubKey, aEncrypted, bNum)
	var subEncrypted, _ = sdk.Sub(aClient.Key.PubKey, aEncrypted, bNum)
	var mulEncrypted, _ = sdk.Mul(aClient.Key.PubKey, aEncrypted, bNum)
	var divEncrypted, _ = sdk.Div(aClient.Key.PubKey, aEncrypted, bNum)

	// Send the encrypted Mul to A to decrypt the value and print the plain
	// Mul.
	var sumDecrypted, _ = aClient.Decrypt(sumEncrypted)
	var subDecrypted, _ = aClient.Decrypt(subEncrypted)
	var mulDecrypted, _ = aClient.Decrypt(mulEncrypted)
	var divDecrypted, _ = aClient.Decrypt(divEncrypted)

	var aSum = sumDecrypted.Float()
	var aSub = subDecrypted.Float()
	var aMul = mulDecrypted.Float()
	var aDiv = divDecrypted.Float()

	fmt.Printf("%f + %f = %f\n", aData, bData, aSum)
	fmt.Printf("%f - %f = %f\n", aData, bData, aSub)
	fmt.Printf("%f * %f = %f\n", aData, bData, aMul)
	fmt.Printf("%f / %f = %f\n", aData, bData, aDiv)
}
