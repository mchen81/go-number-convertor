package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	number := parseArgToInt()
	un := uint32(number)

	bi := make(chan string)
	bih := make(chan string)
	hex := make(chan string)

	go func() {
		bi <- toBinary(un)
	}()

	go func() {
		bih <- toHumanReadBinary(un)
	}()

	go func() {
		hex <- toHex(un)
	}()

	fmt.Printf("(signed) %d\n", number)
	fmt.Printf("(unsigned) %d\n", un)
	fmt.Printf("(binary) 0b%s\n", <-bi)
	fmt.Printf("(binary) %s\n", <-bih)
	fmt.Printf("(hex) 0x%s\n", <-hex)

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
		inputNumber = inputNumber[2:]
		number, err = strconv.ParseInt(inputNumber, 16, 64)
		if err != nil {
			log.Fatalf("%s is not a valid hex number", inputNumber)
		}
	} else {
		// other
		number, err = strconv.ParseInt(inputNumber, 10, 64)
		if err != nil {
			log.Fatalf("%s is not a valid digital number", inputNumber)
		}
		if number <= math.MaxInt32 && number >= math.MinInt32 {
			// int32 value
			return int32(number)
		} else if number > math.MaxInt32 && number <= math.MaxUint32 {
			// uint32 value
			var ui uint64
			ui, _ = strconv.ParseUint(inputNumber, 10, 64)
			number = int64(ui)
		} else if number > math.MaxUint32 || number < math.MinInt32 {
			log.Fatalf("%s is too big or too small(for int32).", inputNumber)
		} else {
			log.Fatalf("Cannot regconize %s.", inputNumber)
		}
	}
	return int32(number)
}
