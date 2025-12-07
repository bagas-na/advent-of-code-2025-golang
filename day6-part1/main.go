package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Title
	fmt.Println("Day 6 - part 1 of advent of code 2025!")

	// Args manipulation
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <file>")
		return
	}
	fileName := os.Args[1]

	// Opening file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)

	// Read file
	problemSet := make([][]int, 0)
	line := ""
	var readError error

	// Read and parse first part only (before any operators)
	for readError == nil {
		line, readError = reader.ReadString('\n')
		if readError == io.EOF {
			break
		} else if readError != nil {
			log.Fatalln("Error reading file: ", err)
		}

		line = strings.TrimSpace(line)

		if strings.ContainsAny(line, "*+") {
			break
		}

		lineNums := strings.Fields(line)
		lineParsedNums := make([]int, 0, len(lineNums))

		for _, v := range lineNums {
			parsedNum, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalln(err)
			}

			lineParsedNums = append(lineParsedNums, parsedNum)
		}

		problemSet = append(problemSet, lineParsedNums)
	}

	// fmt.Printf("Problem set %v\n", problemSet)

	// Parse operators
	operatorSet := strings.Fields(line)
	// fmt.Printf("Operators %v\n", operatorSet)

	// Calculate
	totalSum := 0
	for i := 0; i < len(problemSet[0]); i++ {
		columnSum := 0
		if operatorSet[i] == "*" {
			columnSum = 1
		}

		for j := 0; j < len(problemSet); j++ {
			if operatorSet[i] == "+" {
				columnSum += problemSet[j][i]
			}
			if operatorSet[i] == "*" {
				columnSum *= problemSet[j][i]
			}
		}
		totalSum += columnSum
	}

	println(totalSum)
}
