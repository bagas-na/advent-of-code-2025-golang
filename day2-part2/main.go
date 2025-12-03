package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func expandRangeToString(rangeStr string) []string {
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

	arr := make([]string, 0)

	for i := start; i <= end; i++ {
		arr = append(arr, strconv.FormatUint(i, 10))
	}

	return arr
}

func hasRepeatingDigits(numStr string) bool {
	for repeatLen := 1; repeatLen <= len(numStr)/2; repeatLen++ {
		// fmt.Printf("Num: %v | ", numStr)
		if repeatLen > 1 && len(numStr)%repeatLen != 0 {
			continue
		}

		// fmt.Printf("Comparing %s with %s\n", strings.Repeat(numStr[:repeatLen], len(numStr)/repeatLen), numStr)
		if strings.Repeat(numStr[:repeatLen], len(numStr)/repeatLen) == numStr {
			return true
		}
	}

	return false
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
		if readError != nil {
			log.Fatal(readError)
		}

		rangeString = strings.TrimRight(rangeString, ",\n")

		arr := expandRangeToString(rangeString)
		// fmt.Printf("%q\n", arr)

		for _, v := range arr {
			// println(v)
			if hasRepeatingDigits(v) {
				// fmt.Printf("Digit with repeats: %v\n", v)
				num, _ := strconv.ParseUint(v, 10, 64)
				sum += num
			}
		}
	}

	fmt.Printf("Sum of repeated numbers: %v\n", sum)
}
