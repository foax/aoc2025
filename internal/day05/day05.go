package day05

import (
	"fmt"
	"log/slog"
	"sort"
	"strconv"
	"strings"
)

func Part1(input []string) (string, error) {
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
		slog.Debug("part 1 id loop", "idIdx", idIdx, "rangeIdx", rangeIdx)
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

func Part2(input []string) (string, error) {
	return "unimplemented", nil
}
