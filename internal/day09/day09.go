package day09

import (
	"fmt"
	"log/slog"
	"sort"
	"strconv"
	"strings"
)

type coord struct {
	row int
	col int
}

type line struct {
	start coord
	end   coord
}

func stringToCoord(s string) (c coord) {
	strSlice := strings.Split(s, ",")
	col, _ := strconv.Atoi(strSlice[0])
	row, _ := strconv.Atoi(strSlice[1])
	c.col = col
	c.row = row
	return
}

func parseCoords(input []string) (coords []coord) {
	coords = make([]coord, len(input))
	for i, c := range input {
		coords[i] = stringToCoord(c)
	}
	return
}

func width(i, j coord) int {
	w := i.col - j.col
	if w < 0 {
		w = -w
	}
	w += 1
	return w
}

func height(i, j coord) int {
	h := i.row - j.row
	if h < 0 {
		h = -h
	}
	h += 1
	return h
}

func area(i, j coord) int {
	return height(i, j) * width(i, j)
}

func getLinesForBox(x, y coord) [4]line {
	var minRow, minCol, maxRow, maxCol int
	minRow = x.row
	maxRow = x.row
	minCol = x.col
	maxCol = x.col
	if y.row < minRow {
		minRow = y.row
	}
	if y.row > maxRow {
		maxRow = y.row
	}
	if y.col < minCol {
		minCol = y.col
	}
	if y.col > maxCol {
		maxCol = y.col
	}
	lines := [4]line{
		{coord{minRow, minCol}, coord{minRow, maxCol}},
		{coord{maxRow, minCol}, coord{maxRow, maxCol}},
		{coord{minRow, minCol}, coord{maxRow, minCol}},
		{coord{minRow, maxCol}, coord{maxRow, maxCol}},
	}
	return lines
}

func sortedLine(l line) line {
	a := l
	if l.start.row == l.end.row {
		if l.start.col > l.end.col {
			a.start = l.end
			a.end = l.start
		}
	} else if l.start.row > l.end.row {
		a.start = l.end
		a.end = l.start
	}
	return a
}

func linesIntersect(x, y line) bool {
	a := sortedLine(x)
	b := sortedLine(y)
	if a.start.row == a.end.row {
		// a is horizontal
		if b.start.col == b.end.col {
			// b is vertical
			if b.start.col >= a.start.col && b.start.col <= a.end.col &&
				a.start.row >= b.start.row && a.start.row <= b.end.row {
				return true
			}
		}
	} else {
		// a is vertical
		if b.start.row == b.end.row {
			// y is horizontal
			if a.start.col >= b.start.col && a.start.col <= b.end.col &&
				b.start.row >= a.start.row && b.start.row <= a.end.row {
				return true
			}
		}
	}
	return false
}

func findIntersectingLines(x line, hLineMap map[int][]line, vLineMap map[int][]line) bool {
	if x.start.row == x.end.row {
		// x is horizontal
		for col := x.start.col; col <= x.end.col; col++ {
			if vLines, ok := vLineMap[col]; ok {
				for _, v := range vLines {
					if v.start.row <= x.start.row {
						if v.end.row >= x.start.row {
							return true
						}
					} else {
						continue
					}
				}
			}
		}
	} else {
		for row := x.start.row; row <= x.end.row; row++ {
			if hLines, ok := hLineMap[row]; ok {
				for _, h := range hLines {
					if h.start.col <= x.start.col {
						if h.end.col >= x.start.col {
							return true
						}
					} else {
						continue
					}
				}
			}
		}
	}
	return false
}

