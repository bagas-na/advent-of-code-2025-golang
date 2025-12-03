package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func largestJoltage2(line string) int {
	tensIndex := 0
	tens := math.MinInt
	// onesIndex := 0
	ones := math.MinInt

	for i := 0; i < len(line)-1; i++ {
		currentNum, _ := strconv.Atoi(string(line[i]))
		if currentNum > tens {
			tens = currentNum
			tensIndex = i
		}
	}

	for i := tensIndex + 1; i < len(line); i++ {
		currentNum, _ := strconv.Atoi(string(line[i]))
		if currentNum > ones {
			ones = currentNum
			// onesIndex = i
		}
	}

	return tens*10 + ones
}

func largestJoltageN(line string, digits int) uint64 {
	if len(line) <= digits {
		num, _ := strconv.ParseUint(line, 10, 64)
		return num
	}

	// fmt.Printf("Calculating jolt for bank %q...\n", line)

	cursor := 0 // pointer where the "next" digit starts their search
	jolt := ""

	for digitOrder := range digits {
		maxNum := math.MinInt
		maxNumIndex := cursor

		// fmt.Printf("\tDigit #%d. Finding in substring %q\n", digitOrder, line[cursor:len(line)-digits+digitOrder+1])

		for i := cursor; i < len(line)-digits+digitOrder+1; i++ {
			currentNum, _ := strconv.Atoi(string(line[i]))
			if currentNum > maxNum {
				maxNum = currentNum
				maxNumIndex = i
			}
		}

		jolt += string(line[maxNumIndex])
		cursor = maxNumIndex + 1
	}

	num, _ := strconv.ParseUint(jolt, 10, 64)

	return num
}

func main() {
	fmt.Println("Day 3 of advent of code")
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

	var readError error
	var line string
	var sum uint64

	for readError == nil {
		line, readError = reader.ReadString('\n')
		if readError == io.EOF {
			break
		} else if readError != nil {
			log.Fatalln("Error reading file: ", readError)
		}

		line = strings.TrimRight(line, "\n")
		jolt := largestJoltageN(line, 12)
		// fmt.Printf("%s. Jolt: %d\n", line, jolt)
		sum += jolt
	}

	fmt.Println("total output joltage:", sum)
}
