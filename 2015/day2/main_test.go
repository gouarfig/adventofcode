package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrappingPaper(t *testing.T) {
	testData := []struct {
		dimensions []int
		need       int
	}{
		{[]int{2, 3, 4}, 58},
		{[]int{1, 1, 10}, 43},
	}

	for _, testRun := range testData {
		assert.Equal(t, testRun.need, wrappingPaper(testRun.dimensions))
	}
}

func TestExtractMin(t *testing.T) {
	testData := []struct {
		orig  []int
		after []int
		min   int
	}{
		{[]int{2, 3, 4}, []int{3, 4}, 2},
		{[]int{1, 1, 10}, []int{1, 10}, 1},
		{[]int{2, 1, 3}, []int{2, 3}, 1},
		{[]int{3, 2, 1}, []int{3, 2}, 1},
	}

	for _, testRun := range testData {
		after, min := extractMin(testRun.orig)
		assert.ElementsMatch(t, testRun.after, after)
		assert.Equal(t, testRun.min, min)
	}
}

func TestRibbonLength(t *testing.T) {
	testData := []struct {
		sides  []int
		length int
	}{
		{[]int{2, 3, 4}, 34},
		{[]int{1, 1, 10}, 14},
	}

	for _, testRun := range testData {
		assert.Equal(t, testRun.length, ribbonLength(testRun.sides))
	}
}
