package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type stack struct {
	slice []string
}

func (s *stack) insert(item string) {
	s.slice = append(s.slice, item)
}

func (s *stack) pop() string {
	size := len(s.slice)
	item := s.slice[size-1]
	s.slice = s.slice[0:(size - 1)]
	return item
}

func (s *stack) peek() string {
	size := len(s.slice)
	item := s.slice[size-1]
	return item
}

func (s *stack) size() int {
	return len(s.slice)
}

func getStartChunkMap() map[string]bool {
	return map[string]bool{"(": true, "[": true, "{": true, "<": true}
}
func getEndChunkMap() map[string]bool {
	return map[string]bool{")": true, "]": true, "}": true, ">": true}
}

func getValidPairsMap() map[string]bool {
	return map[string]bool{"()": true, "[]": true, "{}": true, "<>": true}
}

func inputToLines(fileName string) []string {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	contentStr := string(content)
	contentSlice := strings.Split(contentStr, "\n")
	return contentSlice
}

func computeCorruptTally(chunkLines []string) int {
	var tally int
	startChunk := getStartChunkMap()
	endChunk := getEndChunkMap()
	validPairs := getValidPairsMap()
	pointsMap := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}

OUTER:
	for _, line := range chunkLines {
		curStack := stack{}
		charSlice := strings.Split(line, "")
		for _, char := range charSlice {
			if _, ok := startChunk[char]; ok {
				curStack.insert(char)
			}
			if _, ok := endChunk[char]; ok {
				lastChar := curStack.peek()
				if _, ok := validPairs[lastChar+char]; ok {
					curStack.pop()
					continue
				}
				tally += pointsMap[char]
				continue OUTER
			}
		}
	}

	return tally
}

func isIncomplete(line []string) (bool, []string) {
	startChunk := getStartChunkMap()
	endChunk := getEndChunkMap()
	validPairs := getValidPairsMap()

	curStack := stack{}
	for _, char := range line {
		if _, ok := startChunk[char]; ok {
			curStack.insert(char)
		}

		if _, ok := endChunk[char]; ok {
			lastChar := curStack.peek()
			if _, ok := validPairs[lastChar+char]; ok {
				curStack.pop()
				continue
			}
			curStack = stack{}
			break
		}
	}

	if curStack.size() == 0 {
		return false, []string{}
	}
	return true, curStack.slice
}

func computeReplacementScore(chunkLines []string) int {
	var tallies []int
	completionStringsPointsMap := map[string]int{")": 1, "]": 2, "}": 3, ">": 4}
	startToCloseCharMap := map[string]string{"(": ")", "[": "]", "{": "}", "<": ">"}

	for _, line := range chunkLines {
		charSlice := strings.Split(line, "")
		isIncompleteLine, incompleteChars := isIncomplete(charSlice)

		if !isIncompleteLine {
			continue
		}

		var tally int
		for i := len(incompleteChars) - 1; i >= 0; i-- {
			startChar := incompleteChars[i]
			closingChar := startToCloseCharMap[startChar]
			charScore := completionStringsPointsMap[closingChar]
			tally = (tally * 5) + charScore
		}
		tallies = append(tallies, tally)
	}
	sort.Ints(tallies)

	return tallies[len(tallies)/2]
}

func main() {

	chunkLines := inputToLines("input.txt")
	corruptTally := computeCorruptTally(chunkLines)
	middleReplacementScore := computeReplacementScore(chunkLines)

	//Part 1 Answer
	fmt.Println(corruptTally)

	// Part 2 Answer
	fmt.Println(middleReplacementScore)

}
