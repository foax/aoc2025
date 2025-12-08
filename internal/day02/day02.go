package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func numDigits(x int) int {
	if x == 0 {
		return 1
	}
	count := 0
	for x > 0 {
		x /= 10
		count++
	}
	return count
}

func pow10(x int) int {
	pow := 1
	for i := 0; i < x; i++ {
		pow *= 10
	}
	return pow
}

func nextInvalidId(cur int) int {
	id := cur
	digits := numDigits(id)
	if digits%2 != 0 {
		id = pow10(digits)
	}
	idStr := strconv.Itoa(id)
	idHalfStr := idStr[0 : len(idStr)/2]
	newId, _ := strconv.Atoi(idHalfStr + idHalfStr)
	if newId > cur {
		return newId
	}
	idHalfInt, _ := strconv.Atoi(idHalfStr)
	idHalfStr = strconv.Itoa(idHalfInt + 1)
	newId, _ = strconv.Atoi(idHalfStr + idHalfStr)
	return newId
}

func isInvalidId(cur int) bool {
	curStr := strconv.Itoa(cur)
	return curStr[0:len(curStr)/2] == curStr[len(curStr)/2:]
}

func Part1(input []string) (string, error) {
	var invalidIds []int
	for _, r := range strings.Split(input[0], ",") {
		idStrs := strings.Split(r, "-")
		low, _ := strconv.Atoi(idStrs[0])
		high, _ := strconv.Atoi(idStrs[1])
		for x := low; x <= high; x = nextInvalidId(x) {
			if isInvalidId(x) {
				invalidIds = append(invalidIds, x)
			}
		}
	}

	sum := 0
	for _, n := range invalidIds {
		sum += n
	}
	return fmt.Sprintf("%d", sum), nil
}

func Part2(input []string) (string, error) {
	return "unimplemented", nil
}
