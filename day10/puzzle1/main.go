package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const inputFile = "./input.txt"

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("failed to read input %q: %v", inputFile, err)
	}
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[0 : len(lines)-1]
	row := 0
	col := 0
	for i, line := range lines {
		row = i
		col = bytes.IndexByte(line, 'S')
		if col != -1 {
			break
		}
	}
	found := false
	total := func() int {
		depth := dfs(lines, row-1, col, true, &found) + 1
		if found {
			return depth
		}
		depth = dfs(lines, row+1, col, true, &found) + 1
		if found {
			return depth
		}
		depth = dfs(lines, row, col-1, true, &found) + 1
		if found {
			return depth
		}
		depth = dfs(lines, row, col+1, true, &found) + 1
		if found {
			return depth
		}
		return 0
	}()

	fmt.Println("Total steps: ", total/2, row, col)
}

func dfs(grid [][]byte, row, col int, first bool, found *bool) int {
	val := grid[row][col]
	grid[row][col] = '.'
	switch val {
	case '|':
		if canUp(grid, row, col, first) {
			return dfs(grid, row-1, col, false, found) + 1
		}
		if canDown(grid, row, col, first) {
			return dfs(grid, row+1, col, false, found) + 1
		}
	case '-':
		if canLeft(grid, row, col, first) {
			return dfs(grid, row, col-1, false, found) + 1
		}
		if canRight(grid, row, col, first) {
			return dfs(grid, row, col+1, false, found) + 1
		}
	case 'J':
		if canUp(grid, row, col, first) {
			return dfs(grid, row-1, col, false, found) + 1
		}
		if canLeft(grid, row, col, first) {
			return dfs(grid, row, col-1, false, found) + 1
		}
	case '7':
		if canDown(grid, row, col, first) {
			return dfs(grid, row+1, col, false, found) + 1
		}
		if canLeft(grid, row, col, first) {
			return dfs(grid, row, col-1, false, found) + 1
		}
	case 'F':
		if canDown(grid, row, col, first) {
			return dfs(grid, row+1, col, false, found) + 1
		}
		if canRight(grid, row, col, first) {
			return dfs(grid, row, col+1, false, found) + 1
		}
	case 'L':
		if canUp(grid, row, col, first) {
			return dfs(grid, row-1, col, false, found) + 1
		}
		if canRight(grid, row, col, first) {
			return dfs(grid, row, col+1, false, found) + 1
		}
	case 'S':
		*found = true
		return 0
	case '.':
		return 0
	}
	return 0
}

func canDown(grid [][]byte, row, col int, first bool) bool {
	return row < len(grid)-1 && grid[row+1][col] != '.' && !(first && grid[row+1][col] == 'S')
}
func canUp(grid [][]byte, row, col int, first bool) bool {
	return row > 0 && grid[row-1][col] != '.' && !(first && grid[row-1][col] == 'S')
}
func canLeft(grid [][]byte, row, col int, first bool) bool {
	return col > 0 && grid[row][col-1] != '.' && !(first && grid[row][col-1] == 'S')
}
func canRight(grid [][]byte, row, col int, first bool) bool {
	return col < len(grid[row])-1 && grid[row][col+1] != '.' && !(first && grid[row][col+1] == 'S')
}
