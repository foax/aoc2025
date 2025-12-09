package day04

import "fmt"

func checkNeighbours(grid [][]rune, x, y int) (total int) {
	height := len(grid)
	width := len(grid[0])
	for _, b := range []int{-1, 0, 1} {
		for _, a := range []int{-1, 0, 1} {
			if x+a < 0 || x+a >= width || y+b < 0 || y+b >= height || a == 0 && b == 0 {
				continue
			}
			if grid[y+b][x+a] == '@' {
				total++
			}
		}
	}
	return
}

func Part1(input []string) (string, error) {
	total := 0
	grid := make([][]rune, len(input))
	for y, line := range input {
		grid[y] = make([]rune, len(line))
		for x, roll := range line {
			grid[y][x] = roll
		}
	}
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != '@' {
				continue
			}
			n := checkNeighbours(grid, x, y)
			if n < 4 {
				total++
			}
		}
	}
	return fmt.Sprintf("%d", total), nil
}

func Part2(input []string) (string, error) {
	total := 0
	grid := make([][]rune, len(input))
	for y, line := range input {
		grid[y] = make([]rune, len(line))
		for x, roll := range line {
			grid[y][x] = roll
		}
	}

	keepGoing := true
	for keepGoing {
		removeCoords := map[string]bool{}
		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] != '@' {
					continue
				}
				n := checkNeighbours(grid, x, y)
				if n < 4 {
					total++
					removeCoords[fmt.Sprintf("%d,%d", x, y)] = true
				}
			}
		}
		if len(removeCoords) == 0 {
			keepGoing = false
			continue
		}
		for y := range grid {
			for x := range grid[y] {
				if removeCoords[fmt.Sprintf("%d,%d", x, y)] {
					grid[y][x] = '.'
				}
			}
		}
	}
	return fmt.Sprintf("%d", total), nil
}
