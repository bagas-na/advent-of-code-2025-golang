package main

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type numRange struct {
	start int
	end   int
}

func (r numRange) has(num int) bool {
	return num >= r.start && num <= r.end
}

func (r numRange) String() string {
	return fmt.Sprintf("[%d-%d]", r.start, r.end)
}

// Merge two ranges into a contiguous range. If not possible, return an error
func mergeTwoRanges(a, b numRange) (numRange, error) {
	// WLOG, a.start < b.start. Swap otherwise
	if a.start > b.start {
		temp := a
		a = b
		b = temp
	}

	if b.start > a.end+1 {
		return a, fmt.Errorf("cannot merge: %s and %s", a, b)
	}

	if b.end <= a.end {
		return numRange{start: a.start, end: a.end}, nil
	}

	return numRange{start: a.start, end: b.end}, nil
}

func strToRange(strRange string) numRange {
	parts := strings.Split(strRange, "-")
	nums := make([]int, 0, 2)
	for _, v := range parts {
		integer, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		nums = append(nums, integer)
	}

	return numRange{start: nums[0], end: nums[1]}
}

func main() {
	// Title
	fmt.Println("Day 5 - part 1 of advent of code 2025!")

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
	rangeList := make([]numRange, 0)
	line := ""
	var readError error

	// Read and parse first part only
	for readError == nil {
		line, readError = reader.ReadString('\n')
		if readError == io.EOF {
			break
		} else if readError != nil {
			log.Fatalln("Error reading file: ", readError)
		}
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		rangeList = append(rangeList, strToRange(line))
	}
	fmt.Printf("range list:\n%v\n", rangeList)

	// Sort range list
	slices.SortStableFunc(rangeList, func(a, b numRange) int {
		return cmp.Compare(a.start, b.start)
	})
	fmt.Printf("sorted range list:\n%v\n", rangeList)

	// Merge range list
	mergedRanges := make([]numRange, 0)
	curRange := rangeList[0]
	var mergeError error

	for _, v := range rangeList {
		curRange, mergeError = mergeTwoRanges(curRange, v)
		if mergeError != nil {
			mergedRanges = append(mergedRanges, curRange)
			curRange = v
			mergeError = nil
		}
	}
	mergedRanges = append(mergedRanges, curRange)
	fmt.Printf("merged range list:\n%v\n", mergedRanges)

	// Count integers in range list
	freshCount := 0
	for _, v := range mergedRanges {
		freshCount += v.end - v.start + 1
	}

	fmt.Printf("Count of fresh IDs: %d\n", freshCount)
}
