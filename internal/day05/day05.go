package day05

import (
	"fmt"
	"log/slog"
	"sort"
	"strconv"
	"strings"
)

func inRange(x int, r [2]int) bool {
	return x >= r[0] && x <= r[1]
}

func Part1(logger *slog.Logger, input []string) (string, error) {
	logger = logger.With("part", 1)
	var ids []int
	var ranges [][2]int

	parseRanges := true
	for _, line := range input {
		if parseRanges {
			if line == "" {
				parseRanges = false
				continue
			}
			rangeInts := [2]int{}
			for idx, x := range strings.Split(line, "-") {
				y, _ := strconv.Atoi(x)
				rangeInts[idx] = y
			}
			ranges = append(ranges, rangeInts)
		} else {
			y, _ := strconv.Atoi(line)
			ids = append(ids, y)
		}
	}

	sort.Ints(ids)
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		} else {
			return ranges[i][0] < ranges[j][0]
		}
	})

	rangeIdx := 0
	idIdx := 0
	freshTotal := 0
	for idIdx < len(ids) {
		logger.Debug("id loop", "idIdx", idIdx, "rangeIdx", rangeIdx)
		if ids[idIdx] < ranges[rangeIdx][0] {
			idIdx++
			continue
		}
		if ids[idIdx] <= ranges[rangeIdx][1] {
			freshTotal++
			idIdx++
			continue
		}
		rangeIdx++
		if rangeIdx == len(ranges) {
			break
		}
	}

	return fmt.Sprintf("%d", freshTotal), nil
}

func Part2(logger *slog.Logger, input []string) (string, error) {
	logger = logger.With("part", 2)
	var ranges [][2]int
	for _, line := range input {
		if line == "" {
			break
		}
		rangeInts := [2]int{}
		for idx, x := range strings.Split(line, "-") {
			y, _ := strconv.Atoi(x)
			rangeInts[idx] = y
		}
		ranges = append(ranges, rangeInts)
	}

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		} else {
			return ranges[i][0] < ranges[j][0]
		}
	})

	var consolidatedRanges [][2]int = [][2]int{ranges[0]}
	for _, origRange := range ranges {
		lastIdx := len(consolidatedRanges) - 1
		if inRange(origRange[0], consolidatedRanges[lastIdx]) {
			if !inRange(origRange[1], consolidatedRanges[lastIdx]) {
				consolidatedRanges[lastIdx][1] = origRange[1]
			}
		} else {
			consolidatedRanges = append(consolidatedRanges, origRange)
		}
	}

	total := 0
	for _, r := range consolidatedRanges {
		total += r[1] - r[0] + 1
	}
	return fmt.Sprintf("%d", total), nil
}
