package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFile = "./input.txt"

type gearLoc struct {
	row int
	col int
}

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("failed to read input %q: %v", inputFile, err)
	}
	gears := map[gearLoc][]int{}
	total := 0
	grid := bytes.Split(input, []byte("\n"))
	grid = grid[:len(grid)-1] // remove trailing newline
	m := len(grid)
	n := len(grid[0])
	for row := range grid {
		start := 0

		adjGears := map[gearLoc]bool{}
		for col := range grid[row] {
			curr := grid[row][col]
			if !isDigit(curr) {
				start = col + 1
				continue
			}
			findAdjGears(grid, m, n, row, col, adjGears)
			// if group is adj and this is the last digit add to total
			if col == n-1 || !isDigit(grid[row][col+1]) {
				// final digit
				num, err := strconv.Atoi(string(grid[row][start : col+1]))
				if err != nil {
					log.Fatalf("failed to convert number: %v", err)
				}
				for loc, _ := range adjGears {
					gears[loc] = append(gears[loc], num)
				}
				adjGears = map[gearLoc]bool{}
			}
		}
	}
	for _, nums := range gears {
		if len(nums) == 2 {
			total += nums[0] * nums[1]
		}
	}
	fmt.Println("Total of gear ratios: ", total)
}

func findAdjGears(grid [][]byte, m, n, row, col int, locs map[gearLoc]bool) {
	// left
	if col > 0 && isSymbol(grid[row][col-1]) {
		locs[gearLoc{row, col - 1}] = true
	}
	if row > 0 {
		// up left
		if col > 0 && isSymbol(grid[row-1][col-1]) {
			locs[gearLoc{row - 1, col - 1}] = true
		}
		// up
		if isSymbol(grid[row-1][col]) {
			locs[gearLoc{row - 1, col}] = true
		}
		// up right
		if col < n-1 && isSymbol(grid[row-1][col+1]) {
			locs[gearLoc{row - 1, col + 1}] = true
		}
	}
	// right
	if col < n-1 && isSymbol(grid[row][col+1]) {
		locs[gearLoc{row, col + 1}] = true
	}
	if row < m-1 {
		// down right
		if col < n-1 && isSymbol(grid[row+1][col+1]) {
			locs[gearLoc{row + 1, col + 1}] = true
		}

		// down
		if isSymbol(grid[row+1][col]) {
			locs[gearLoc{row + 1, col}] = true
		}

		// down left
		if col > 0 && isSymbol(grid[row+1][col-1]) {
			locs[gearLoc{row + 1, col - 1}] = true
		}
	}
}

func isSymbol(char byte) bool {
	// not number or .
	return char == '*'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
