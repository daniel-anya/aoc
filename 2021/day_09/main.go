package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type LowPointLoc []int

func isLowPoint(point int, others []int) bool {
	for _, otherPoint := range others {
		if point >= otherPoint {
			return false
		}
	}
	return true
}

func sumInts(slice []int) int {
	result := 0
	for _, v := range slice {
		result += v
	}
	return result
}

func txtFileToMatrix(fileName string) [][]int {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var matrix [][]int
	contentStr := string(content)
	contentSlice := strings.Split(contentStr, "\n")
	for _, line := range contentSlice {
		var row []int
		lineSlice := strings.Split(line, "")
		for _, digitStr := range lineSlice {
			d, err := strconv.Atoi(digitStr)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, d)
		}
		matrix = append(matrix, row)
	}
	return matrix
}
func getPoint(matrix [][]int, i, j int) int {
	rowLength := len(matrix)
	colLength := len(matrix[0])

	if (i >= rowLength) || (i < 0) || (j >= colLength) || (j < 0) {
		return math.MaxInt64
	}
	return matrix[i][j]
}

func computeRiskLevelSum(matrix [][]int) (int, []LowPointLoc) {
	rowLength := len(matrix)
	var lowPoints []int
	var lowPointLocs []LowPointLoc

	for i := 0; i < rowLength; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			point := getPoint(matrix, i, j)
			left := getPoint(matrix, i, j-1)
			up := getPoint(matrix, i-1, j)
			right := getPoint(matrix, i, j+1)
			down := getPoint(matrix, i+1, j)

			if isLowPoint(point, []int{left, up, right, down}) {
				lowPoints = append(lowPoints, point)
				lowPointLocs = append(lowPointLocs, LowPointLoc{i, j})
			}
		}
	}
	return sumInts(lowPoints) + len(lowPoints), lowPointLocs
}

func countUntilBoundary(matrix [][]int, i, j int, seen map[string]bool) int {
	rowLength := len(matrix)
	colLength := len(matrix[0])

	iStr := strconv.Itoa(i)
	jStr := strconv.Itoa(j)
	pointStr := iStr + jStr

	if _, ok := seen[pointStr]; ok {
		return 0
	}

	if (i >= rowLength) || (i < 0) || (j >= colLength) || (j < 0) {
		return 0
	}

	if matrix[i][j] == 9 {
		return 0
	}

	seen[iStr+jStr] = true
	left := countUntilBoundary(matrix, i, j-1, seen)
	up := countUntilBoundary(matrix, i-1, j, seen)
	right := countUntilBoundary(matrix, i, j+1, seen)
	down := countUntilBoundary(matrix, i+1, j, seen)

	return 1 + left + up + right + down

}

func computeLargestBasinMultiple(matrix [][]int, LowPointLocs []LowPointLoc) int {
	var basinSizes []int
	for _, pointLoc := range LowPointLocs {
		i := pointLoc[0]
		j := pointLoc[1]
		seen := make(map[string]bool)
		basinSize := countUntilBoundary(matrix, i, j, seen)
		basinSizes = append(basinSizes, basinSize)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func main() {
	matrix := txtFileToMatrix("input.txt")
	riskLevelSum, lowPointLocs := computeRiskLevelSum(matrix)

	// Part 1 Answer
	fmt.Printf("Part 1 answer: %d \n", riskLevelSum)

	largestBasinMultiple := computeLargestBasinMultiple(matrix, lowPointLocs)

	// Part 2 Answer
	fmt.Printf("Part 2 answer: %d \n", largestBasinMultiple)

}
