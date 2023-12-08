package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

type handStrength int8

const (
	none handStrength = iota
	pair
	twoPair
	threeOfaKind
	fullHouse
	fourOfaKind
	fiveOfaKind
)

const inputFile = "./input.txt"

func getVal(b byte) int {
	if val, ok := vals[b]; ok {
		return val
	}
	return int(b)
}

var vals = map[byte]int{
	'A': '9' + 5,
	'K': '9' + 4,
	'Q': '9' + 3,
	'J': '0' - 1,
	'T': '9' + 1,
}

type cardHand struct {
	cards    []byte
	strength handStrength
	bid      int
}

func cmpHand(hand1, hand2 *cardHand) int {
	if hand1.strength < hand2.strength {
		return -1
	}
	if hand1.strength > hand2.strength {
		return 1
	}
	if hand1.strength == hand2.strength {
		for i := range hand1.cards {
			b1 := getVal(hand1.cards[i])
			b2 := getVal(hand2.cards[i])
			if b1 < b2 {
				return -1
			}
			if b1 > b2 {
				return 1
			}
		}
	}
	return 0
}

func New(cards []byte, bid int) *cardHand {
	hand := cardHand{cards: cards, bid: bid}
	seen := make(map[byte]int, len(cards))
	max := 0
	for _, card := range cards {
		seen[card]++
		if card != 'J' && seen[card] > max {
			max = seen[card]
		}
	}
	sortedVal := []int{}
	for card, val := range seen {
		if card == 'J' {
			continue
		}
		sortedVal = append(sortedVal, val)
	}
	slices.Sort(sortedVal)
	if len(sortedVal) == 0 {
		sortedVal = []int{0}
	}
	sortedVal[len(sortedVal)-1] += seen['J']

loop:
	for i := len(sortedVal) - 1; i >= 0; i-- {
		val := sortedVal[i]
		switch val {
		case 5:
			hand.strength = fiveOfaKind
			break loop
		case 4:
			hand.strength = fourOfaKind
			break loop
		case 3:
			if hand.strength == pair {
				hand.strength = fullHouse
				break loop
			}
			hand.strength = threeOfaKind
		case 2:
			if hand.strength == threeOfaKind {
				hand.strength = fullHouse
				break loop
			}
			if hand.strength == pair {
				hand.strength = twoPair
				break loop
			}
			hand.strength = pair
		}
	}

	return &hand
}

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("failed to read input %q: %v", inputFile, err)
	}
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[0 : len(lines)-1]
	hands := make([]*cardHand, 0, len(lines))
	for _, line := range lines {
		parts := bytes.Fields(line)
		bid, _ := strconv.Atoi(string(parts[1]))
		hands = append(hands, New(parts[0], bid))
	}
	slices.SortFunc(hands, cmpHand)
	total := 0

	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}

	fmt.Println("Total winnnings: ", total)
}
