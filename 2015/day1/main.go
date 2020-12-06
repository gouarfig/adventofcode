package main

import (
	"fmt"
	"log"

	"github.com/gouarfig/adventofcode/helpers"
)

const (
	Up   byte = '('
	Down byte = ')'
)

func main() {
	instructions := helpers.GetInput("input.txt")
	floor := getFloor(instructions)
	fmt.Printf("1. go to floor %d\n", floor)
	pos := getBasementInstruction(instructions)
	fmt.Printf("1. got to the basement at position %d\n", pos)
}

func getFloor(instructions []byte) int {
	floor := 0
	for _, instruction := range instructions {
		switch instruction {
		case Up:
			floor++
		case Down:
			floor--
		default:
			log.Fatalf("invalid instruction %s", string(instruction))
		}
	}
	return floor
}

func getBasementInstruction(instructions []byte) int {
	floor := 0
	for i := 0; i < len(instructions); i++ {
		switch instructions[i] {
		case Up:
			floor++
		case Down:
			floor--
			if floor == -1 {
				// got to the basement at last
				return i + 1
			}
		default:
			log.Fatalf("invalid instruction %s", string(instructions[i]))
		}
	}
	log.Fatal("instructions finished without entering the basement")
	return 0 // dead code
}
