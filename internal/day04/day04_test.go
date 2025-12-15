package day04

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
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	want := "13"
	got, err := Part1(logger, strings.Split(input, "\n"))
	assert.NoError(err)
	assert.Equal(want, got)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	want := "43"
	got, err := Part2(logger, strings.Split(input, "\n"))
	assert.NoError(err)
	assert.Equal(want, got)
}
