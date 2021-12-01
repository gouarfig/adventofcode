package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	wires := strings.Split(strings.TrimSpace(string(input)), "\n")
	// closest := findClosestIntersection(wires)
	// fmt.Printf("part 1: closest = %d\n", closest)
	steps := findShortestIntersection(wires)
	fmt.Printf("part 2: steps = %d\n", steps)
}

func makePaths(wires []string) [][]point {
	paths := make([][]point, 2)
	for i, wire := range wires {
		path := make([]point, 0)
		pos := point{0, 0}
		instructions := strings.Split(wire, ",")
		for _, instruction := range instructions {
			segment := pos.getPath(instruction)
			pos = segment[len(segment)-1]
			path = append(path, segment...)
		}
		paths[i] = path
	}
	return paths
}

func findClosestIntersection(wires []string) int {
	paths := makePaths(wires)

	closest := 0
	// detect matching points
	for _, point1 := range paths[0] {
		for _, point2 := range paths[1] {
			if point1.equal(point2) {
				dist := distance(point1)
				// log.Printf("found match: x=%d, y=%d, distance=%d", point1.x, point1.y, dist)
				if dist < closest || closest == 0 {
					closest = dist
				}
			}
		}
	}
	return closest
}

func findShortestIntersection(wires []string) int {
	paths := makePaths(wires)

	closest := 0
	// detect matching points
	for steps1, point1 := range paths[0] {
		for steps2, point2 := range paths[1] {
			if point1.equal(point2) {
				log.Printf("found match: x=%d, y=%d, steps1=%d, steps2=%d", point1.x, point1.y, steps1+1, steps2+1)
				if steps1+steps2+2 < closest || closest == 0 {
					closest = steps1 + steps2 + 2
				}
			}
		}
	}
	return closest
}
