package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	fixtures := []struct {
		instruction string
		path        []point
	}{
		{"L3", []point{{-1, 0}, {-2, 0}, {-3, 0}}},
		{"R3", []point{{1, 0}, {2, 0}, {3, 0}}},
		{"U3", []point{{0, -1}, {0, -2}, {0, -3}}},
		{"D3", []point{{0, 1}, {0, 2}, {0, 3}}},
	}

	for _, fixture := range fixtures {
		assert.ElementsMatch(t, fixture.path, point{0, 0}.getPath(fixture.instruction))
	}
}

func TestExamplePart1(t *testing.T) {
	example := []struct {
		wires    []string
		distance int
		steps    int
	}{
		{[]string{
			"R8,U5,L5,D3",
			"U7,R6,D4,L4",
		}, 6, 30},
		{[]string{
			"R75,D30,R83,U83,L12,D49,R71,U7,L72",
			"U62,R66,U55,R34,D71,R55,D58,R83",
		}, 159, 610},
		{[]string{
			"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
		}, 135, 410},
	}

	for _, fixture := range example {
		closest := findClosestIntersection(fixture.wires)
		assert.Equal(t, fixture.distance, closest)

		steps := findShortestIntersection(fixture.wires)
		assert.Equal(t, fixture.steps, steps)
	}
}
