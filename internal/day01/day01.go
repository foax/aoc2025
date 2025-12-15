package day01

import (
	"fmt"
	"log/slog"
	"strconv"
)

func Part1(logger *slog.Logger, input []string) (string, error) {
	var dialPos int = 50
	zeroCount := 0
	for _, line := range input {
		turns, _ := strconv.Atoi(line[1:])
		dialPos, _ = turnDial(dialPos, rune(line[0]), turns, 100)
		if dialPos == 0 {
			zeroCount++
		}
	}
	return fmt.Sprintf("%d", zeroCount), nil
}

func Part2(logger *slog.Logger, input []string) (string, error) {
	logger = logger.With("part", 2)
	var dialPos int = 50
	zeroCount := 0
	var zeroes int
	logger.Debug("init", "dialPos", dialPos, "zeroCount", zeroCount)
	for _, line := range input {
		turns, _ := strconv.Atoi(line[1:])
		dialPos, zeroes = turnDial(dialPos, rune(line[0]), turns, 100)
		zeroCount += zeroes
		logger.Debug("turn", "direction", string(rune(line[0])), "turns", turns, "dialPos", dialPos, "zeroes", zeroes, "zeroCount", zeroCount)
	}
	return fmt.Sprintf("%d", zeroCount), nil
}

func turnDial(start int, direction rune, clicks int, dialSize int) (int, int) {
	var mult int
	switch direction {
	case 'L':
		mult = -1
	case 'R':
		mult = 1
	}

	totalClicks := start + mult*clicks
	zeroCount := totalClicks / dialSize
	newDial := totalClicks % dialSize

	if zeroCount < 0 {
		zeroCount = -zeroCount
		if start != 0 {
			zeroCount++
		}
	} else if zeroCount == 0 && newDial <= 0 && start != 0 {
		zeroCount = 1
	}

	if newDial < 0 {
		newDial = dialSize + newDial
	}
	return newDial, zeroCount
}
