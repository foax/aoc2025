package day03

import (
	"io"
	"log/slog"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMaxJoltage(t *testing.T) {
	tests := map[string]struct {
		batteries       string
		batteriesNeeded int
		maxJoltage      int
	}{
		"987654321111111 2":  {batteries: "987654321111111", batteriesNeeded: 2, maxJoltage: 98},
		"811111111111119 2":  {batteries: "811111111111119", batteriesNeeded: 2, maxJoltage: 89},
		"234234234234278 2":  {batteries: "234234234234278", batteriesNeeded: 2, maxJoltage: 78},
		"818181911112111 2":  {batteries: "818181911112111", batteriesNeeded: 2, maxJoltage: 92},
		"987654321111111 12": {batteries: "987654321111111", batteriesNeeded: 12, maxJoltage: 987654321111},
		"811111111111119 12": {batteries: "811111111111119", batteriesNeeded: 12, maxJoltage: 811111111119},
		"234234234234278 12": {batteries: "234234234234278", batteriesNeeded: 12, maxJoltage: 434234234278},
		"818181911112111 12": {batteries: "818181911112111", batteriesNeeded: 12, maxJoltage: 888911112111},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			j := getMaxJoltage(test.batteries, test.batteriesNeeded)
			assert.Equal(test.maxJoltage, j)
		})
	}
}

func TestPart1(t *testing.T) {
	assert := assert.New(t)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	want := "357"
	got, err := Part1(logger, strings.Split(input, "\n"))
	assert.NoError(err)
	assert.Equal(want, got)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	want := "3121910778619"
	got, err := Part2(logger, strings.Split(input, "\n"))
	assert.NoError(err)
	assert.Equal(want, got)
}
