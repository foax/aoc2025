package day01

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTurnDial(t *testing.T) {
	tests := map[string]struct {
		start        int
		direction    rune
		clicks       int
		dialSize     int
		expDialPos   int
		expZeroCount int
	}{
		"Right turn that does not past zero": {
			start:        50,
			direction:    'R',
			clicks:       5,
			dialSize:     100,
			expDialPos:   55,
			expZeroCount: 0,
		},
		"Left turn that does not past zero": {
			start:        50,
			direction:    'L',
			clicks:       5,
			dialSize:     100,
			expDialPos:   45,
			expZeroCount: 0,
		},
		"Right turn lands on zero": {
			start:        99,
			direction:    'R',
			clicks:       1,
			dialSize:     100,
			expDialPos:   0,
			expZeroCount: 1,
		},
		"Left turn lands on zero": {
			start:        1,
			direction:    'L',
			clicks:       1,
			dialSize:     100,
			expDialPos:   0,
			expZeroCount: 1,
		},
		"Right turn starts on zero": {
			start:        0,
			direction:    'R',
			clicks:       1,
			dialSize:     100,
			expDialPos:   1,
			expZeroCount: 0,
		},
		"Left turn starts on zero": {
			start:        0,
			direction:    'L',
			clicks:       1,
			dialSize:     100,
			expDialPos:   99,
			expZeroCount: 0,
		},
		"Right turn goes past zero": {
			start:        99,
			direction:    'R',
			clicks:       2,
			dialSize:     100,
			expDialPos:   1,
			expZeroCount: 1,
		},
		"Left turn goes past zero": {
			start:        1,
			direction:    'L',
			clicks:       2,
			dialSize:     100,
			expDialPos:   99,
			expZeroCount: 1,
		},
		"Right turn multiple revolutions": {
			start:        99,
			direction:    'R',
			clicks:       102,
			dialSize:     100,
			expDialPos:   1,
			expZeroCount: 2,
		},
		"Left turn multiple revolutions": {
			start:        1,
			direction:    'L',
			clicks:       102,
			dialSize:     100,
			expDialPos:   99,
			expZeroCount: 2,
		},
		"Right turn multiple revolutions lands on zero": {
			start:        99,
			direction:    'R',
			clicks:       201,
			dialSize:     100,
			expDialPos:   0,
			expZeroCount: 3,
		},
		"Left turn multiple revolutions lands on zero": {
			start:        1,
			direction:    'L',
			clicks:       201,
			dialSize:     100,
			expDialPos:   0,
			expZeroCount: 3,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			dialPos, zeroCount := turnDial(test.start, test.direction, test.clicks, test.dialSize)
			assert.Equal(test.expDialPos, dialPos, "Dial position should be the same")
			assert.Equal(test.expZeroCount, zeroCount, "Zero count should be the same")
		})
	}
}

func TestPart1(t *testing.T) {
	assert := assert.New(t)
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	want := "3"
	got, err := Part1(strings.Split(input, "\n"))
	assert.NoError(err)
	assert.Equal(want, got)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	want := "6"
	got, err := Part2(strings.Split(input, "\n"))
	assert.NoError(err)
	assert.Equal(want, got)
}
