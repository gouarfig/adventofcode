package main

import (
	"fmt"
	"log"

	"github.com/gouarfig/adventofcode/helpers"
)

type Coord struct {
	x int
	y int
}

func main() {
	// === first part
	houses := make(map[Coord]int)
	var x, y int
	input := helpers.GetInput("input.txt")
	// deliver on (0, 0)
	deliver(houses, x, y)

	for _, char := range input {
		x, y = nextHouse(char, x, y)
		deliver(houses, x, y)
	}

	// Check the delivered houses
	deliveries, presents := getDeliveries(houses)
	fmt.Printf("1. deliveries: %d houses, presents: %d\n", deliveries, presents)

	// === second part
	var santaX, santaY, robotX, robotY int
	houses = make(map[Coord]int)
	// 2 deliveries on (0, 0)
	deliver(houses, santaX, santaY)
	deliver(houses, robotX, robotY)
	pos := 0
	for pos < len(input) {
		santaX, santaY = nextHouse(input[pos], santaX, santaY)
		deliver(houses, santaX, santaY)
		pos++
		robotX, robotY = nextHouse(input[pos], robotX, robotY)
		deliver(houses, robotX, robotY)
		pos++
	}
	deliveries, presents = getDeliveries(houses)
	fmt.Printf("2. deliveries: %d houses, presents: %d\n", deliveries, presents)
}

func nextHouse(move byte, x, y int) (int, int) {
	switch move {
	case '<':
		x--
	case '>':
		x++
	case '^':
		y--
	case 'v':
		y++
	default:
		log.Fatal("incorrect move")
	}
	return x, y
}

func deliver(houses map[Coord]int, x, y int) {
	index := Coord{x, y}
	if _, ok := houses[index]; !ok {
		houses[index] = 0
	}
	houses[index]++
}

func getDeliveries(houses map[Coord]int) (int, int) {
	deliveries := 0
	totalPresents := 0
	for index, presents := range houses {
		if presents == 0 {
			fmt.Printf("no present at %d, %d\n", index.x, index.y)
		}
		deliveries++
		totalPresents += presents
	}
	return deliveries, totalPresents
}
