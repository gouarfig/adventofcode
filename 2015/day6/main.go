package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gouarfig/adventofcode/2015/day6/part1"
	"github.com/gouarfig/adventofcode/2015/day6/part2"
	"github.com/gouarfig/adventofcode/helpers"
)

func main() {
	fmt.Printf("================= Part 1 =====================\n")
	grid1 := part1.NewGrid()
	fmt.Printf("empty grid: %d lights on\n", part1.CountLights(grid1, true))
	part1.TurnOn(grid1, part1.Coord{0, 0}, part1.Coord{999, 999})
	fmt.Printf("full grid: %d lights on\n", part1.CountLights(grid1, true))
	part1.TurnOff(grid1, part1.Coord{0, 0}, part1.Coord{999, 999})
	part1.Toggle(grid1, part1.Coord{0, 0}, part1.Coord{999, 0})
	fmt.Printf("toggled first line: %d lights on\n", part1.CountLights(grid1, true))
	part1.Toggle(grid1, part1.Coord{0, 0}, part1.Coord{999, 0})
	part1.TurnOn(grid1, part1.Coord{499, 499}, part1.Coord{500, 500})
	fmt.Printf("turn on the middle: %d lights on\n", part1.CountLights(grid1, true))
	part1.TurnOff(grid1, part1.Coord{499, 499}, part1.Coord{500, 500})
	fmt.Printf("back to initial state: %d lights on\n", part1.CountLights(grid1, true))

	input := helpers.GetLines("input.txt")
	for lineID, line := range input {
		instruction := strings.TrimSpace(string(line))
		if len(instruction) == 0 {
			continue
		}
		err := part1.Execute(grid1, instruction)
		if err != nil {
			log.Fatalf("error line %d: %v", lineID+1, err)
		}
	}

	fmt.Printf("\nFirst run of Santa's script: %d lights on\n", part1.CountLights(grid1, true))

	fmt.Printf("================= Part 2 =====================\n")
	grid2 := part2.NewGrid()
	fmt.Printf("empty grid, total brightness: %d\n", part2.CountBrigthness(grid2))
	part2.Increase(grid2, part2.Coord{0, 0}, part2.Coord{999, 999})
	fmt.Printf("full grid, total brightness: %d\n", part2.CountBrigthness(grid2))
	part2.Decrease(grid2, part2.Coord{0, 0}, part2.Coord{999, 999})
	part2.DoubleIncrease(grid2, part2.Coord{0, 0}, part2.Coord{999, 0})
	fmt.Printf("toggled first line, total brightness: %d\n", part2.CountBrigthness(grid2))
	grid2 = part2.NewGrid()
	part2.Increase(grid2, part2.Coord{499, 499}, part2.Coord{500, 500})
	fmt.Printf("turn on the middle, total brightness: %d\n", part2.CountBrigthness(grid2))
	part2.Decrease(grid2, part2.Coord{499, 499}, part2.Coord{500, 500})
	fmt.Printf("back to initial state, total brightness: %d\n", part2.CountBrigthness(grid2))

	for lineID, line := range input {
		instruction := strings.TrimSpace(string(line))
		if len(instruction) == 0 {
			continue
		}
		err := part2.Execute(grid2, instruction)
		if err != nil {
			log.Fatalf("error line %d: %v", lineID+1, err)
		}
	}

	fmt.Printf("\nSecond run of Santa's script, total brightness: %d\n", part2.CountBrigthness(grid2))
}
