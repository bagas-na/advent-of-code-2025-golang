package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Title
	fmt.Println("Day 7 - part 2 of advent of code 2025!")

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
	prevBeam := make([]int, 0)
	nextBeam := make([]int, 0)
	curLine := make([]byte, 0)
	splitCount := 0

	var readError error

	for readError == nil {
		curLine, readError = reader.ReadBytes('\n')
		if readError == io.EOF {
			break
		} else if readError != nil {
			log.Fatalln("Error reading file: ", err)
		}

		curLine = bytes.TrimSpace(curLine)

		// If there's no beam entrance or splitter, skip to next line
		if !bytes.ContainsAny(curLine, "S^") {
			continue
		}
		fmt.Printf("%v\n", strings.Split(string(curLine), ""))

		// fills nextBeam with all '.'
		for range curLine {
			nextBeam = append(nextBeam, 0)
		}

		// Finds where there's S on the current line, add this to nextBeam
		for i := range curLine {
			if curLine[i] == 'S' {
				nextBeam[i] = 1
			}
		}

		// Find where there's a ^ on the current line AND a | on prevBeam
		// Otherise an unhindered beam will continue on

		if len(prevBeam) > 0 {
			for i := range curLine {
				if curLine[i] == '^' && prevBeam[i] > 0 {
					nextBeam[i-1] += prevBeam[i]
					nextBeam[i+1] += prevBeam[i]
					splitCount++
				} else if prevBeam[i] > 0 {
					nextBeam[i] += prevBeam[i]
				}
			}
		}

		// Flush nextBeam to prevBeam
		fmt.Printf("%v\n", nextBeam)
		prevBeam = nextBeam
		nextBeam = make([]int, 0)
	}

	fmt.Printf("Total splits: %d\n", splitCount)

	total := 0
	for _, v := range prevBeam {
		total += v
	}

	fmt.Printf("Total timelines: %d\n", total)
}
