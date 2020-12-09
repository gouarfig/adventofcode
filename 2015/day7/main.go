package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gouarfig/adventofcode/helpers"
)

func main() {
	input := helpers.GetLines("input.txt")
	rules, err := loadInput(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("found %d rules\n", len(rules))

	fmt.Printf("=== part 1===\n")
	wires := make(map[string]uint16, len(rules))
	a, err := resolveWire("a", rules, wires)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("final result: a = %d\n", a)

	fmt.Printf("=== part 2===\n")
	wires = make(map[string]uint16, len(rules))
	wires["b"] = a
	a, err = resolveWire("a", rules, wires)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("final result: a = %d\n", a)
}

func loadInput(input [][]byte) (map[string]Rule, error) {
	var err error
	wires := make(map[string]Rule)
	fmt.Printf("loaded %d instructions\n", len(input))
	for lineID, line := range input {
		instruction := strings.TrimSpace(string(line))
		if len(instruction) == 0 {
			continue
		}
		operation := strings.Split(instruction, " -> ")
		if len(operation) != 2 {
			return nil, fmt.Errorf("line %d: invalid instruction %q", lineID+1, instruction)
		}
		wire := strings.TrimSpace(operation[1])
		if _, ok := wires[wire]; ok {
			return nil, fmt.Errorf("wire %s already defined", wire)
		}
		wires[wire], err = loadOperation(operation[0])
		if err != nil {
			return nil, fmt.Errorf("line %d: %v", lineID+1, err)
		}
	}
	return wires, nil
}

func loadOperation(operation string) (Rule, error) {
	var found bool
	var op Operator
	var sources []string
	for _, operator := range Operators {
		if strings.Contains(operation, operator.Code) {
			found = true
			op = operator.Operator
			switch operator.Type {
			case TwoInput:
				sources = strings.Split(operation, operator.Code)
				if len(sources) != 2 {
					return Rule{}, fmt.Errorf("invalid operation %q", operation)
				}
				sources[0] = strings.TrimSpace(sources[0])
				sources[1] = strings.TrimSpace(sources[1])

			case OneInput:
				sources = []string{strings.TrimSpace(strings.TrimPrefix(operation, operator.Code))}
			}
			break
		}
	}
	if !found {
		// no operator => it should just be a number, or a wire name
		sources = []string{strings.TrimSpace(operation)}
	}
	return Rule{
		Operator: op,
		Sources:  sources,
	}, nil
}

func resolveWire(wire string, rules map[string]Rule, wires map[string]uint16) (uint16, error) {
	if value, ok := wires[wire]; ok {
		// return the value already calculated
		return value, nil
	}
	var err error
	var resolve uint16
	if rule, ok := rules[wire]; ok {
		// resolve needed values
		var first, second uint16

		if len(rule.Sources) >= 1 {
			first, err = getValue(rule.Sources[0], rules, wires)
			if err != nil {
				return 0, fmt.Errorf("rule %+v: %v", rule, err)
			}
		}

		if len(rule.Sources) == 2 {
			second, err = getValue(rule.Sources[1], rules, wires)
			if err != nil {
				return 0, fmt.Errorf("rule %+v: %v", rule, err)
			}
		}
		switch rule.Operator {
		case Noop:
			resolve = first

		case Not:
			resolve = ^first

		case LShift:
			resolve = first << second

		case RShift:
			resolve = first >> second

		case And:
			resolve = first & second

		case Or:
			resolve = first | second
		}
		// fmt.Printf("%s: %d\n", wire, resolve)
		wires[wire] = resolve
	} else {
		return 0, fmt.Errorf("rule not found for %q", wire)
	}
	return resolve, nil
}

func getValue(wire string, rules map[string]Rule, wires map[string]uint16) (uint16, error) {
	var value uint16
	// try if it is a number
	temp, err := strconv.ParseUint(wire, 10, 16)
	if err == nil {
		// this is a number
		value = uint16(temp)
	} else {
		// not a number
		value, err = resolveWire(wire, rules, wires)
		if err != nil {
			return 0, err
		}
	}
	return value, nil
}
