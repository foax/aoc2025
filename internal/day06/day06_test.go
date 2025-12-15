package day06

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
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	want := "4277556"
	got, err := Part1(logger, strings.Split(input, "\n"))
	assert.NoError(err)
	assert.Equal(want, got)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	want := "3263827"
	got, err := Part2(logger, strings.Split(input, "\n"))
	assert.NoError(err)
	assert.Equal(want, got)
}
