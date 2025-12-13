package day09

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Dim    = "\033[2m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

type coord struct {
	row int
	col int
}

func colourStr(s, c string) string {
	x := c + s
	if c != "" {
		x += Reset
	}
	return x
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

func addCoord(i, j coord) (newCoord coord) {
	newCoord.row = i.row + j.row
	newCoord.col = i.col + j.col
	return
}

func printGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}

func printGridFillLoop(grid [][]rune, queue []coord, visited map[coord]struct{}, p coord) {
	queuedMap := make(map[coord]struct{})
	for _, p := range queue {
		queuedMap[p] = struct{}{}
	}

	for r := range grid {
		line := ""
		for c, x := range grid[r] {
			colour := ""
			if _, ok := visited[coord{r, c}]; !ok {
				colour += Dim
			}
			if _, ok := queuedMap[coord{r, c}]; ok {
				colour += Yellow
			} else if x == '#' {
				colour += Green
			} else if x == 'X' {
				colour += Red
			}
			if p.row == r && p.col == c {
				colour += Bold
			}
			line += colourStr(string(x), colour)
		}
		fmt.Println(string(line))
	}
}

func minInt(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func maxInt(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func checkBoundaries(grid [][]bool, i, j coord) bool {
	minRow := minInt(i.row, j.row)
	minCol := minInt(i.col, j.col)
	maxRow := maxInt(i.row, j.row)
	maxCol := maxInt(i.col, j.col)
	for c := minCol; c <= maxCol; c++ {
		if !grid[minRow][c] || !grid[maxRow][c] {
			return false
		}
	}
	for r := minRow; r <= maxRow; r++ {
		if !grid[r][minCol] || !grid[r][maxCol] {
			return false
		}
	}
	return true
}

func addToQueue(queue []coord, queued map[coord]struct{}, coords ...coord) []coord {
	for _, c := range coords {
		if _, ok := queued[c]; !ok {
			queue = append(queue, c)
			queued[c] = struct{}{}
		}
	}
	return queue
}

func fillLoop(grid [][]bool) {
	rows := len(grid)
	cols := len(grid[0])
	visited := make(map[coord]struct{})

	dirs := []coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := []coord{}
	queued := make(map[coord]struct{})

	for c := range cols {
		queue = addToQueue(queue, queued, coord{0, c}, coord{rows - 1, c})
	}
	for r := range rows {
		queue = addToQueue(queue, queued, coord{r, 0}, coord{r, cols - 1})
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		delete(queued, p)

		if len(visited)%10000000 == 0 {
			slog.Debug("fillLoop", "queueLen", len(queue), "visited", len(visited), "p", p)
			// printGridFillLoop(grid, queue, visited, p)
		}

		if p.row < 0 || p.row >= rows || p.col < 0 || p.col >= cols {
			// slog.Debug("fillLoop", "msg", "out of bounds", "p", p)
			continue
		}
		if grid[p.row][p.col] {
			// slog.Debug("fillLoop", "msg", "grid boundary", "p", p)
			continue
		}
		if _, ok := visited[p]; ok {
			// slog.Debug("fillLoop", "msg", "already visited", "p", p)
			continue
		}

		visited[p] = struct{}{}

		for _, d := range dirs {
			newP := coord{p.row + d.row, p.col + d.col}
			if newP.row < 0 || newP.row >= rows || newP.col < 0 || newP.col >= cols {
				continue
			}
			if _, ok := visited[newP]; ok {
				continue
			}
			if _, ok := queued[newP]; ok {
				continue
			}

			queue = addToQueue(queue, queued, newP)
		}
	}

	for r := range rows {
		for c := range cols {
			if _, ok := visited[coord{r, c}]; !ok && !grid[r][c] {
				grid[r][c] = true
			}
		}
	}
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

func Part2(input []string) (string, error) {
	coords := parseCoords(input)
	minCoord := coords[0]
	maxCoord := coords[0]
	for _, c := range coords {
		if c.row < minCoord.row {
			minCoord.row = c.row
		}
		if c.col < minCoord.col {
			minCoord.col = c.col
		}
		if c.row > maxCoord.row {
			maxCoord.row = c.row
		}
		if c.col > maxCoord.col {
			maxCoord.col = c.col
		}
	}
	slog.Debug("part 2", "minCoord", minCoord, "maxCoord", maxCoord)

	for i := range coords {
		coords[i].row -= minCoord.row
		coords[i].col -= minCoord.col
	}

	grid := make([][]bool, maxCoord.row-minCoord.row+1)
	for i := range grid {
		grid[i] = make([]bool, maxCoord.col-minCoord.col+1)
		for j := range grid[i] {
			grid[i][j] = false
		}
	}
	slog.Debug("part 2", "msg", "grid initialised")

	for i := range coords {
		grid[coords[i].row][coords[i].col] = true
		j := i + 1
		if j == len(coords) {
			j = 0
		}

		vector := coord{0, 0}
		if coords[i].row != coords[j].row {
			if coords[i].row < coords[j].row {
				vector.row = 1
			} else {
				vector.row = -1
			}
		}
		if coords[i].col != coords[j].col {
			if coords[i].col < coords[j].col {
				vector.col = 1
			} else {
				vector.col = -1
			}
		}

		for c := addCoord(coords[i], vector); c != coords[j]; c = addCoord(c, vector) {
			grid[c.row][c.col] = true
		}
	}
	slog.Debug("part 2", "msg", "grid borders added")

	// for i := range grid {
	// 	inside := false
	// 	for j := range grid[i] {
	// 		if !inside {
	// 			if grid[i][j] != '.' {
	// 				inside = true
	// 			}
	// 		} else {
	// 			if grid[i][j] == '.' {
	// 				grid[i][j] = 'X'
	// 			} else {
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	fillLoop(grid)
	slog.Debug("part 2", "msg", "grid borders filled")

	// printGrid(grid)

	maxArea := 0
	for i := 0; i < len(coords)-1; i++ {
		slog.Debug("part 2", "i", i)
		for j := i + 1; j < len(coords); j++ {
			if !checkBoundaries(grid, coords[i], coords[j]) {
				continue
			}
			area := area(coords[i], coords[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return fmt.Sprintf("%d", maxArea), nil
}
