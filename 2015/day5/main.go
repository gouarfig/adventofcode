package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gouarfig/adventofcode/helpers"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "async" {
		runAsync()
	} else {
		run()
	}
}

func run() {
	fmt.Print("reading from one big buffer\n")
	lines := helpers.GetLines("input.txt")
	count := 0
	nice1 := 0
	nice2 := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		count++
		if isNice(string(line), rulesPart1) {
			nice1++
		}
		if isNice(string(line), rulesPart2) {
			nice2++
		}
	}
	fmt.Printf("read %d lines:\n1. %d are nice\n2. %d are nice\n", count, nice1, nice2)
}

func runAsync() {
	fmt.Print("reading from a channel\n")
	lineChan := make(chan string)
	go helpers.ReadLines("input.txt", lineChan)
	count := 0
	nice1 := 0
	nice2 := 0
	for line := range lineChan {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		count++
		if isNice(string(line), rulesPart1) {
			nice1++
		}
		if isNice(string(line), rulesPart2) {
			nice2++
		}
	}
	fmt.Printf("read %d lines:\n1. %d are nice\n2. %d are nice\n", count, nice1, nice2)
}
