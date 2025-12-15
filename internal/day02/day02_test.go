package day02

import (
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextInvalidId(t *testing.T) {
	tests := map[string]struct {
		curId      int
		expectedId int
	}{
		"Next invalid ID is same number of digits and greater than current ID": {
			curId:      1000,
			expectedId: 1010,
		},
		"Next invalid ID is greater than current ID which is not invalid": {
			curId:      1020,
			expectedId: 1111,
		},
		"Current ID is an invalid ID, next invalid ID is same number of digits": {
			curId:      2424,
			expectedId: 2525,
		},
		"Current ID is an invalid ID, next invalid ID is more digits": {
			curId:      9999,
			expectedId: 100100,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			gotId := nextInvalidId(test.curId)
			assert.Equal(test.expectedId, gotId)
		})
	}
}

func TestIsInvalidId(t *testing.T) {
	tests := map[string]struct {
		id             int
		expectedResult bool
	}{
		"Single digit is not invalid": {
			id:             1,
			expectedResult: false,
		},
		"Two identical digits is invalid": {
			id:             11,
			expectedResult: true,
		},
		"Two sequential digits is invalid": {
			id:             34,
			expectedResult: false,
		},
		"Odd number of digits is not invalid": {
			id:             234,
			expectedResult: false,
		},
		"Even number of digits with matching sequence of digits is invalid": {
			id:             456456,
			expectedResult: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			result := isInvalidId(test.id)
			assert.Equal(test.expectedResult, result)
		})
	}

}

func TestIsInvalidIdPart2(t *testing.T) {
	tests := map[string]struct {
		id             int
		expectedResult bool
	}{
		"Single digit is not an invalid ID": {
			id:             1,
			expectedResult: false,
		},
		"Two identical digits is an invalid ID": {
			id:             11,
			expectedResult: true,
		},
		"Two sequential digits is not an invalid ID": {
			id:             34,
			expectedResult: false,
		},
		"Three non repeating digits is not an invalid ID": {
			id:             235,
			expectedResult: false,
		},
		"Three repeating digits is an invalid ID": {
			id:             666,
			expectedResult: true,
		},
		"10 digit number with two repeating sequences is an invalid ID": {
			id:             1143211432,
			expectedResult: true,
		},
		"9 digit number with three repeating sequences is an invalid ID": {
			id:             369369369,
			expectedResult: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			result := isInvalidIdPart2(test.id)
			assert.Equal(test.expectedResult, result)
		})
	}
}

func TestPart1(t *testing.T) {
	assert := assert.New(t)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	input := []string{"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"}
	want := "1227775554"
	got, err := Part1(logger, input)
	assert.NoError(err)
	assert.Equal(want, got)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	input := []string{"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"}
	want := "4174379265"
	got, err := Part2(logger, input)
	assert.NoError(err)
	assert.Equal(want, got)
}
