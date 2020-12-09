package main

import (
	"strings"
)

type IsNice func(value string) bool

var (
	rulesPart1 = []IsNice{hasThreeVowels, hasDoubleLetter, noForbiddenPattern}
	rulesPart2 = []IsNice{hasPairOfTwoLetters, letterRepeats}
)

func isNice(value string, rules []IsNice) bool {
	for _, rule := range rules {
		if !rule(value) {
			return false
		}
	}
	return true
}

func hasThreeVowels(value string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	total := 0
	for _, vowel := range vowels {
		total += strings.Count(value, vowel)
		if total >= 3 {
			// fmt.Printf("%q has 3 (or more) vowels\n", value)
			return true
		}
	}
	return false
}

func hasDoubleLetter(value string) bool {
	for i := 0; i < len(value)-1; i++ {
		if value[i] == value[i+1] {
			// fmt.Printf("%q has at least one double letter\n", value)
			return true
		}
	}
	return false
}

func noForbiddenPattern(value string) bool {
	forbidden := []string{"ab", "cd", "pq", "xy"}
	for _, forbid := range forbidden {
		if strings.Contains(value, forbid) {
			return false
		}
	}
	// fmt.Printf("%q does not contain any forbidden string\n", value)
	return true
}

func hasPairOfTwoLetters(value string) bool {
	for i := 0; i < len(value)-1; i++ {
		if strings.Count(value, value[i:i+2]) > 1 {
			// fmt.Printf("%q has at least two instances of %q\n", value, value[i:i+2])
			return true
		}
	}
	return false
}

func letterRepeats(value string) bool {
	for i := 0; i < len(value)-2; i++ {
		if value[i] == value[i+2] {
			// fmt.Printf("%q has a letter that repeats %q\n", value, value[i:i+3])
			return true
		}
	}
	return false
}
