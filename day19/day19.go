package day19

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strings"
	"time"
)

func Run() {
	input, err := utils.ReadInputRaw("day19/input.txt")
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
	fmt.Printf("Part 2: %d (%v)\n", part2Result, elapsed)
}

func part1(input string) int {
	parts := strings.Split(input, "\n\n")
	available := strings.Split(parts[0], ", ")
	desired := strings.Split(parts[1], "\n")

	memo := make(map[string]bool)

	possibleTowels := 0
	for _, pattern := range desired {
		if canMakeTowel(pattern, available, &memo) {
			possibleTowels++
		}
	}

	return possibleTowels
}

func part2(input string) int {
	parts := strings.Split(input, "\n\n")
	available := strings.Split(parts[0], ", ")
	desired := strings.Split(parts[1], "\n")

	memo := make(map[string]int)

	towelWays := 0
	for _, pattern := range desired {
		towelWays += waysToMakeTowel(pattern, available, &memo)
	}

	return towelWays
}

func canMakeTowel(towelDesign string, available []string, memo *map[string]bool) bool {
	if possible, ok := (*memo)[towelDesign]; ok {
		return possible
	}

	if len(towelDesign) == 0 {
		(*memo)[towelDesign] = true
		return true
	}

	for _, stripes := range available {
		if strings.HasPrefix(towelDesign, stripes) {
			if canMakeTowel(towelDesign[len(stripes):], available, memo) {
				(*memo)[towelDesign] = true
				return true
			}
		}
	}

	(*memo)[towelDesign] = false
	return false
}

func waysToMakeTowel(towelDesign string, available []string, memo *map[string]int) int {
	if ways, ok := (*memo)[towelDesign]; ok {
		return ways
	}

	if len(towelDesign) == 0 {
		return 1
	}

	ways := 0
	for _, stripes := range available {
		if strings.HasPrefix(towelDesign, stripes) {
			ways += waysToMakeTowel(towelDesign[len(stripes):], available, memo)
		}
	}

	(*memo)[towelDesign] = ways
	return ways
}
