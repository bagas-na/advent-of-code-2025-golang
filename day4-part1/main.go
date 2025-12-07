package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadToEOF(reader *bufio.Reader) [][]byte {
	var readError error
	var line []byte
	var grid [][]byte

	for readError == nil {
		line, readError = reader.ReadBytes('\n')
		if readError == io.EOF {
			break
		} else if readError != nil {
			log.Fatalln("Error reading file: ", readError)
		}
		line = bytes.TrimSpace(line)

		grid = append(grid, line)
	}

	return grid
}

func main() {
	fmt.Println("Day 4 - part 1 of advent of code 2025!")
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

	grid := ReadToEOF(reader)

	sizeY := len(grid)
	sizeX := len(grid[0])

	countAccessible := 0

	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			if grid[y][x] != '@' {
				continue
			}

			count := 0

			// Check Verticals
			if y > 0 && grid[y-1][x] == '@' {
				count++
			}
			if y < sizeY-1 && grid[y+1][x] == '@' {
				count++
			}

			// Check horizontals
			if x > 0 && grid[y][x-1] == '@' {
				count++
			}
			if x < sizeX-1 && grid[y][x+1] == '@' {
				count++
			}

			// Check diagonals
			if y > 0 && x > 0 && grid[y-1][x-1] == '@' {
				count++
			}
			if y > 0 && x < sizeX-1 && grid[y-1][x+1] == '@' {
				count++
			}
			if y < sizeY-1 && x > 0 && grid[y+1][x-1] == '@' {
				count++
			}
			if y < sizeY-1 && x < sizeX-1 && grid[y+1][x+1] == '@' {
				count++
			}

			if count < 4 {
				countAccessible++
			}
		}
	}

	println("Accessible rolls of paper: ", countAccessible)
}
