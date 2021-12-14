package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

type Node struct {
	nodeId   string
	children []*Node
}

type Seen struct {
	parents []string
}

func (s *Seen) Contains(elem string) bool {
	for _, item := range s.parents {
		if item == elem {
			return true
		}
	}
	return false
}

func (s *Seen) Insert(elem string) {
	s.parents = append(s.parents, elem)
}

func (s *Seen) DiscardLast() {
	size := len(s.parents)
	s.parents = s.parents[0:(size - 1)]
}

func (s *Seen) TallyNodes() (map[string]int, bool) {
	counter := make(map[string]int)
	smallCaveVisitedTwice := false

	for _, elem := range s.parents {
		counter[elem] += 1
	}

	for elem, count := range counter {
		if isLower(elem) && count == 2 {
			smallCaveVisitedTwice = true
		}
	}

	return counter, smallCaveVisitedTwice
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

func parseLinesToNodes(edges []string) map[string]*Node {
	nodeMap := make(map[string]*Node)
	for _, edge := range edges {
		nodes := strings.Split(edge, "-")
		a := nodes[0]
		b := nodes[1]

		if _, ok := nodeMap[a]; !ok {
			nodeMap[a] = &Node{
				nodeId: a,
			}
		}

		if _, ok := nodeMap[b]; !ok {
			nodeMap[b] = &Node{
				nodeId: b,
			}
		}

		startNode := nodeMap[a]
		endNode := nodeMap[b]
		startNode.children = append(startNode.children, endNode)
		endNode.children = append(endNode.children, startNode)
	}

	return nodeMap
}

func getOptimalPaths(nodeMap map[string]*Node, predicate func(nodeId string, s *Seen) bool) []string {
	startNode := nodeMap["start"]
	seen := &Seen{[]string{}}
	paths := getPaths(startNode, seen, predicate)

	var optimalPaths []string
	for _, path := range paths {
		if strings.HasPrefix(path, "start") && strings.HasSuffix(path, "end") {
			optimalPaths = append(optimalPaths, path)
		}
	}
	return optimalPaths
}

func partOnePredicate(nodeId string, s *Seen) bool {
	return isLower(nodeId) && s.Contains(nodeId)
}

func partTwoPredicate(nodeId string, s *Seen) bool {
	nodeTally, smallCaveVisitedTwice := s.TallyNodes()
	if (nodeId == "start" && nodeTally["start"] == 1) || (nodeId == "end" && nodeTally["end"] == 1) {
		return true
	}

	if isLower(nodeId) && nodeTally[nodeId] == 2 {
		return true
	}

	if isLower(nodeId) && nodeTally[nodeId] == 1 && smallCaveVisitedTwice {
		return true
	}
	return false
}

func getPaths(node *Node, seen *Seen, predicate func(nodeId string, s *Seen) bool) []string {
	var nodePath []string
	if node.nodeId == "end" {
		return []string{node.nodeId}
	}

	if predicate(node.nodeId, seen) {
		return []string{""}
	}

	seen.Insert(node.nodeId)
	for _, child := range node.children {
		childPaths := getPaths(child, seen, predicate)
		for _, path := range childPaths {
			nodePath = append(nodePath, fmt.Sprintf("%s,%s", node.nodeId, path))
		}
	}
	seen.DiscardLast()
	return nodePath
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func main() {
	inputs := txtFileToNewLineSepStrings("input.txt")
	nodeMap := parseLinesToNodes(inputs)
	numOptimalPaths := len(getOptimalPaths(nodeMap, partOnePredicate))
	numOptimalPathsPartTwo := len(getOptimalPaths(nodeMap, partTwoPredicate))

	// Part 1 Answer
	fmt.Printf("Part 1 answer: %d \n", numOptimalPaths)

	// Part 2 Answer
	fmt.Printf("Part 2 answer: %d \n", numOptimalPathsPartTwo)
}
