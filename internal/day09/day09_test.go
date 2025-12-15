package day09

import (
	"io"
	"log/slog"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert := assert.New(t)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	input := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

	want := "50"
	got, err := Part1(logger, strings.Split(input, "\n"))
	assert.NoError(err)
	assert.Equal(want, got)
}

func TestLinesIntersect(t *testing.T) {
	tests := map[string]struct {
		hLine line
		vLine line
		exp   bool
	}{
		"Intersecting lines": {
			hLine: line{coord{row: 3, col: 2}, coord{row: 3, col: 7}},
			vLine: line{coord{row: 1, col: 5}, coord{row: 5, col: 5}},
			exp:   true,
		},
		"Non-intersecting lines": {
			hLine: line{coord{row: 3, col: 2}, coord{row: 3, col: 4}},
			vLine: line{coord{row: 1, col: 5}, coord{row: 5, col: 5}},
			exp:   false,
		},
		"Touching lines": {
			hLine: line{coord{row: 3, col: 2}, coord{row: 3, col: 5}},
			vLine: line{coord{row: 1, col: 5}, coord{row: 5, col: 5}},
			exp:   true,
		},
		"test case": {
			hLine: line{coord{5, 9}, coord{5, 2}},
			vLine: line{coord{8, 8}, coord{4, 8}},
			exp:   true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			got := linesIntersect(test.hLine, test.vLine)
			assert.Equal(test.exp, got)
		})
	}
}

func TestFindIntersectingLines(t *testing.T) {
	hLineMap := map[int][]line{
		1: {
			{coord{row: 1, col: 7}, coord{row: 1, col: 11}},
		},
		3: {
			{coord{row: 3, col: 2}, coord{row: 3, col: 7}},
		},
		5: {
			{coord{row: 5, col: 2}, coord{row: 5, col: 9}},
		},
		7: {
			{coord{row: 7, col: 9}, coord{row: 7, col: 11}},
		},
	}
	vLineMap := map[int][]line{
		2: {
			{coord{row: 3, col: 2}, coord{row: 5, col: 2}},
		},
		7: {
			{coord{row: 1, col: 7}, coord{row: 3, col: 7}},
		},
		9: {
			{coord{row: 5, col: 9}, coord{row: 7, col: 9}},
		},
		11: {
			{coord{row: 1, col: 11}, coord{row: 7, col: 11}},
		},
	}
	tests := map[string]struct {
		line line
		exp  bool
	}{
		"Intersecting lines": {
			line: line{coord{row: 3, col: 2}, coord{row: 3, col: 11}},
			exp:  true,
		},
		"Non-intersecting lines": {
			line: line{coord{row: 2, col: 2}, coord{row: 2, col: 4}},
			exp:  false,
		},
		"Touching lines": {
			line: line{coord{row: 3, col: 2}, coord{row: 3, col: 5}},
			exp:  true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			got := findIntersectingLines(test.line, hLineMap, vLineMap)
			assert.Equal(test.exp, got)
		})
	}
}

func TestIsForwards(t *testing.T) {
	tests := map[string]struct {
		l        line
		expected bool
	}{
		"horizontal line heading right is going forward": {
			l:        line{coord{5, 2}, coord{5, 6}},
			expected: true,
		},
		"horizontal line heading left is not going forward": {
			l:        line{coord{4, 12}, coord{4, 3}},
			expected: false,
		},
		"vertical line heading down is going forward": {
			l:        line{coord{3, 10}, coord{7, 10}},
			expected: true,
		},
		"vertical line heading up is not going forward": {
			l:        line{coord{8, 10}, coord{2, 10}},
			expected: false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			got := isForwards(test.l)
			assert.Equal(test.expected, got)
		})
	}
}

func TestIsHorizontal(t *testing.T) {
	tests := map[string]struct {
		l        line
		expected bool
	}{
		"horizontal line heading right is horizontal": {
			l:        line{coord{5, 2}, coord{5, 6}},
			expected: true,
		},
		"horizontal line heading left is horizontal": {
			l:        line{coord{4, 12}, coord{4, 3}},
			expected: true,
		},
		"vertical line heading down is not horizontal": {
			l:        line{coord{3, 10}, coord{7, 10}},
			expected: false,
		},
		"vertical line heading up is not horizontal": {
			l:        line{coord{8, 10}, coord{2, 10}},
			expected: false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			got := isHorizontal(test.l)
			assert.Equal(test.expected, got)
		})
	}
}

