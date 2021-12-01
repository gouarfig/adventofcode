package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	Right = 'R'
	Left  = 'L'
	Up    = 'U'
	Down  = 'D'
)

type point struct {
	x int
	y int
}

func (p point) equal(q point) bool {
	return p.x == q.x && p.y == q.y
}

func (p point) getPath(instruction string) []point {
	direction := instruction[0]
	length, err := strconv.Atoi(instruction[1:])
	if err != nil {
		panic(err)
	}
	path := make([]point, 0, length)

	switch direction {
	case Left:
		for i := p.x - 1; i > p.x-1-length; i-- {
			path = append(path, point{i, p.y})
		}
	case Right:
		for i := p.x + 1; i < p.x+1+length; i++ {
			path = append(path, point{i, p.y})
		}
	case Up:
		for i := p.y - 1; i > p.y-1-length; i-- {
			path = append(path, point{p.x, i})
		}
	case Down:
		for i := p.y + 1; i < p.y+1+length; i++ {
			path = append(path, point{p.x, i})
		}
	default:
		panic(fmt.Sprintf("unknown direction %s", string(direction)))
	}
	return path
}

func distance(p point) int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}
