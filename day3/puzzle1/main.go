package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFile = "./input.txt"

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("failed to read input %q: %v", inputFile, err)
	}

	total := 0
	grid := bytes.Split(input, []byte("\n"))
	grid = grid[:len(grid)-1] // remove trailing newline
	m := len(grid)
	n := len(grid[0])
	for row := range grid {
		found := false
		start := 0
		for col := range grid[row] {
			curr := grid[row][col]
			if !isDigit(curr) {
				start = col + 1
				continue
			}
			if !found && isAdj(grid, m, n, row, col) {
				found = true
			}
			// if group is adj and this is the last digit add to total
			if found && (col == n-1 || !isDigit(grid[row][col+1])) {
				// final digit
				num, err := strconv.Atoi(string(grid[row][start : col+1]))
				if err != nil {
					log.Fatalf("failed to convert number: %v", err)
				}
				total += num
				found = false
			}
		}
	}
	fmt.Println("Total of product numbers: ", total)
}

func isAdj(grid [][]byte, m, n, row, col int) bool {
	// left
	if col > 0 && isSymbol(grid[row][col-1]) {
		return true
	}
	if row > 0 {
		// up left
		if col > 0 && isSymbol(grid[row-1][col-1]) {
			return true
		}
		// up
		if isSymbol(grid[row-1][col]) {
			return true
		}
		// up right
		if col < n-1 && isSymbol(grid[row-1][col+1]) {
			return true
		}
	}
	// right
	if col < n-1 && isSymbol(grid[row][col+1]) {
		return true
	}
	if row < m-1 {
		// down right
		if col < n-1 && isSymbol(grid[row+1][col+1]) {
			return true
		}

		// down
		if isSymbol(grid[row+1][col]) {
			return true
		}

		// down left
		if col > 0 && isSymbol(grid[row+1][col-1]) {
			return true
		}
	}
	return false
}

func isSymbol(char byte) bool {
	// not number or .
	return !(isDigit(char) || char == '.')
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
