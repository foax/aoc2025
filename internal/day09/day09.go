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

// stringToCoord converts a string in the format of "x,y" to a coord struct
// coord{row: y, col: x}. Returns the converted coord
func stringToCoord(s string) (c coord) {
	strSlice := strings.Split(s, ",")
	col, _ := strconv.Atoi(strSlice[0])
	row, _ := strconv.Atoi(strSlice[1])
	c.col = col
	c.row = row
	return
}

// parseCoords parses a list of coordinate strings.
// Returns a slice of coord structs.
func parseCoords(input []string) (coords []coord) {
	coords = make([]coord, len(input))
	for i, c := range input {
		coords[i] = stringToCoord(c)
	}
	return
}

// width returns the width of the tiles between a and b
func width(a, b coord) int {
	w := a.col - b.col
	if w < 0 {
		w = -w
	}
	w += 1
	return w
}

// height returns the height of the tiles between i and j
func height(a, b coord) int {
	h := a.row - b.row
	if h < 0 {
		h = -h
	}
	h += 1
	return h
}

// area returns the area of the tiles in the rectangle with corners a and b
func area(a, b coord) int {
	return height(a, b) * width(a, b)
}

// getLinesForBox returns the lines that make up the border of the rectangle
// with corners a and b.
func getLinesForBox(a, b coord) []line {
	var minRow, minCol, maxRow, maxCol int
	minRow = a.row
	maxRow = a.row
	minCol = a.col
	maxCol = a.col
	if b.row < minRow {
		minRow = b.row
	}
	if b.row > maxRow {
		maxRow = b.row
	}
	if b.col < minCol {
		minCol = b.col
	}
	if b.col > maxCol {
		maxCol = b.col
	}

	linesSet := map[line]struct{}{}
	linesSet[line{coord{minRow, minCol}, coord{minRow, maxCol}}] = struct{}{}
	linesSet[line{coord{maxRow, minCol}, coord{maxRow, maxCol}}] = struct{}{}
	linesSet[line{coord{minRow, minCol}, coord{maxRow, minCol}}] = struct{}{}
	linesSet[line{coord{minRow, maxCol}, coord{maxRow, maxCol}}] = struct{}{}

	lines := []line{}
	for l := range linesSet {
		if l.start != l.end {
			lines = append(lines, l)
		}
	}
	return lines
}

// sortedLine sorts the coords in the line so that the upper left coord
// comes first in the list.
func sortedLine(l line) line {
	sorted := l
	if l.start.row == l.end.row {
		if l.start.col > l.end.col {
			sorted.start = l.end
			sorted.end = l.start
		}
	} else if l.start.row > l.end.row {
		sorted.start = l.end
		sorted.end = l.start
	}
	return sorted
}

// linesIntersect returns true if the line a and b intersect.
func linesIntersect(a, b line) bool {
	x := sortedLine(a)
	y := sortedLine(b)
	if x.start.row == x.end.row {
		// x is horizontal
		if y.start.col == y.end.col {
			// y is vertical
			if y.start.col >= x.start.col && y.start.col <= x.end.col &&
				x.start.row >= y.start.row && x.start.row <= y.end.row {
				return true
			}
		}
	} else {
		// x is vertical
		if y.start.row == y.end.row {
			// y is horizontal
			if x.start.col >= y.start.col && x.start.col <= y.end.col &&
				y.start.row >= x.start.row && y.start.row <= x.end.row {
				return true
			}
		}
	}
	return false
}

// findIntersectingLines check if l intersects with any horizontal or vertical lines
// in hLineMap and vLineMap. Returns true if l intersects.
func findIntersectingLines(l line, hLineMap map[int][]line, vLineMap map[int][]line) bool {
	if l.start.row == l.end.row {
		// l is horizontal
		for col := l.start.col; col <= l.end.col; col++ {
			if vLines, ok := vLineMap[col]; ok {
				for _, v := range vLines {
					if v.start.row <= l.start.row {
						if v.end.row >= l.start.row {
							return true
						}
					} else {
						continue
					}
				}
			}
		}
	} else {
		// l is vertical
		for row := l.start.row; row <= l.end.row; row++ {
			if hLines, ok := hLineMap[row]; ok {
				for _, h := range hLines {
					if h.start.col <= l.start.col {
						if h.end.col >= l.start.col {
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

// lineHeading returns a coord indicating the direction the line l is heading in
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

// isHorizontal returns if a line is horizontal
func isHorizontal(l line) bool {
	return lineHeading(l).row == 0
}

// isForwards returns if a line is heading in a positive direction
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

// createBorderLines creates a list of lines that forms a perimeter around
// the tile loop formed by linking all tiles in coords.
func createBorderLines(coords []coord) []line {
	lines := []line{}

	// find the upper-leftmost coord in the list as our starting point
	// so we know we're starting on the outside of the loop
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

		// create two lines, a short one and a long one. Either the next
		// edge will get in the way of the border, in which case it will
		// intersect with the long line. If it does, use the short line.
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

func Part1(logger *slog.Logger, input []string) (string, error) {
	logger = logger.With("part", 1)
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

func Part2(logger *slog.Logger, input []string) (string, error) {
	logger = logger.With("part", 2)
	coords := parseCoords(input)
	borderLines := createBorderLines(coords)
	logger.Debug("borderLines created", "len", len(borderLines))

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
	logger.Debug("horizontal and vertical lines grouped")

	for idx := range horizontalLines {
		sort.Slice(horizontalLines[idx], func(i, j int) bool {
			return horizontalLines[idx][i].start.col < horizontalLines[idx][j].start.col
		})
	}
	logger.Debug("horizontalLines sorted")

	for idx := range verticalLines {
		sort.Slice(verticalLines[idx], func(i, j int) bool {
			return verticalLines[idx][i].start.row < verticalLines[idx][j].start.row
		})
	}
	logger.Debug("verticalLines sorted")

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
