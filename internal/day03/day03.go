package day03

import (
	"log/slog"
	"strconv"
)

func getMaxJoltage(bank string, batteries int) (total int) {
	bankInt := make([]int, len(bank))
	max := make([]int, batteries)
	maxIdx := make([]int, batteries)
	for idx, j := range bank {
		bankInt[idx] = int(j - '0')
	}

	for b := 0; b < batteries; b++ {
		var startIdx int
		if b == 0 {
			startIdx = 0
		} else {
			startIdx = maxIdx[b-1] + 1
		}
		endIdx := len(bankInt) - batteries + b
		// b = 0, len = 15, batteries = 12: 15 - 12 + 0 = 3
		// b = 11, len = 15, batteries = 12: 15 - 12 + 11 = 14
		for idx := startIdx; idx <= endIdx; idx++ {
			if bankInt[idx] > max[b] {
				max[b] = bankInt[idx]
				maxIdx[b] = idx
			}
		}
	}

	for _, b := range max {
		total *= 10
		total += b
	}
	return
}

func Part1(logger *slog.Logger, input []string) (string, error) {
	total := 0
	for _, bank := range input {
		total += getMaxJoltage(bank, 2)
	}
	return strconv.Itoa(total), nil
}

func Part2(logger *slog.Logger, input []string) (string, error) {
	total := 0
	for _, bank := range input {
		total += getMaxJoltage(bank, 12)
	}
	return strconv.Itoa(total), nil
}
