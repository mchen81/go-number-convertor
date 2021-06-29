package main

import (
	"fmt"
	"log"
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
	number := parseArg()
	fmt.Println(number)

	// output:
	// baseed 10 sign
	// based 10 unsign
	// binary (humand read)
	// hex

}

func parseArg() int64 {
	args := os.Args[1:]
	fmt.Println(args)
	argsLength := len(args)

	// arg types
	var number int64
	var err error
	// case1 a number
	if argsLength < 1 {
		log.Fatal("No Args Found")
	} else if argsLength == 1 {
		var tmpInt int
		tmpInt, err = strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("%s is not a digital number.", args[1])
		}
		number = int64(tmpInt)
	} else {
		inputType := args[0]
		inputNumber := args[1]
		switch inputType {
		case "-b":
			number, err = strconv.ParseInt(inputNumber, 2, 64)
			if err != nil {
				log.Fatalf("%s is not a valid binary number", inputNumber)
			}
		case "-u":
			var ui uint64
			ui, err = strconv.ParseUint(inputNumber, 10, 64)
			if err != nil {
				log.Fatalf("%s is not a valid unsign int", inputNumber)
			}
			number = int64(ui)
		case "-x":
			number, err = strconv.ParseInt(inputNumber, 16, 64)
			if err != nil {
				log.Fatalf("%s is not a valid hex number", inputNumber)
			}
		default:
			log.Fatalf("Cannot understand the input type: %s", inputType)
		}

	}
	return number
}
