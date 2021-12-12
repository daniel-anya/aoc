package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Submarine struct {
	Depth         int
	HorizontalPos int
	Aim           int
}

func (s *Submarine) UpdateDepth(cmd command) {
	s.Depth += cmd.value
}

func (s *Submarine) UpdateHorizontalPos(cmd command) {
	s.HorizontalPos += cmd.value
}

func (s *Submarine) UpdateAimOrDepth(cmd command) {
	if cmd.referenceState == "depth" {
		s.Aim += cmd.value
	} else {
		s.Depth += cmd.value * s.Aim
	}
}

type command struct {
	value          int
	referenceState string
}

func newCommand(cmd string) command {
	commandSlice := strings.Split(cmd, " ")
	cmdStr := commandSlice[0]
	valStr := commandSlice[1]

	command := command{}
	if cmdStr == "forward" {
		command.referenceState = "horizontal-pos"
	} else {
		command.referenceState = "depth"
	}

	val, _ := strconv.Atoi(valStr)
	if cmdStr == "up" {
		command.value = -1 * val
	} else {
		command.value = val
	}

	return command
}

func txtFileToNewLineSepStrings(fileName string) []string {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	contentStr := string(content)
	contentSlice := strings.Split(contentStr, "\n")
	return contentSlice
}

func computeFinalPosition(cmds []string) int {
	sub := Submarine{}
	for _, cmdStr := range cmds {
		cmd := newCommand(cmdStr)
		if cmd.referenceState == "depth" {
			sub.UpdateDepth(cmd)
		} else {
			sub.UpdateHorizontalPos(cmd)
		}
	}

	return sub.HorizontalPos * sub.Depth
}

func computeFinalPositionWithAim(cmds []string) int {
	sub := Submarine{}
	for _, cmdStr := range cmds {
		cmd := newCommand(cmdStr)
		if cmd.referenceState == "horizontal-pos" {
			sub.UpdateHorizontalPos(cmd)
		}
		sub.UpdateAimOrDepth(cmd)
	}

	return sub.HorizontalPos * sub.Depth
}

func main() {
	input := txtFileToNewLineSepStrings("input.txt")
	finalPos := computeFinalPosition(input)

	// Part 1 Answer
	fmt.Printf("Part 1 answer: %d \n", finalPos)

	finalPosWithAim := computeFinalPositionWithAim(input)

	// Part 2 Answer
	fmt.Printf("Part 2 answer: %d \n", finalPosWithAim)
}
