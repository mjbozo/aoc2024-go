package day23

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"slices"
	"strings"
	"time"
)

func Run() {
	input, err := utils.ReadInput("day23/input.txt", 23)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	start := time.Now()
	part1Result := part1(input)
	elapsed := time.Since(start)
	fmt.Printf("Part 1: %d (%v)\n", part1Result, elapsed)

	start = time.Now()
	part2Result := part2(input)
	elapsed = time.Since(start)
	fmt.Printf("Part 2: %s (%v)\n", part2Result, elapsed)
}

func part1(lines []string) int {
	connections := make(map[string]utils.HashSet[string])
	for _, line := range lines {
		computers := strings.Split(line, "-")
		computer1 := computers[0]
		computer2 := computers[1]

		// add to both map entries
		if links, ok := connections[computer1]; ok {
			links.Insert(computer2)
			connections[computer1] = links
		} else {
			links := make(utils.HashSet[string])
			links.Insert(computer2)
			connections[computer1] = links
		}

		if links, ok := connections[computer2]; ok {
			links.Insert(computer1)
			connections[computer2] = links
		} else {
			links := make(utils.HashSet[string])
			links.Insert(computer1)
			connections[computer2] = links
		}
	}

	total := 0
	triples := make(utils.HashSet[string])
	for key, value := range connections {
		if strings.HasPrefix(key, "t") {
			for computer := range value {
				otherConnections := connections[computer]
				for other := range otherConnections {
					nodes := []string{key, computer, other}
					slices.Sort(nodes)
					nodesCombined := strings.Join(nodes, "-")
					if value.Contains(other) && !triples.Contains(nodesCombined) && other != key {
						triples.Insert(nodesCombined)
						total++
					}
				}
			}
		}
	}

	return total
}

func part2(lines []string) string {
	// adjacency list of nodes
	connections := make(map[string]utils.HashSet[string])
	P := make(utils.HashSet[string])

	for _, line := range lines {
		computers := strings.Split(line, "-")
		computer1 := computers[0]
		computer2 := computers[1]

		P.Insert(computer1)
		P.Insert(computer2)

		// add to both map entries
		if links, ok := connections[computer1]; ok {
			links.Insert(computer2)
			connections[computer1] = links
		} else {
			links := make(utils.HashSet[string])
			links.Insert(computer2)
			connections[computer1] = links
		}

		if links, ok := connections[computer2]; ok {
			links.Insert(computer1)
			connections[computer2] = links
		} else {
			links := make(utils.HashSet[string])
			links.Insert(computer1)
			connections[computer2] = links
		}
	}

	R := make(utils.HashSet[string])
	X := make(utils.HashSet[string])

	maxCliques := make([]utils.HashSet[string], 0)
	bronKerbosch(R, P, X, connections, &maxCliques)

	var maximumClique utils.HashSet[string]
	for _, clique := range maxCliques {
		if len(clique) > len(maximumClique) {
			maximumClique = clique
		}
	}

	maximumNodes := make([]string, 0)
	for clique := range maximumClique {
		maximumNodes = append(maximumNodes, clique)
	}

	slices.Sort(maximumNodes)
	return strings.Join(maximumNodes, ",")
}

func bronKerbosch(R, P, X utils.HashSet[string], graph map[string]utils.HashSet[string], maxCliques *[]utils.HashSet[string]) {
	if len(P) == 0 && len(X) == 0 {
		// found maximal clique
		*maxCliques = append(*maxCliques, R)
		return
	}

	for v := range P {
		nextR := make(utils.HashSet[string])
		nextP := make(utils.HashSet[string])
		nextX := make(utils.HashSet[string])

		for node := range R {
			nextR.Insert(node)
		}
		nextR.Insert(v)

		for neighbour := range graph[v] {
			if P.Contains(neighbour) {
				nextP.Insert(neighbour)
			}

			if X.Contains(neighbour) {
				nextX.Insert(neighbour)
			}
		}

		bronKerbosch(nextR, nextP, nextX, graph, maxCliques)

		P.Remove(v)
		X.Insert(v)
	}
}
