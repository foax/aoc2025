package day03

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMaxJoltage(t *testing.T) {
	tests := map[string]int{
		"987654321111111": 98,
		"811111111111119": 89,
		"234234234234278": 78,
		"818181911112111": 92,
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			j := getMaxJoltage(name)
			assert.Equal(test, j)
		})
	}
}

func TestPart1(t *testing.T) {
	assert := assert.New(t)
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	want := "357"
	got, err := Part1(strings.Split(input, "\n"))
	assert.NoError(err)
	assert.Equal(want, got)
}
