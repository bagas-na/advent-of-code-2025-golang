package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func expandRangeString(rangeStr string) []uint64 {
	startStr, endStr, isRange := strings.Cut(rangeStr, "-")
	if !isRange {
		log.Panicln("Invalid range string")
	}

	start, err := strconv.ParseUint(startStr, 10, 64)
	if err != nil {
		log.Panicln(err)
	}

	end, err := strconv.ParseUint(endStr, 10, 64)
	if err != nil {
		log.Panicln(err)
	}

	intArray := make([]uint64, 0)

	for i := start; i <= end; i++ {
		intArray = append(intArray, i)
	}

	return intArray
}

func isRepeatedTwice(num uint64) bool {
	tempNum := num
	count := 1 // counts the number of characters needed to represent the number. E.g. 20 => 2, and 123 => 3

	for tempNum > 9 {
		tempNum = tempNum / 10
		count++
	}

	// Number can't repeat twice if it has an odd number of character representation
	if count%2 == 1 {
		return false
	}

	var divisor uint64 = 1
	for range count / 2 {
		divisor *= 10
	}

	return num/divisor == num%divisor
}

func main() {
	fmt.Println("Day 2 of advent of code")
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <file>")
		return
	}

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)

	var sum uint64 = 0

	var readError error
	rangeString := ""

	for readError == nil {
		rangeString, readError = reader.ReadString(',')
		if err != nil {
			log.Fatal(err)
		}

		rangeString = strings.TrimRight(rangeString, ",\n")

		intArray := expandRangeString(rangeString)

		for _, v := range intArray {
			if isRepeatedTwice(v) {
				// fmt.Printf("%v is repeated twice.\n", v)
				sum += v
			}
		}
	}

	fmt.Printf("Sum of repeated numbers: %v\n", sum)
}
