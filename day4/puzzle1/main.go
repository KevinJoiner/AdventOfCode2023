package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFile = "./input.txt"

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("failed to read input %q: %v", inputFile, err)
	}
	cards := bytes.Split(input, []byte("\n"))
	cards = cards[0 : len(cards)-1]
	total := 0
	for _, card := range cards {
		colon := bytes.IndexByte(card, ':')
		if colon == -1 {
			continue
		}
		card = card[colon+1:]
		parts := bytes.Split(card, []byte("|"))
		desired := bytes.TrimSpace(parts[0])
		desiredList := strings.Split(string(desired), " ")
		have := bytes.TrimSpace(parts[1])
		winners := getWinners(bytes.Split(have, []byte(" ")), desiredList)
		if winners > 0 {
			total += 1 << (winners - 1)
		}
	}
	fmt.Println("Total points: ", total)
}

func getWinners(have [][]byte, desiredList []string) int {
	winners := 0
	// we are dealing with small slices so a double loop is fine over using a map
	for _, num := range have {
		for _, desired := range desiredList {
			if desired != "" && string(num) == desired {
				winners++
				break
			}
		}
	}
	return winners
}
