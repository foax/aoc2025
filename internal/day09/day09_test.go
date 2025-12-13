package day09

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert := assert.New(t)
	input := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

	want := "50"
	got, err := Part1(strings.Split(input, "\n"))
	assert.NoError(err)
	assert.Equal(want, got)
}

func TestPart2(t *testing.T) {
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
			input: `1,0
6,0
6,2
9,2
9,5
5,5
5,3
2,3
2,8
10,8
10,10
1,10`,
			want: "30",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			got, err := Part2(strings.Split(test.input, "\n"))
			assert.NoError(err)
			assert.Equal(test.want, got)
		})
	}
}
