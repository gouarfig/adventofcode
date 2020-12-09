package part1

import (
	"fmt"
	"strings"
)

type Grid [][]bool

type Coord struct {
	X uint
	Y uint
}

func NewGrid() Grid {
	grid := make([][]bool, 1000)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]bool, 1000)
	}
	return grid
}

func turnOnOff(lights Grid, from, to Coord, value bool) {
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			lights[x][y] = value
		}
	}
}

func Toggle(lights Grid, from, to Coord) {
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			lights[x][y] = !lights[x][y]
		}
	}
}

func TurnOn(lights Grid, from, to Coord) {
	turnOnOff(lights, from, to, true)
}

func TurnOff(lights Grid, from, to Coord) {
	turnOnOff(lights, from, to, false)
}

func CountLights(lights Grid, value bool) int {
	count := 0
	for _, row := range lights {
		for _, column := range row {
			if column == value {
				count++
			}
		}
	}
	return count
}

var (
	commands = map[string]func(lights Grid, from, to Coord){
		"turn on":  TurnOn,
		"turn off": TurnOff,
		"toggle":   Toggle,
	}
)

func Execute(lights Grid, instruction string) error {
	command := ""
	for name := range commands {
		if strings.HasPrefix(instruction, name) {
			command = name
			break
		}
	}
	if command == "" {
		return fmt.Errorf("invalid command %q", instruction)
	}
	instruction = strings.TrimPrefix(instruction, command)
	// for the remaining of the instruction, a scanf is quite fitted
	var x1, y1, x2, y2 uint
	found, err := fmt.Sscanf(instruction, "%d,%d through %d,%d", &x1, &y1, &x2, &y2)
	if err != nil {
		return err
	}
	if found != 4 {
		return fmt.Errorf("cannot parse %q", instruction)
	}
	// fmt.Printf("%s from (%d, %d) to (%d, %d)\n", command, x1, y1, x2, y2)
	commands[command](lights, Coord{x1, y1}, Coord{x2, y2})
	return nil
}
