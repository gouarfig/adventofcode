package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	example := []struct {
		start  string
		finish string
	}{
		{"1,0,0,0,99", "2,0,0,0,99"},
		{"2,3,0,3,99", "2,3,0,6,99"},
		{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
		{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
	}

	for _, fixture := range example {
		memory := loadMemory(fixture.start)
		runProgram(memory)
		output := saveMemory(memory)
		assert.Equal(t, fixture.finish, output)
	}
}
