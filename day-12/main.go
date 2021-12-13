/* Day 12: Passage Pathing */
package main

import (
	"AoC2021/utils/log"
	"AoC2021/utils/timer"
	"bufio"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		log.Println(log.DEBUG, "Oh snap!")
		panic(e)
	}
}

// Assumptions: The input data won't ever put us in an infinite loop of large caverns
// This might be something to protect against if this was production code

func main() {
	log.Println(log.NORMAL, "--- Day 12: Passage Pathing ---")
	timer.Start()
	part1()
	timer.Tick()
	log.Println(log.NORMAL)

	part2()
	timer.Tick()
	log.Println(log.NORMAL)
}

func part1() {
	log.Println(log.NORMAL, "* Part 1 *")
	log.Println(log.NORMAL, " Goal: Find all valid path to navigate")
	log.Println(log.NORMAL, " Answer: How many paths through this cave system are there that visit small caves at most once?")

	nodeMap := loadNodeMap()
	log.Println(log.DIAGNOSTIC, nodeMap)

	log.Println(log.DEBUG, "The start node is", nodeMap["start"], "with connections", nodeMap["start"].connections)

	paths := getPaths(nodeMap, 0)

	log.Printf(log.NORMAL, "Found %d paths from 'start' to 'end'!\n", len(paths))
}

func part2() {
	log.Println(log.NORMAL, "* Part 2 *")
	log.Println(log.NORMAL, " Goal: Find all valid path to navigate")
	log.Println(log.NORMAL, " Answer: Given these new rules, how many paths through this cave system are there?")

	nodeMap := loadNodeMap()
	log.Println(log.DIAGNOSTIC, nodeMap)

	log.Println(log.DEBUG, "The start node is", nodeMap["start"], "with connections", nodeMap["start"].connections)

	paths := getPaths(nodeMap, 1)

	// for idx, p := range paths {
	// 	if p[0] != "start" || p[len(p)-1] != "end" {
	// 		panic(fmt.Sprintf("Error: path %d is invalid: %s\n", idx, p))
	// 	}
	// }

	log.Printf(log.NORMAL, "Found %d paths from 'start' to 'end'!\n", len(paths))
}

type Node struct {
	name        string
	isLarge     bool
	connections *[]string
}

type NodeMap = map[string]Node

func loadNodeMap() NodeMap {

	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := 0
	nodeMap := make(NodeMap)
	connections := 0

	for scanner.Scan() {
		nodeStrs := strings.Split(scanner.Text(), "-")
		lines++
		if len(nodeStrs) != 2 {
			panic("Did not read in expected 2 nodes from input line")
		}
		// Check if either node already exists, add them to the slice, and add them to each other's connections
		node1Str, node2Str := nodeStrs[0], nodeStrs[1]
		log.Printf(log.DIAGNOSTIC, "Processing nodes %s & %s: ", node1Str, node2Str)

		node1, node1exists := nodeMap[node1Str]
		if !node1exists {
			isLarge := unicode.IsUpper(rune(node1Str[0]))
			nodeMap[node1Str] = Node{node1Str, isLarge, &[]string{}}
			node1 = nodeMap[node1Str]
		}

		node2, node2exists := nodeMap[node2Str]
		if !node2exists {
			isLarge := unicode.IsUpper(rune(node2Str[0]))
			nodeMap[node2Str] = Node{node2Str, isLarge, &[]string{}}
			node2 = nodeMap[node2Str]
		}

		*node1.connections = append(*node1.connections, node2Str)
		*node2.connections = append(*node2.connections, node1Str)
		log.Println(log.DIAGNOSTIC, node1, node2)
		connections += 2
	}

	log.Printf(log.NORMAL, "Read %d nodes with %d connections from %d lines from input\n", len(nodeMap), connections, lines)

	return nodeMap
}

func getPaths(nodeMap NodeMap, revisitsAllowed int) [][]string {
	visitCountMap := initVisitCountMap(nodeMap)

	// start at 'start' and work our way to 'end'
	path := []string{"start"}
	log.Printf(log.DEBUG, "Starting iteration of nodes from 'start'\n")

	paths := getValidPathsFrom("start", path, nodeMap, visitCountMap, revisitsAllowed)

	return paths
}

type VisitMap = map[string]int

func initVisitCountMap(nodeMap NodeMap) VisitMap {
	visitCountMap := make(VisitMap, len(nodeMap))
	for n := range nodeMap {
		visitCountMap[n] = 0
	}

	return visitCountMap
}

func getValidPathsFrom(node string, path []string, nodeMap NodeMap, visitMap VisitMap, revisitsAllowed int) [][]string {
	var paths [][]string
	// just stop once we reach the end
	if node == "end" {
		log.Printf(log.DIAGNOSTIC, "The final visit count looks like %s\n", visitMap)
		return [][]string{path}
	}

	n, exists := nodeMap[node]
	if exists {
		visitMap[node] += 1
		for _, cStr := range *n.connections {
			// do not return to the 'start' node
			// allow traveling to large nodes
			// allow traveling to small nodes that we have not visited
			// allow traveling to small nodes when we still have some 'revisitAllowed' points left
			if cStr != "start" && (nodeMap[cStr].isLarge || visitMap[cStr] == 0 || revisitsAllowed > 0) {
				// if we triggered due to 'revisitAllowed' points, decrease the point pool for the next time
				nextRevisitsAllowed := revisitsAllowed
				if !nodeMap[cStr].isLarge && visitMap[cStr] > 0 && revisitsAllowed > 0 {
					nextRevisitsAllowed -= 1
				}
				newBasePath := append(path, cStr)
				log.Printf(log.DIAGNOSTIC, "Processed path %s\n", newBasePath)
				newPaths := getValidPathsFrom(cStr, newBasePath, nodeMap, copyMap(visitMap), nextRevisitsAllowed)
				if len(newPaths) > 0 {
					paths = append(paths, newPaths...)
				}
			} else {
				log.Printf(log.DIAGNOSTIC, "Did NOT follow %s %s (visit count: %d) and pruned this path\n", path, cStr, visitMap[cStr])
			}
		}
	}

	return paths
}

func copyMap(v1 VisitMap) VisitMap {
	v2 := VisitMap{}
	for v := range v1 {
		v2[v] = v1[v]
	}
	return v2
}
