package day06

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input []string) (string, error) {
	problems := make([][]int, len(input)-1)
	operators := make([]rune, len(strings.Fields(input[0])))
	totals := make([]int, len(input[0]))

	for i, line := range input {
		fields := strings.Fields(line)
		if i == len(input)-1 {
			for j, op := range fields {
				operators[j] = rune(op[0])
			}
		} else {
			problems[i] = make([]int, len(fields))
			for j, numStr := range fields {
				problems[i][j], _ = strconv.Atoi(numStr)
			}
		}
	}

	for i := range operators {
		for j := range problems {
			x := problems[j][i]
			if j == 0 {
				totals[i] = x
			} else {
				switch operators[i] {
				case '+':
					totals[i] += x
				case '*':
					totals[i] *= x
				}
			}
		}
	}

	total := 0
	for _, x := range totals {
		total += x
	}

	return fmt.Sprintf("%d", total), nil
}

func Part2(input []string) (string, error) {
	type operator struct {
		op  rune
		col int
	}
	operators := []operator{}
	for col, x := range input[len(input)-1] {
		if x == '*' || x == '+' {
			operators = append(operators, operator{op: x, col: col})
		}
	}

	totals := make([]int, len(operators))
	for idx, op := range operators {
		var startPos int
		if idx == len(operators)-1 {
			startPos = len(input[len(input)-1]) - 1
		} else {
			startPos = operators[idx+1].col - 2
		}
		for col := startPos; col >= op.col; col-- {
			rowNum := 0
			for row := range input[0 : len(input)-1] {
				if input[row][col] == ' ' {
					continue
				}
				rowNum = rowNum*10 + int(input[row][col]-'0')
			}
			if totals[idx] == 0 {
				totals[idx] = rowNum
			} else if op.op == '+' {
				totals[idx] += rowNum
			} else {
				totals[idx] *= rowNum
			}
		}
	}

	total := 0
	for _, x := range totals {
		total += x
	}

	return fmt.Sprintf("%d", total), nil
}
