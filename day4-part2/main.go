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

func CountAdjacent(grid [][]byte, x, y, sizeX, sizeY int, char byte) int {
	count := 0

	// Check Verticals
	if y > 0 && grid[y-1][x] == char {
		count++
	}
	if y < sizeY-1 && grid[y+1][x] == char {
		count++
	}

	// Check horizontals
	if x > 0 && grid[y][x-1] == char {
		count++
	}
	if x < sizeX-1 && grid[y][x+1] == char {
		count++
	}

	// Check diagonals
	if y > 0 && x > 0 && grid[y-1][x-1] == char {
		count++
	}
	if y > 0 && x < sizeX-1 && grid[y-1][x+1] == char {
		count++
	}
	if y < sizeY-1 && x > 0 && grid[y+1][x-1] == char {
		count++
	}
	if y < sizeY-1 && x < sizeX-1 && grid[y+1][x+1] == char {
		count++
	}

	return count
}

func UpdateGrid(grid [][]byte) (updatedGrid [][]byte, accessible int) {
	for i := range grid {
		line := grid[i]
		lineCopy := make([]byte, len(line))
		copy(lineCopy, line)
		updatedGrid = append(updatedGrid, lineCopy)
	}

	sizeY := len(grid)
	sizeX := len(grid[0])

	countAccessible := 0

	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			if grid[y][x] != '@' {
				continue
			}

			if CountAdjacent(grid, x, y, sizeX, sizeY, '@') < 4 {
				countAccessible++
				updatedGrid[y][x] = '.'
			}
		}
	}

	accessible = countAccessible
	return
}

func main() {
	fmt.Println("Day 4 - part 2 of advent of code 2025!")
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

	println("Initial state")
	grid := ReadToEOF(reader)
	for i := range grid {
		fmt.Printf("%s\n", grid[i])
	}
	println()

	totalAccessible := 0

	updatedGrid, removed := UpdateGrid(grid)
	// fmt.Printf("Remove %d rolls of paper:\n", removed)
	// for i := range updatedGrid {
	// 	fmt.Printf("%s\n", updatedGrid[i])
	// }
	// println()
	totalAccessible += removed

	for removed > 0 {
		updatedGrid, removed = UpdateGrid(updatedGrid)

		// fmt.Printf("Remove %d rolls of paper:\n", removed)
		// for i := range updatedGrid {
		// 	fmt.Printf("%s\n", updatedGrid[i])
		// }
		// println()

		totalAccessible += removed
	}

	println("Final state")
	for i := range updatedGrid {
		fmt.Printf("%s\n", updatedGrid[i])
	}
	println()
	fmt.Printf("A total of %d rolls of paper can be removed.\n", totalAccessible)
}
