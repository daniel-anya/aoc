package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

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

func incrementMatrixByVal(matrix [][]int, val int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] += val
		}
	}
	return matrix
}

func shouldFlashFn(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			point := matrix[i][j]
			if point > 9 {
				return true
			}
		}
	}
	return false
}

func getSurroundingPoints(matrix [][]int, i, j int) [][]int {
	var surroundingPoints [][]int
	rowLength := len(matrix)
	colLength := len(matrix[0])

	potentialPoints := [][]int{
		{i, j + 1},
		{i, j - 1},
		{i + 1, j},
		{i - 1, j},
		{i + 1, j + 1},
		{i + 1, j - 1},
		{i - 1, j + 1},
		{i - 1, j - 1},
	}
	for _, point := range potentialPoints {
		if (point[0] >= 0) && (point[0] < rowLength) && (point[1] >= 0) && (point[1] < colLength) {
			surroundingPoints = append(surroundingPoints, point)
		}
	}
	return surroundingPoints
}

func incrementPointsByVal(matrix, points [][]int, val int) {
	for _, point := range points {
		i := point[0]
		j := point[1]

		matrix[i][j] += 1
	}
}

func setPointsToVal(matrix, points [][]int, val int) {
	for _, point := range points {
		i := point[0]
		j := point[1]

		matrix[i][j] = val
	}
}

func flash(matrix [][]int, alreadyFlashedInCurrentStep map[string]bool) {
	var toFlash [][]int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			point := matrix[i][j]
			if point > 9 {
				toFlash = append(toFlash, []int{i, j})
			}
		}
	}

	for _, point := range toFlash {

		iStr := strconv.Itoa(point[0])
		jStr := strconv.Itoa(point[1])
		pointStr := iStr + jStr

		if _, ok := alreadyFlashedInCurrentStep[pointStr]; ok {
			continue
		}
		surroundingPoints := getSurroundingPoints(matrix, point[0], point[1])
		incrementPointsByVal(matrix, surroundingPoints, 1)
		alreadyFlashedInCurrentStep[pointStr] = true
	}
}

func getPointsSliceFromMapKeys(pointsMap map[string]bool) [][]int {
	var pointsSlice [][]int
	for key := range pointsMap {
		point := strings.Split(key, "")
		i, _ := strconv.Atoi(point[0])
		j, _ := strconv.Atoi(point[1])
		pointsSlice = append(pointsSlice, []int{i, j})
	}
	return pointsSlice
}

func simulateStep(matrix [][]int) int {
	matrix = incrementMatrixByVal(matrix, 1)
	alreadyFlashedInCurrentStep := make(map[string]bool)
	shouldFlash := shouldFlashFn(matrix)
	for shouldFlash {
		flash(matrix, alreadyFlashedInCurrentStep)
		flashedPoints := getPointsSliceFromMapKeys(alreadyFlashedInCurrentStep)
		setPointsToVal(matrix, flashedPoints, 0)
		shouldFlash = shouldFlashFn(matrix)
	}

	return len(alreadyFlashedInCurrentStep)
}

func simulateSteps(matrix [][]int, numSteps int) int {
	var totalFlashes int
	for i := 0; i < numSteps; i++ {
		numFlashes := simulateStep(matrix)
		totalFlashes += numFlashes
	}
	return totalFlashes
}

func getFirstStepOfUnisonFlash(matrix [][]int) int {
	var firstStepOfUnisonFlash int
	for {
		firstStepOfUnisonFlash += 1
		numFlashes := simulateStep(matrix)
		if numFlashes == 100 {
			break
		}
	}
	return firstStepOfUnisonFlash
}

func main() {
	matrix := txtFileToMatrix("input.txt")
	totalFlashes := simulateSteps(matrix, 100)

	// Part 1 Answer
	fmt.Printf("Part 1 answer: %d \n", totalFlashes)

	// Part 2 Answer
	matrix = txtFileToMatrix("input.txt")
	firstStep := getFirstStepOfUnisonFlash(matrix)
	fmt.Printf("Part 2 answer: %d \n", firstStep)
}
