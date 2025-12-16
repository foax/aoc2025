package day10

import (
	"fmt"
	"log/slog"
	"math"
	"math/bits"
	"strconv"
	"strings"
)

type schematic struct {
	lights   indicatorLights
	buttons  []int
	joltages []int
}

type indicatorLights struct {
	lights     int
	lightCount int
}

func (l indicatorLights) String() string {
	lightStr := make([]rune, l.lightCount)
	for i := range lightStr {
		mask := 1 << i
		idx := l.lightCount - 1 - i
		if l.lights&mask == 0 {
			lightStr[idx] = '.'
		} else {
			lightStr[idx] = '#'
		}
	}
	return "{ [" + string(lightStr) + "] }"
}

func (s schematic) String() string {
	return fmt.Sprintf("{%v %v %v}", s.lights, s.buttons, s.joltages)
}

func bitIndexes(x int) []int {
	var bitList []int
	for x != 0 {
		idx := bits.TrailingZeros(uint(x))
		bitList = append(bitList, idx)
		x &= x - 1
	}
	return bitList
}

func parseInputLine(line string) schematic {
	s := schematic{}
	for _, f := range strings.Fields(line) {
		switch f[0] {
		case '[':
			for i := 1; i < len(f)-1; i++ {
				s.lights.lights <<= 1
				if f[i] == '#' {
					s.lights.lights++
				}
			}
			s.lights.lightCount = len(f) - 2
		case '(':
			button := 0
			buttonStrs := strings.Split(f[1:len(f)-1], ",")
			for _, bStr := range buttonStrs {
				b, _ := strconv.Atoi(bStr)
				button |= 1 << (s.lights.lightCount - b - 1)
			}
			s.buttons = append(s.buttons, button)
		case '{':
			joltageStrs := strings.Split(f[1:len(f)-1], ",")
			s.joltages = make([]int, len(joltageStrs))
			for idx, jStr := range joltageStrs {
				j, _ := strconv.Atoi(jStr)
				s.joltages[idx] = j
			}
		}
	}
	return s
}

func Part1(logger *slog.Logger, input []string) (string, error) {
	logger = logger.With("part", 1)
	schematics := make([]schematic, len(input))
	for idx, line := range input {
		schematics[idx] = parseInputLine(line)
	}

	total := 0
	for _, s := range schematics {
		min := math.MaxInt
		// x goes from 0b1 to 0b11111 (where there are 5 button presses to compare)
		for x := 1; x < 1<<len(s.buttons); x++ {
			indexes := bitIndexes(x)
			if len(indexes) > min {
				continue
			}

			y := 0
			for _, bIdx := range indexes {
				y ^= s.buttons[bIdx]
			}
			if y == s.lights.lights {
				min = len(indexes)
			}
		}
		total += min
	}
	return fmt.Sprintf("%d", total), nil
}

func Part2(logger *slog.Logger, input []string) (string, error) {
	return "unimplemented", nil
}
