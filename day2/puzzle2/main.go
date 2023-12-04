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
	for _, line := range bytes.Split(input, []byte("\n")) {
		colon := bytes.IndexByte(line, ':')
		if colon == -1 {
			continue
		}
		total += gamePower(bytes.Split(line[colon+1:], []byte(";")))
	}
	fmt.Println("Total game Powers: ", total)
}

func gamePower(games [][]byte) int {
	maxRed := 1
	maxGreen := 1
	maxBlue := 1

	for _, game := range games {
		for _, pick := range bytes.Split(game, []byte(",")) {
			pick = bytes.TrimSpace(pick)
			parts := bytes.Split(pick, []byte(" "))
			num, err := strconv.Atoi(string(parts[0]))
			if err != nil {
				log.Fatalf("failed to get colors: %v", err)
			}
			switch string(parts[1]) {
			case "red":
				if num > maxRed {
					maxRed = num
				}
			case "blue":
				if num > maxBlue {
					maxBlue = num
				}
			case "green":
				if num > maxGreen {
					maxGreen = num
				}
			}
		}
	}
	return maxBlue * maxGreen * maxRed
}
