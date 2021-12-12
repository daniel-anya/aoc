package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func txtFileToIntSlice(fileName string) []int {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	contentStr := string(content)
	contentSlice := strings.Split(contentStr, "\n")

	var intSlice []int

	for _, strElem := range contentSlice{
		intElem, _ := strconv.Atoi(strElem)
		intSlice = append(intSlice, intElem)
	}

	return intSlice
}

func computeDepthIncreases(input []int) int{
	inputSize := len(input)
	curr, next := 0, 1

	var numIncreases int

	for next < inputSize {
		if input[next] > input[curr] {
			numIncreases += 1
		}
		curr += 1
		next += 1
	}
	return numIncreases
}

func computeSlidingWindowSums(input []int) []int{
	inputSize := len(input)
	 one, two, three := 0, 1, 2

	var sums []int

	for three < inputSize {
		sum := input[one] + input[two] + input[three]
		sums = append(sums, sum)
		one += 1
		two += 1
		three += 1
	}

	return sums
}

func main() {
	inputInts := txtFileToIntSlice("input.txt")
	numIncreases := computeDepthIncreases(inputInts)

	//Part 1 Answer
	fmt.Printf("Part 1 answer: %d \n", numIncreases)

	slidingWindowSums := computeSlidingWindowSums(inputInts)
	numSlidingWindowIncreases := computeDepthIncreases(slidingWindowSums)

	//Part 2 Answer
	fmt.Printf("Part 2 answer: %d \n", numSlidingWindowIncreases)
}