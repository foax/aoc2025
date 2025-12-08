package day03

import "strconv"

func getMaxJoltage(bank string) int {
	bankInt := make([]int, len(bank))
	max := []int{0, 0}
	maxIdx := 0
	for idx, j := range bank {
		bankInt[idx] = int(j - '0')
		if idx < len(bank)-1 && bankInt[idx] > max[0] {
			max[0] = bankInt[idx]
			maxIdx = idx
		}
	}

	for idx := maxIdx + 1; idx < len(bankInt); idx++ {
		if bankInt[idx] > max[1] {
			max[1] = bankInt[idx]
		}
	}
	return max[0]*10 + max[1]
}

func Part1(input []string) (string, error) {
	total := 0
	for _, bank := range input {
		total += getMaxJoltage(bank)
	}
	return strconv.Itoa(total), nil
}

func Part2(input []string) (string, error) {
	return "Unimplemented", nil
}
