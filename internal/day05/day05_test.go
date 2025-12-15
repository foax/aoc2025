package day05

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
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	want := "3"
	got, err := Part1(logger, strings.Split(input, "\n"))
	assert.NoError(err)
	assert.Equal(want, got)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	want := "14"
	got, err := Part2(logger, strings.Split(input, "\n"))
	assert.NoError(err)
	assert.Equal(want, got)
}
