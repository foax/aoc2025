package day07

import "fmt"

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

	for _, row := range manifold {
		fmt.Println(string(row))
	}
	return fmt.Sprintf("%d", splitCount), nil
}

func Part2(input []string) (string, error) {
	return "unimplemented", nil
}
