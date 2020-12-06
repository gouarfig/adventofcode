package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloors(t *testing.T) {
	testData := []struct {
		instructions string
		floor        int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, testRun := range testData {
		assert.Equal(t, testRun.floor, getFloor([]byte(testRun.instructions)))
	}
}

func TestBasementInstruction(t *testing.T) {
	testData := []struct {
		instructions string
		position     int
	}{
		{")", 1},
		{"()())", 5},
	}

	for _, testRun := range testData {
		assert.Equal(t, testRun.position, getBasementInstruction([]byte(testRun.instructions)))
	}
}
