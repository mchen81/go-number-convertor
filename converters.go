package main

import (
	"strconv"
)

func toBinary(number uint32) string {
	return strconv.FormatInt(int64(number), 2)
}

func toHumanReadBinary(number uint32) string {

	var chars [39]byte
	var m byte
	counter := 0
	for i := range chars {

		if counter == 4 {
			chars[38-i] = ' '
			counter = 0
			continue
		}

		if number > 0 {
			m = byte((number % 2) + '0')
			number /= 2
			chars[38-i] = m
		} else {
			chars[38-i] = '0'
		}

		counter++
	}

	return string(chars[:])
}

func toHex(number uint32) string {
	return strconv.FormatInt(int64(number), 16)
}
