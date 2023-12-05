package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const inputFile = "./input.txt"

type Range struct {
	srcStart  int
	destStart int
	run       int
}

func (r *Range) Get(src int) (int, bool) {
	if r.srcStart > src || src > r.srcStart+r.run {
		return 0, false
	}
	return r.destStart + src - r.srcStart, true
}

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("failed to read input %q: %v", inputFile, err)
	}
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[0 : len(lines)-1]
	seedLine := lines[0]
	lines = lines[1:]
	currMap := []Range{}
	maps := [][]Range{}
	for _, line := range lines {
		if bytes.Contains(line, []byte("map")) {
			maps = append(maps, currMap)
			currMap = []Range{}
			continue
		}
		if len(line) == 0 {
			continue
		}

		parts := bytes.Split(line, []byte(" "))
		destStart, _ := strconv.Atoi(string(parts[0]))
		srcStart, _ := strconv.Atoi(string(parts[1]))
		run, _ := strconv.Atoi(string(parts[2]))
		currMap = append(currMap, Range{srcStart: srcStart, destStart: destStart, run: run})
	}
	maps = append(maps, currMap)
	min := math.MaxInt
	seeds := bytes.Split(seedLine, []byte(" "))[1:]
	for _, seed := range seeds {
		seedInt, _ := strconv.Atoi(string(seed))
		loc := getLocation(seedInt, maps)
		if loc < min {
			min = loc
		}
	}
	fmt.Println("Minimum location is: ", min)

}

func getLocation(start int, maps [][]Range) int {
	cur := start
	for _, curMap := range maps {
		for _, curRange := range curMap {
			next, ok := curRange.Get(cur)
			if ok {
				cur = next
				break
			}
		}
	}
	return cur
}