func Part1(input []string) (string, error) {
	coords := parseCoords(input)
	maxArea := 0
	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			area := area(coords[i], coords[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return fmt.Sprintf("%d", maxArea), nil
}

func lineHeading(l line) coord {
	if l.start.row == l.end.row {
		if l.start.col < l.end.col {
			return coord{0, 1}
		} else {
			return coord{0, -1}
		}
	} else {
		if l.start.row < l.end.row {
			return coord{1, 0}
		} else {
			return coord{-1, 0}
		}
	}
}

func isHorizontal(l line) bool {
	return lineHeading(l).row == 0
}

func isForwards(l line) bool {
	h := lineHeading(l)
	var x int
	if h.row == 0 {
		x = h.col
	} else {
		x = h.row
	}
	return x == 1
}

func createBorderLines(coords []coord) []line {
	lines := []line{}
	minCoord := coords[0]
	minIdx := 0
	for i, c := range coords {
		if c.row <= minCoord.row && c.col < minCoord.col {
			minCoord = c
			minIdx = i
		}
	}

	start := coord{minCoord.row - 1, minCoord.col - 1}
	for idx := minIdx; len(lines) < len(coords); idx++ {
		i := idx % len(coords)
		j := (idx + 1) % len(coords)
		k := (idx + 2) % len(coords)
		edgeLine := line{coords[i], coords[j]}
		nextEdgeLine := line{coords[j], coords[k]}
		var newLineShort, newLineLong line
		if len(lines) > 0 {
			newLineShort.start = lines[len(lines)-1].end
		} else {
			newLineShort.start = start
		}
		newLineLong.start = newLineShort.start
		if isHorizontal(edgeLine) {
			newLineShort.end = coord{row: newLineShort.start.row}
			newLineLong.end = newLineShort.end
			if isForwards(edgeLine) {
				newLineShort.end.col = coords[j].col - 1
				newLineLong.end.col = coords[j].col + 1
			} else {
				newLineShort.end.col = coords[j].col + 1
				newLineLong.end.col = coords[j].col - 1
			}
		} else {
			newLineShort.end = coord{col: newLineShort.start.col}
			newLineLong.end = newLineShort.end
			if isForwards(edgeLine) {
				newLineShort.end.row = coords[j].row - 1
				newLineLong.end.row = coords[j].row + 1
			} else {
				newLineShort.end.row = coords[j].row + 1
				newLineLong.end.row = coords[j].row - 1
			}
		}
		if linesIntersect(newLineLong, nextEdgeLine) {
			lines = append(lines, newLineShort)
		} else {
			lines = append(lines, newLineLong)
		}
	}
	return lines
}

func Part2(input []string) (string, error) {
	coords := parseCoords(input)
	borderLines := createBorderLines(coords)
	slog.Debug("part 2", "msg", "borderLines created", "len", len(borderLines))

	horizontalLines := make(map[int][]line)
	verticalLines := make(map[int][]line)
	for _, l := range borderLines {
		l = sortedLine(l)
		if isHorizontal(l) {
			if _, ok := horizontalLines[l.start.row]; !ok {
				horizontalLines[l.start.row] = []line{}
			}
			horizontalLines[l.start.row] = append(horizontalLines[l.start.row], l)
		} else {
			if _, ok := verticalLines[l.start.col]; !ok {
				verticalLines[l.start.col] = []line{}
			}
			verticalLines[l.start.col] = append(verticalLines[l.start.col], l)
		}
	}
	slog.Debug("part 2", "msg", "horizontal and vertical lines grouped")

	for idx := range horizontalLines {
		sort.Slice(horizontalLines[idx], func(i, j int) bool {
			return horizontalLines[idx][i].start.col < horizontalLines[idx][j].start.col
		})
	}
	slog.Debug("part 2", "msg", "horizontalLines sorted")

	for idx := range verticalLines {
		sort.Slice(verticalLines[idx], func(i, j int) bool {
			return verticalLines[idx][i].start.row < verticalLines[idx][j].start.row
		})
	}
	slog.Debug("part 2", "msg", "verticalLines sorted")

	maxArea := 0
	for i := range coords {
		if i == len(coords)-1 {
			break
		}
		for j := i + 1; j < len(coords); j++ {
			lines := getLinesForBox(coords[i], coords[j])
			intersecting := false
			for _, l := range lines {
				if findIntersectingLines(l, horizontalLines, verticalLines) {
					intersecting = true
					break
				}
			}
			if !intersecting {
				area := area(coords[i], coords[j])
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return fmt.Sprintf("%d", maxArea), nil
}
