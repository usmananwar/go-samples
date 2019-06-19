package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

func main() {

	//calcDiff()

	aString := "1"

	shaResult := sha256.Sum256([]byte(aString))

	z := new(big.Int).SetBytes(shaResult[:])

	fmt.Println(z)

}

func calcDiff() {
	baseString := "Hello, world!"
	fmt.Println("****TARGET****")
	number := getAbigNumber("2")
	var targetNumber = number.Exp(number, getAbigNumber("219"), nil)
	fmt.Println("Number: \n", targetNumber)

	fmt.Println("--------------------------------------------------------------------------------------------")

	fmt.Println("****SOLUTIONS****")
	bignum := getAbigNumber("0")
	for {
		somestring := baseString + bignum.String()
		resultBytes := sha256.Sum256([]byte(somestring))
		result := new(big.Int).SetBytes(resultBytes[:])

		fmt.Println("Hashed string:", somestring)
		fmt.Printf("Hex: %x\n", result)
		fmt.Println("Number: \n", result)

		if result.Cmp(targetNumber) <= 0 {
			fmt.Println("Solution found")
			break
		}

		incrementBigInt(bignum)
		fmt.Println("Number: \n", bignum)
	}

}

func incrementBigInt(num *big.Int) {
	z, success := new(big.Int).SetString("1", 10)
	if !success {
		panic("Error converting ... ")
	}
	num.Add(num, z)
}

func getAbigNumber(source string) *big.Int {
	var targetNumber, ok = new(big.Int).SetString(source, 10)
	if !ok {
		panic("Error converting ... ")
	}
	return targetNumber
}
