package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNicePart1(t *testing.T) {
	testData := []struct {
		value string
		nice  bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, testRun := range testData {
		t.Run(testRun.value, func(t *testing.T) {
			assert.Equal(t, testRun.nice, isNice(testRun.value, rulesPart1))
		})
	}
}

func TestNicePart2(t *testing.T) {
	testData := []struct {
		value string
		nice  bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	for _, testRun := range testData {
		t.Run(testRun.value, func(t *testing.T) {
			assert.Equal(t, testRun.nice, isNice(testRun.value, rulesPart2))
		})
	}
}
