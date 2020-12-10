package main

import (
	"regexp"
	"strconv"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/gouarfig/adventofcode/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculateString(t *testing.T) {
	pattern := regexp.MustCompile(`^\(string\) \(len=(\d+)\) (".*")\n$`)
	lines := helpers.GetLines("demo_input.txt")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		code, _, encoded := calculateString(string(line))
		dump := spew.Sdump(string(line))
		match := pattern.FindStringSubmatch(dump)
		require.Lenf(t, match, 3, "%s didn't match", dump)
		length, _ := strconv.Atoi(match[1])
		assert.Equal(t, length, code)
		assert.Equal(t, len(match[2]), encoded)
		t.Logf("%s -> %s", string(line), match[2])
	}
}
