package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func intReducer(list []int, reducer func(acc, cur int) int, start int) int {
	if len(list) == 0 {
		return 0
	} else if len(list) == 1 {
		return list[0]
	}

	acc := start

	for i := range list {
		acc = reducer(acc, list[i])
	}

	return acc
}

func main() {
	// Title
	fmt.Println("Day 6 - part 2 of advent of code 2025!")

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
	problemString := make([][]byte, 0)
	line := []byte{}
	var readError error

	for readError == nil {
		line, readError = reader.ReadBytes('\n')
		if readError == io.EOF {
			break
		} else if readError != nil {
			log.Fatalln("Error reading file: ", err)
		}

		line = bytes.TrimSuffix(line, []byte{'\n'})

		if bytes.ContainsAny(line, "+*") {
			break
		}

		problemString = append(problemString, line)
	}

	// Find location of separators
	sepCols := make([]int, 0)
	sizeX := len(problemString[0])
	sizeY := len(problemString)

xLoop:
	for x := range sizeX {
		for y := range sizeY {
			// fmt.Printf("checking [%d][%d]: %s\n", x, y, string(problemString[y][x]))
			if problemString[y][x] != ' ' {
				continue xLoop
			}
		}
		sepCols = append(sepCols, x)
	}
	sepCols = append(sepCols, sizeX) // for End of Array
	// fmt.Printf("Separator indices: %v\n", sepCols)

	// Calculate Problems
	totalSum := 0

	operatorSet := strings.Fields(string(line))
	problemIndex := 0 // index for sepCols and operatorSet

	currentProblemSet := make([]int, 0)
	for x := range sizeX {
		// fmt.Printf("totalSum: %v | ", totalSum)
		// fmt.Printf("currentProblemSet: (%s) %v\n", operatorSet[problemIndex], currentProblemSet)

		// Flush problems on columns where it's a separator column
		if x == sepCols[problemIndex] {
			if operatorSet[problemIndex] == "+" {
				totalSum += intReducer(currentProblemSet, func(acc, cur int) int { return acc + cur }, 0)
			}
			if operatorSet[problemIndex] == "*" {
				totalSum += intReducer(currentProblemSet, func(acc, cur int) int { return acc * cur }, 1)
			}

			currentProblemSet = []int{}
			problemIndex++
			continue
		}

		// Parse number column by column
		curNumString := ""
		for y := range sizeY {
			curNumString += string(problemString[y][x])
		}

		curNumString = strings.TrimSpace(curNumString)

		curNumParsed, err := strconv.Atoi(curNumString)
		if err != nil {
			log.Fatalln(err)
		}

		currentProblemSet = append(currentProblemSet, curNumParsed)
	}

	// Final flush
	if operatorSet[problemIndex] == "+" {
		totalSum += intReducer(currentProblemSet, func(acc, cur int) int { return acc + cur }, 0)
	}
	if operatorSet[problemIndex] == "*" {
		totalSum += intReducer(currentProblemSet, func(acc, cur int) int { return acc * cur }, 1)
	}

	fmt.Printf("Grand total: %d\n", totalSum)
}
