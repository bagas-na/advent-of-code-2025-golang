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

func largestJoltage(line string) int {
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

	var readError error
	var line string
	sum := 0

	for readError == nil {
		line, readError = reader.ReadString('\n')
		if readError == io.EOF {
			break
		} else if readError != nil {
			log.Fatalln("Error reading file: ", readError)
		}

		line = strings.TrimRight(line, "\n")
		jolt := largestJoltage(line)
		// fmt.Printf("%s. Jolt: %d\n", line, jolt)
		sum += jolt
	}

	fmt.Println("total output joltage:", sum)
}
