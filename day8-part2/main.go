package main

import (
	"bufio"
	"bytes"
	"cmp"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
)

type Point3D [3]int64

type PointDistance struct {
	indexA   int
	indexB   int
	distance int64
}

func pointDistance(a, b Point3D) int64 {
	return (a[0]-b[0])*(a[0]-b[0]) +
		(a[1]-b[1])*(a[1]-b[1]) +
		(a[2]-b[2])*(a[2]-b[2])
}

type PointSet struct {
	length   int
	elements map[Point3D]interface{}
}

func NewPointSet() PointSet {
	return PointSet{
		length:   0,
		elements: make(map[Point3D]interface{}),
	}
}

// Checks if this set contains a point
func (p *PointSet) has(point Point3D) bool {
	_, ok := p.elements[point]
	return ok
}

// Adds a point to this set. Returns an error if the point already exists
func (p *PointSet) add(point Point3D) error {
	_, ok := p.elements[point]
	if ok {
		return errors.New("duplicate value")
	}
	p.elements[point] = struct{}{}
	p.length += 1

	return nil
}

// Empties out the set. Should never return an error
func (p *PointSet) clear() error {
	p.length = 0
	p.elements = make(map[Point3D]interface{})

	return nil
}

// Delete an entry from the set
func (p *PointSet) delete(point Point3D) error {
	_, ok := p.elements[point]
	if !ok {
		return errors.New("value does not exist")
	}
	delete(p.elements, point)
	p.length -= 1

	return nil
}

// Returns a list of the entries, in no particular order
func (p *PointSet) entries() []Point3D {
	out := make([]Point3D, 0, p.length)

	for k := range p.elements {
		out = append(out, k)
	}

	return out
}

func (p *PointSet) size() int {
	return p.length
}

// Joins another set to this set, and empties it out.
func (p *PointSet) union(set *PointSet) error {
	for _, v := range set.entries() {
		err := p.add(v)
		if err != nil {
			return errors.New("error joining sets")
		}
		err = set.delete(v)
		if err != nil {
			return errors.New("error joining sets")
		}
	}

	return nil
}

func main() {
	// Title
	fmt.Println("Day 8 - part 2 of advent of code 2025!")

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

	pointList := make([]Point3D, 0)

	curLine := make([]byte, 0)
	var readError error

	for readError == nil {
		curLine, readError = reader.ReadBytes('\n')
		if readError == io.EOF {
			break
		} else if readError != nil {
			log.Fatalln("Error reading file: ", err)
		}

		curLine = bytes.TrimSpace(curLine)

		var point Point3D

		split := bytes.Split(curLine, []byte(","))

		for i := range point {
			num, err := strconv.ParseInt(string(split[i]), 10, 64)
			if err != nil {
				log.Fatalln(err)
			}
			point[i] = num
		}

		pointList = append(pointList, point)
	}
	// fmt.Printf("%#+v\n", pointList)
	fmt.Printf("Point list length: %d\n", len(pointList))

	// Calculate pair-wise distances
	distances := make([]PointDistance, 0)

	for i := 0; i < len(pointList); i++ {
		for j := i + 1; j < len(pointList); j++ {
			distance := pointDistance(pointList[i], pointList[j])
			// fmt.Printf("%d ", distance)

			distances = append(distances, PointDistance{
				indexA:   i,
				indexB:   j,
				distance: distance,
			})
		}
		// fmt.Printf("\n")
	}

	slices.SortStableFunc(distances, func(a, b PointDistance) int {
		return cmp.Compare(a.distance, b.distance)
	})

	fmt.Printf("Connections length: %d\n", len(distances))

	// fmt.Printf("%v\n", distances)

	// Connect junction boxes
	pointsCount := len(pointList)
	distancesCount := len(distances) // amount of connections to be made
	ListOfSets := make([]PointSet, 0)

	for i := range distancesCount {
		boxA := pointList[distances[i].indexA]
		boxB := pointList[distances[i].indexB]

		if len(ListOfSets) == 0 {
			newSet := NewPointSet()
			newSet.add(boxA)
			newSet.add(boxB)

			ListOfSets = append(ListOfSets, newSet)
			continue
		}

		// Checks whether boxA or boxB is in a set. -1 means it's not in a set
		boxASetIndex := -1
		boxBSetIndex := -1

		for i := range ListOfSets {
			if ListOfSets[i].has(boxA) {
				boxASetIndex = i
			}
			if ListOfSets[i].has(boxB) {
				boxBSetIndex = i
			}
		}

		// 4 cases:
		// (4) A and B both are not in a set; new set
		// (3) One of A or B are in a set, the other is not; add
		// (1) A and B both are in the same set; nothing happens
		// (2) A and B both are both in a different set; union
		if boxASetIndex == -1 && boxBSetIndex == -1 {
			newSet := NewPointSet()
			newSet.add(boxA)
			newSet.add(boxB)

			ListOfSets = append(ListOfSets, newSet)
		} else if boxASetIndex == -1 {
			ListOfSets[boxBSetIndex].add(boxA)

			if ListOfSets[boxBSetIndex].size() == pointsCount {
				fmt.Printf("Last boxes: %#v, and %#v\n", boxA, boxB)
			}

		} else if boxBSetIndex == -1 {
			ListOfSets[boxASetIndex].add(boxB)

			if ListOfSets[boxASetIndex].size() == pointsCount {
				fmt.Printf("Last boxes: %#v, and %#v\n", boxA, boxB)
			}

		} else if boxASetIndex != boxBSetIndex {
			ListOfSets[boxASetIndex].union(&ListOfSets[boxBSetIndex])

			if ListOfSets[boxASetIndex].size() == pointsCount {
				fmt.Printf("Last boxes: %#v, and %#v\n", boxA, boxB)
			}

		}

		// fmt.Println("BoxA set length: ", func() int {
		// 	if boxASetIndex == -1 {
		// 		return -1
		// 	} else {
		// 		return ListOfSets[boxASetIndex].size()
		// 	}
		// }())
		// fmt.Println("BoxB set length: ", func() int {
		// 	if boxBSetIndex == -1 {
		// 		return -1
		// 	} else {
		// 		return ListOfSets[boxBSetIndex].size()
		// 	}
		// }())
	}

	fmt.Println("==========================")

	slices.SortStableFunc(ListOfSets, func(a, b PointSet) int {
		return -cmp.Compare(a.size(), b.size())
	})

	for i := range 3 {
		fmt.Printf("len: %d \n", ListOfSets[i].size())
	}

}
