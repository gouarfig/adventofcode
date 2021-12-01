package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	part1, part2 := totalFuelPart1And2(file)
	fmt.Printf("part 1, total fuel = %d\npart 2, total fuel = %d\n", part1, part2)
}

func fuel(mass int64) int64 {
	return int64(math.Floor(float64(mass)/3)) - 2
}

func totalFuel(mass int64) int64 {
	total := fuel(mass)
	value := total
	for value > 0 {
		value = fuel(value)
		if value > 0 {
			total += value
		}
	}
	return total
}

func totalFuelPart1And2(r io.Reader) (part1, part2 int64) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		mass, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		part1 += fuel(mass)
		part2 += totalFuel(mass)
	}
	return part1, part2
}
