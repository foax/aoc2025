package day02

import (
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

func TestPart1(t *testing.T) {
	assert := assert.New(t)
	input := []string{"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"}
	want := "1227775554"
	got, err := Part1(input)
	assert.NoError(err)
	assert.Equal(want, got)
}
