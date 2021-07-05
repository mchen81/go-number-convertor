package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// functionalities:

// base2 -> base10
// base10 -> base2

// nt -b (binary number)
// nt (default base10 signed number)
// nt -u (base10 unsigned number)

func main() {

	number := parseArgToInt()
	un := uint32(number)
	bi := toBinary(un)
	bih := toHumanReadBinary(un)
	hex := toHex(un)

	fmt.Printf("(signed) %d\n", number)
	fmt.Printf("(unsigned) %d\n", un)
	fmt.Printf("(binary) 0b%s\n", bi)
	fmt.Printf("(binary) %s\n", bih)
	fmt.Printf("(unsigned) 0x%s\n", hex)

}

func parseArgToInt() int32 {
	args := os.Args[1:]
	argsLength := len(args)

	// arg types
	var number int64
	var err error
	// case1 a number
	if argsLength < 1 {
		log.Fatal("No Args Found")
	}
	inputNumber := args[0]
	if len(inputNumber) < 2 {
		number, err = strconv.ParseInt(inputNumber, 10, 32)
		if err != nil {
			log.Fatalf("%s is not a valid binary number", inputNumber)
		}
		return int32(number)
	}

	prefix := inputNumber[0:2]

	if prefix == "0b" {
		// binary
		inputNumber = inputNumber[2:]
		number, err = strconv.ParseInt(inputNumber, 2, 64)
		if err != nil {
			log.Fatalf("%s is not a valid binary number", inputNumber)
		}
	} else if prefix == "0x" {
		// hex
		number, err = strconv.ParseInt(inputNumber, 16, 64)
		if err != nil {
			log.Fatalf("%s is not a valid hex number", inputNumber)
		}
	} else {
		// base10
		// unsign
		const MaxUint = ^uint(0)
		const MinUint = 0
		const MaxInt = int(MaxUint >> 1)
		const MinInt = -MaxInt - 1

		number, err = strconv.ParseInt(inputNumber, 10, 64)
		if err != nil {
			log.Fatalf("not support %s", inputNumber)
		}
		if number <= math.MaxInt32 && number >= math.MinInt32 {
			return int32(number)
		} else if number > math.MaxInt32 && number <= math.MaxUint32 {
			var ui uint64
			ui, err = strconv.ParseUint(inputNumber, 10, 64)
			if err != nil {
				log.Fatalf("%s is not a valid unsign int", inputNumber)
			}
			number = int64(ui)
		} else {
			log.Fatalf("not support %s", inputNumber)
		}
	}
	return int32(number)
}