func TestLineHeading(t *testing.T) {
	tests := map[string]struct {
		l        line
		expected coord
	}{
		"horizontal left": {
			l:        line{coord{5, 2}, coord{5, 6}},
			expected: coord{0, 1},
		},
		"horizontal right": {
			l:        line{coord{3, 10}, coord{3, 6}},
			expected: coord{0, -1},
		},
		"vertical down": {
			l:        line{coord{0, 1}, coord{10, 1}},
			expected: coord{1, 0},
		},
		"vertical up": {
			l:        line{coord{7, 8}, coord{1, 8}},
			expected: coord{-1, 0},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			got := lineHeading(test.l)
			assert.Equal(test.expected, got)
		})
	}
}

func TestSortedLine(t *testing.T) {
	tests := map[string]struct {
		l        line
		expected line
	}{
		"forwards horizontal line does not change": {
			l:        line{coord{1, 6}, coord{1, 8}},
			expected: line{coord{1, 6}, coord{1, 8}},
		},
		"backwards horizontal line gets reversed": {
			l:        line{coord{1, 8}, coord{1, 6}},
			expected: line{coord{1, 6}, coord{1, 8}},
		},
		"downwards vertical line does not change": {
			l:        line{coord{6, 1}, coord{8, 1}},
			expected: line{coord{6, 1}, coord{8, 1}},
		},
		"upwards vertical line gets reversed": {
			l:        line{coord{8, 1}, coord{6, 1}},
			expected: line{coord{6, 1}, coord{8, 1}},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			got := sortedLine(test.l)
			assert.Equal(test.expected, got)
		})
	}
}

func TestCreateBorderLines(t *testing.T) {
	tests := map[string]struct {
		coords []coord
		lines  []line
	}{
		"AoC sample input": {
			coords: []coord{{1, 7}, {1, 11}, {7, 11}, {7, 9}, {5, 9}, {5, 2}, {3, 2}, {3, 7}},
			lines: []line{
				{coord{0, 6}, coord{0, 12}},
				{coord{0, 12}, coord{8, 12}},
				{coord{8, 12}, coord{8, 8}},
				{coord{8, 8}, coord{6, 8}},
				{coord{6, 8}, coord{6, 1}},
				{coord{6, 1}, coord{2, 1}},
				{coord{2, 1}, coord{2, 6}},
				{coord{2, 6}, coord{0, 6}}},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			got := createBorderLines(test.coords)
			assert.Equal(test.lines, got)
		})
	}
}

func TestGetLinesForBox(t *testing.T) {
	tests := map[string]struct {
		x        coord
		y        coord
		expected []line
	}{
		"1": {
			x: coord{0, 6},
			y: coord{2, 6},
			expected: []line{
				{coord{0, 6}, coord{2, 6}},
			},
		},
		"2": {
			x: coord{0, 6},
			y: coord{2, 9},
			expected: []line{
				{coord{0, 6}, coord{0, 9}},
				{coord{2, 6}, coord{2, 9}},
				{coord{0, 6}, coord{2, 6}},
				{coord{0, 9}, coord{2, 9}},
			},
		},
		"3": {
			x: coord{0, 6},
			y: coord{4, 5},
			expected: []line{
				{coord{0, 5}, coord{0, 6}},
				{coord{4, 5}, coord{4, 6}},
				{coord{0, 5}, coord{4, 5}},
				{coord{0, 6}, coord{4, 6}},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			got := getLinesForBox(test.x, test.y)
			assert.ElementsMatch(test.expected, got)
		})
	}

}

func TestPart2(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	tests := map[string]struct {
		input string
		want  string
	}{
		"1": {
			input: `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`,
			want: "24",
		},
		"2": {
			input: `6,0
6,2
9,2
9,5
5,5
5,4
2,4
2,8
10,8
10,10
11,10
11,0`,
			want: "32",
		},
	}
	for name, test := range tests {
		if name == "1" {
			continue
		}
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			got, err := Part2(logger, strings.Split(test.input, "\n"))
			assert.NoError(err)
			assert.Equal(test.want, got)
		})
	}
}
