package day07

import (
	"fmt"
	"strings"
)

func countBeamOptions(grid []string, cache map[[2]int]int, startCoord [2]int) int {
	if startCoord[0] < 0 || startCoord[1] < 0 || startCoord[0] >= len(grid) || startCoord[1] >= len(grid[0]) {
		return 0
	}
	for row := startCoord[0]; row < len(grid)-1; row++ {
		if grid[row+1][startCoord[1]] == '^' {
			count, ok := cache[[2]int{row + 1, startCoord[1]}]
			if !ok {
				count = countBeamOptions(grid, cache, [2]int{row + 1, startCoord[1] - 1})
				count += countBeamOptions(grid, cache, [2]int{row + 1, startCoord[1] + 1})
				cache[[2]int{row + 1, startCoord[1]}] = count
			}
			return count
		}
	}
	return 1
}

func Part1(input []string) (string, error) {
	manifold := make([][]rune, len(input))
	splitCount := 0
	for row, line := range input {
		manifold[row] = make([]rune, len(line))
		for col, char := range line {
			if row == 0 {
				manifold[row][col] = char
			} else if manifold[row][col] == '|' {
				continue
			} else if char == '.' {
				if manifold[row-1][col] == '|' || manifold[row-1][col] == 'S' {
					manifold[row][col] = '|'
				} else {
					manifold[row][col] = '.'
				}
			} else if char == '^' {
				if manifold[row-1][col] == '|' || manifold[row-1][col] == 'S' {
					splitCount++
					if col-1 >= 0 {
						manifold[row][col-1] = '|'
					}
					if col+1 < len(line) {
						manifold[row][col+1] = '|'
					}
				}
				manifold[row][col] = '^'
			} else if char == '|' {
				manifold[row][col] = '|'
			}
		}
	}

	return fmt.Sprintf("%d", splitCount), nil
}

func Part2(input []string) (string, error) {
	cache := make(map[[2]int]int)
	count := countBeamOptions(input, cache, [2]int{0, strings.Index(input[0], "S")})
	return fmt.Sprintf("%d", count), nil
}
