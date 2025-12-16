package day10

import (
	"io"
	"log/slog"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndicatorLightsString(t *testing.T) {
	tests := map[string]struct {
		l        indicatorLights
		expected string
	}{
		"[.##.]": {
			l:        indicatorLights{0b0110, 4},
			expected: "[.##.]",
		},
		"[...#.]": {
			l:        indicatorLights{0b00010, 5},
			expected: "[...#.]",
		},
		"[##.#.#]": {
			l:        indicatorLights{0b110101, 6},
			expected: "[##.#.#]",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			got := test.l.String()
			assert.Equal(test.expected, got)
		})
	}
}

func TestBitIndexes(t *testing.T) {
	tests := map[string]struct {
		x        int
		expected []int
	}{
		"0b0110": {
			x:        0b0110,
			expected: []int{1, 2},
		},
		"0b10110100": {
			x:        0b10110100,
			expected: []int{2, 4, 5, 7},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			got := bitIndexes(test.x)
			assert.Equal(test.expected, got)
		})
	}
}

func TestPart1(t *testing.T) {
	assert := assert.New(t)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	input := `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`

	want := "7"
	got, err := Part1(logger, strings.Split(input, "\n"))
	assert.NoError(err)
	assert.Equal(want, got)
}
