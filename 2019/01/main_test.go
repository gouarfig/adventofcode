package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamplePart1(t *testing.T) {

	example := []struct {
		mass int64
		fuel int64
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, fixture := range example {
		assert.Equal(t, fixture.fuel, fuel(fixture.mass))
	}
}

func TestExamplePart2(t *testing.T) {

	example := []struct {
		mass int64
		fuel int64
	}{
		{12, 2},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, fixture := range example {
		assert.Equal(t, fixture.fuel, totalFuel(fixture.mass))
	}
}
