package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

const inputFile = "./input.txt"

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("failed to read input %q: %v", inputFile, err)
	}

	total := 0
	headerLen := len("Game ")
	for _, line := range bytes.Split(input, []byte("\n")) {
		colon := bytes.IndexByte(line, ':')
		if colon == -1 {
			continue
		}
		numB := line[headerLen:colon]
		gameID, err := strconv.Atoi(string(numB))
		if err != nil {
			log.Fatalf("failed to get gameID: %v", err)
		}
		if isPossible(bytes.Split(line[colon+1:], []byte(";"))) {
			total += gameID
		}

	}
	fmt.Println("Total gameIDs: ", total)
}

func isPossible(games [][]byte) bool {
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
					return false
				}
			case "blue":
				if num > maxBlue {
					return false
				}
			case "green":
				if num > maxGreen {
					return false
				}
			}
		}
	}
	return true
}
