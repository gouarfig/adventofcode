package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	base := strings.TrimSpace(string(raw))
	memory := loadMemory(base)
	// initial setup
	memory[1] = 12
	memory[2] = 2
	runProgram(memory)
	fmt.Printf("part 1: output is %d\n", memory[0])

	expected := int64(19690720)
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			memory := loadMemory(base)
			memory[1] = int64(noun)
			memory[2] = int64(verb)
			runProgram(memory)
			if memory[0] == expected {
				value := 100*noun + verb
				fmt.Printf("part2: value is %d (match noun = %d, verb = %d)\n", value, noun, verb)
			}
		}
	}
}

func loadMemory(input string) []int64 {
	temp := strings.Split(input, ",")
	memory := make([]int64, len(temp))
	for pos, value := range temp {
		opcode, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			panic(err)
		}
		memory[pos] = opcode
	}
	return memory
}

func saveMemory(memory []int64) string {
	temp := make([]string, len(memory))
	for pos, opcode := range memory {
		temp[pos] = strconv.FormatInt(opcode, 10)
	}
	return strings.Join(temp, ",")
}

func runProgram(memory []int64) {
	const instructionAdd = 1
	const instructionMultiply = 2
	const instructionHalt = 99
	for pos := 0; pos < len(memory); pos++ {
		opcode := memory[pos]
		switch opcode {
		case instructionAdd:
			pointer1 := memory[pos+1]
			pointer2 := memory[pos+2]
			pointer3 := memory[pos+3]
			memory[pointer3] = memory[pointer1] + memory[pointer2]
			pos += 3
		case instructionMultiply:
			pointer1 := memory[pos+1]
			pointer2 := memory[pos+2]
			pointer3 := memory[pos+3]
			memory[pointer3] = memory[pointer1] * memory[pointer2]
			pos += 3
		case instructionHalt:
			return
		default:
			panic(fmt.Sprintf("invalid opcode %d", opcode))
		}
	}
}
