package main

import (
	"fmt"
	"log"

	"github.com/gouarfig/adventofcode/helpers"
)

func main() {
	presents := helpers.GetLines("input.txt")
	fmt.Printf("found %d presents\n", len(presents))
	paper := 0
	ribbon := 0
	for _, present := range presents {
		dimensions := getSides(string(present))
		paper += wrappingPaper(dimensions)
		ribbon += ribbonLength(dimensions)
	}
	fmt.Printf("1. total wrapping paper needed: %d\n", paper)
	fmt.Printf("2. total ribbon length needed: %d\n", ribbon)
}

func getSides(dimensions string) []int {
	l, w, h := 0, 0, 0
	found, err := fmt.Sscanf(dimensions, "%dx%dx%d", &l, &w, &h)
	if err != nil {
		log.Fatal(err)
	}
	if found != 3 {
		log.Fatalf("incorrect format: %q", dimensions)
	}
	return []int{l, w, h}
}

func getSidesArea(sides []int) []int {
	if len(sides) != 3 {
		log.Fatal("invalid input")
	}
	return []int{
		sides[0] * sides[1],
		sides[1] * sides[2],
		sides[2] * sides[0],
	}
}

func wrappingPaper(dimensions []int) int {
	sides := getSidesArea(dimensions)
	return area(sides) + min(sides)
}

func min(values []int) int {
	if len(values) == 0 {
		return 0
	}
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min
}

func area(sides []int) int {
	area := 0
	for _, side := range sides {
		area += 2 * side
	}
	return area
}

func extractMin(values []int) ([]int, int) {
	if len(values) == 0 {
		return values, 0
	}
	pos := 0
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
			pos = i
		}
	}
	// copy the values into a new slice
	newValues := make([]int, 0, len(values)-1)
	for i, value := range values {
		if i == pos {
			continue
		}
		newValues = append(newValues, value)
	}
	return newValues, min
}

func ribbonLength(sides []int) int {
	twoSides, min1 := extractMin(sides)
	_, min2 := extractMin(twoSides)
	length := 2*(min1+min2) + sides[0]*sides[1]*sides[2]
	// fmt.Printf("ribbon length: %d\n", length)
	return length
}
