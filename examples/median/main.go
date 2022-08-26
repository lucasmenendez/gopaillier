package main

import (
	"fmt"

	"github.com/lucasmenendez/gopaillier/internal/number"
	"github.com/lucasmenendez/gopaillier/pkg/sdk"
)

func main() {
	var aClient, _ = sdk.InitClient()

	// Get first number record and encode it
	var numbers = []int64{4, 27, 2, 39, 25, 37, 85, 17, 15, 21, 58, 27, 77, 4, 91, 64, 90, 78, 48, 43, 40, 55, 56, 57, 92, 50, 78, 6, 42, 64, 19, 14, 7, 61, 87, 86, 73, 82, 72, 48, 28, 76, 49, 65, 34, 81, 40, 10, 83, 70, 30, 55, 35, 85, 45, 6, 41, 24, 42, 61, 34, 54, 88, 14, 99, 23, 9, 69, 36, 18, 59, 49, 48, 14, 13, 11, 42, 80, 91, 50, 35, 26, 90, 60, 41, 26, 85, 84, 9, 79, 30, 81, 51, 90, 16, 21, 13, 69, 57, 71}
	var first = numbers[0]
	var encodedfirst = new(number.Number).SetInt(first)

	// Instance the sumatories with the first value
	var rawSumatory = float64(first)
	var encryptedSumatory, _ = aClient.Encrypt(encodedfirst)

	// Complete the sumatories iterating over the rest of the records
	for _, num := range numbers[1:] {
		rawSumatory += float64(num)

		var encoded = new(number.Number).SetInt(num)
		encryptedSumatory = sdk.Add(aClient.Key.PubKey, encryptedSumatory, encoded)
	}

	// Get decrypted median dividing the decrypted sumatory by the number of items
	var encodedLen = new(number.Number).SetInt(int64(len(numbers)))
	var encryptedMedian = sdk.Div(aClient.Key.PubKey, encryptedSumatory, encodedLen)

	// Decrypt it and decode it
	var decryptedMedian, _ = aClient.Decrypt(encryptedMedian)
	var decodedMedian = decryptedMedian.Float()

	// Calc raw median
	var median = rawSumatory / float64(len(numbers))

	// Make some prints
	fmt.Printf("\nPerform median of %d numbers: \n%v\n", len(numbers), numbers)
	fmt.Printf("\t- Raw sum: %.2f, Encrypted sum: %d\n", rawSumatory, encryptedSumatory)
	fmt.Printf("\t- Raw median: %.2f, Decrypted median: %.2f\n\n", median, decodedMedian)
}
