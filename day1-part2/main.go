package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func amountGoThroughZero(initialPosition int, rotationAmount int, dialSize int) (finalPosition, countZero int) {
	if rotationAmount == 0 {
		panic("Rotation must not be 0")
	}

	if dialSize <= 1 {
		panic("Dial size must be larger than 1")
	}

	// direction := "no"
	// if rotationAmount > 0 {
	// 	direction = "right"
	// } else if rotationAmount < 0 {
	// 	direction = "left"
	// }

	finalPosition = initialPosition + rotationAmount

	if finalPosition == 0 {
		countZero = 1
		// (1) means the final value is zero, and it doesn't rotate through '0' before the dial gets tehere
		// fmt.Printf("(1) Dial at %d is rotated %q %d times, to %d, passing '0' %d times.\n", initialPosition, direction, int(math.Abs(float64(rotationAmount))), finalPosition, countZero)
		return
	}

	if finalPosition < 0 {
		for finalPosition < 0 {
			finalPosition += dialSize
			countZero++
		}

		// Edge cases where the initial value is zero, where turning the dial to the left before full rotation (to get to '0' again) is already counted
		if initialPosition == 0 {
			countZero--
		}

		// Edge case where if the final value is zero, it's not being counted at the for loop above
		if finalPosition == 0 {
			countZero++
		}
		// (2) means the dial is rotated to the left, and it rotate through '0' atleast once
		// fmt.Printf("(2) Dial at %d is rotated %q %d times, to %d, passing '0' %d times.\n", initialPosition, direction, int(math.Abs(float64(rotationAmount))), finalPosition, countZero)
		return
	}

	if finalPosition >= dialSize {
		for finalPosition >= dialSize {
			finalPosition -= dialSize
			countZero++
		}

		// (3) means the dial is rotated to the right, and it rotates through '0' at least once
		// fmt.Printf("(3) Dial at %d is rotated %q %d times, to %d, passing '0' %d times.\n", initialPosition, direction, int(math.Abs(float64(rotationAmount))), finalPosition, countZero)
		return

	}

	// (4) means the dial is rotated either to the left or to the right, without ever touching the number '0'
	// fmt.Printf("(4) Dial at %d is rotated %q %d times, to %d, passing '0' %d times.\n", initialPosition, direction, int(math.Abs(float64(rotationAmount))), finalPosition, countZero)
	return
}

func main() {
	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)

	countZero := 0
	dialPos := 50
	fmt.Printf("Current dial position: %d, '0' counted: %d\n", dialPos, countZero)

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
			count := 0
			dialPos, count = amountGoThroughZero(dialPos, -amount, 100)
			countZero += count
		}

		if strings.HasPrefix(line, "R") {
			amount, _ = strconv.Atoi(line[1:])
			count := 0
			dialPos, count = amountGoThroughZero(dialPos, amount, 100)
			countZero += count
		}
	}

	fmt.Printf("Final position: %d; '0' counted: %d\n", dialPos, countZero)
}
