package day08

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Don't use math.Sqrt so we avoid using floats
func distance(a, b [3]int) (x int) {
	for i := range a {
		y := a[i] - b[i]
		x += y * y
	}
	return
}

func Part1(input []string) (string, error) {
	return part1Inner(input, 1000)
}

func part1Inner(input []string, circuitConnect int) (string, error) {
	circuits := make(map[int]map[[3]int]struct{})
	coords := make([][3]int, len(input))
	coordCircuits := make(map[[3]int]int)
	for idx, line := range input {
		coord := strings.Split(line, ",")
		for x := range coord {
			c, _ := strconv.Atoi(coord[x])
			coords[idx][x] = c
		}
		circuits[idx] = make(map[[3]int]struct{})
		circuits[idx][coords[idx]] = struct{}{}
		coordCircuits[coords[idx]] = idx
	}

	distances := make(map[[2][3]int]int)
	var coordPairs [][2][3]int
	for x := range coords {
		if x == len(coords)-1 {
			break
		}
		for y := x + 1; y < len(coords); y++ {
			d := distance(coords[x], coords[y])
			coordSlice := [][3]int{coords[x], coords[y]}
			sort.Slice(coordSlice, func(i, j int) bool {
				if coordSlice[i][0] == coordSlice[j][0] {
					if coordSlice[i][1] == coordSlice[j][1] {
						return coordSlice[i][2] < coordSlice[j][2]
					}
					return coordSlice[i][1] < coordSlice[j][1]
				}
				return coordSlice[i][0] < coordSlice[j][0]
			})
			var coordPair [2][3]int
			copy(coordPair[:], coordSlice)
			distances[coordPair] = d
			coordPairs = append(coordPairs, coordPair)
		}
	}

	sort.Slice(coordPairs, func(i, j int) bool {
		return distances[coordPairs[i]] < distances[coordPairs[j]]
	})

	for i := range circuitConnect {
		circuitId := coordCircuits[coordPairs[i][0]]
		oldCircuitId := coordCircuits[coordPairs[i][1]]
		if circuitId != oldCircuitId {
			for coord := range circuits[oldCircuitId] {
				coordCircuits[coord] = circuitId
				circuits[circuitId][coord] = struct{}{}
			}
			delete(circuits, oldCircuitId)
		}
	}

	circuitSizes := make([]int, len(circuits))
	i := 0
	for _, c := range circuits {
		circuitSizes[i] = len(c)
		i++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(circuitSizes)))
	total := circuitSizes[0] * circuitSizes[1] * circuitSizes[2]

	return fmt.Sprintf("%d", total), nil
}

func Part2(input []string) (string, error) {
	circuits := make(map[int]map[[3]int]struct{})
	coords := make([][3]int, len(input))
	coordCircuits := make(map[[3]int]int)
	for idx, line := range input {
		coord := strings.Split(line, ",")
		for x := range coord {
			c, _ := strconv.Atoi(coord[x])
			coords[idx][x] = c
		}
		circuits[idx] = make(map[[3]int]struct{})
		circuits[idx][coords[idx]] = struct{}{}
		coordCircuits[coords[idx]] = idx
	}

	distances := make(map[[2][3]int]int)
	var coordPairs [][2][3]int
	for x := range coords {
		if x == len(coords)-1 {
			break
		}
		for y := x + 1; y < len(coords); y++ {
			d := distance(coords[x], coords[y])
			coordSlice := [][3]int{coords[x], coords[y]}
			sort.Slice(coordSlice, func(i, j int) bool {
				if coordSlice[i][0] == coordSlice[j][0] {
					if coordSlice[i][1] == coordSlice[j][1] {
						return coordSlice[i][2] < coordSlice[j][2]
					}
					return coordSlice[i][1] < coordSlice[j][1]
				}
				return coordSlice[i][0] < coordSlice[j][0]
			})
			var coordPair [2][3]int
			copy(coordPair[:], coordSlice)
			distances[coordPair] = d
			coordPairs = append(coordPairs, coordPair)
		}
	}

	sort.Slice(coordPairs, func(i, j int) bool {
		return distances[coordPairs[i]] < distances[coordPairs[j]]
	})

	finalPair := 0
	for i := range coordPairs {
		circuitId := coordCircuits[coordPairs[i][0]]
		oldCircuitId := coordCircuits[coordPairs[i][1]]
		if circuitId != oldCircuitId {
			for coord := range circuits[oldCircuitId] {
				coordCircuits[coord] = circuitId
				circuits[circuitId][coord] = struct{}{}
			}
			delete(circuits, oldCircuitId)
		}
		if len(circuits) == 1 {
			finalPair = i
			break
		}
	}

	return fmt.Sprintf("%d", coordPairs[finalPair][0][0]*coordPairs[finalPair][1][0]), nil
}
