package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Modulus(number int, modulo int) int {
	if modulo <= 0 {
		log.Panic("Modulo must be a non negative integer")
	}

	mod := number % modulo

	if mod < 0 {
		return mod + modulo
	}
	return mod
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)

	countZero := 0
	dial := 50
	fmt.Printf("Current dial position: %d, '0' counted: %d\n", dial, countZero)

	var readError error
	line := ""

	for readError == nil {
		line, readError = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		line = strings.TrimRight(line, "\n")
		amount := 0

		if strings.HasPrefix(line, "L") {
			amount, _ = strconv.Atoi(line[1:])
			dial = Modulus(dial-amount, 100)
		}

		if strings.HasPrefix(line, "R") {
			amount, _ = strconv.Atoi(line[1:])
			dial = Modulus(dial+amount, 100)
		}

		if dial == 0 {
			countZero++
		}

		// if line != "" {
		// 	fmt.Printf("Rotating dial to %s direction, %d times. Current position: %d. '0' counted: %d.\n", string(line[0]), amount, dial, countZero)
		// }
	}

	fmt.Printf("Final '0' counted: %d\n", countZero)
}
