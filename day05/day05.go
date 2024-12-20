package day05

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func Run() {
	input, err := utils.ReadInputRaw("day05/input.txt", 5)
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
	rules := parts[0]
	pages := parts[1]

	ruleMap := buildOrderingMap(rules)

	goodLines := make([]string, 0)
	for _, line := range strings.Split(pages, "\n") {
		nums := strings.Split(line, ",")
		if isGoodLine(nums, ruleMap) {
			goodLines = append(goodLines, line)
		}
	}

	sum := 0
	for _, line := range goodLines {
		nums := strings.Split(line, ",")
		middleIndex := (len(nums) - 1) / 2
		x, _ := strconv.Atoi(nums[middleIndex])
		sum += x
	}

	return sum
}

func part2(lines string) int {
	parts := strings.Split(lines, "\n\n")
	rules := parts[0]
	pages := parts[1]

	ruleMap := buildOrderingMap(rules)
	badLines := make([]string, 0)

	for _, line := range strings.Split(pages, "\n") {
		nums := strings.Split(line, ",")
		if !isGoodLine(nums, ruleMap) {
			badLines = append(badLines, line)
		}
	}

	sum := 0
	for _, line := range badLines {
		nums := strings.Split(line, ",")

		for i := 0; i < len(nums)-1; i++ {
			for j := i + 1; j < len(nums); j++ {
				if allowedAfter, ok := ruleMap[nums[j]]; ok && allowedAfter.Contains(nums[i]) {
					nums[i], nums[j] = nums[j], nums[i]
				}
			}
		}

		middleIndex := (len(nums) - 1) / 2
		x, _ := strconv.Atoi(nums[middleIndex])
		sum += x
	}

	return sum
}

func buildOrderingMap(rules string) map[string]utils.HashSet[string] {
	ruleMap := make(map[string]utils.HashSet[string])
	for _, rule := range strings.Split(rules, "\n") {
		nums := strings.Split(rule, "|")
		if list, ok := ruleMap[nums[0]]; ok {
			list.Insert(nums[1])
			ruleMap[nums[0]] = list
		} else {
			newList := make(utils.HashSet[string])
			newList.Insert(nums[1])
			ruleMap[nums[0]] = newList
		}
	}

	return ruleMap
}

func isGoodLine(nums []string, ruleMap map[string]utils.HashSet[string]) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if allowedAfter, ok := ruleMap[nums[j]]; ok && allowedAfter.Contains(nums[i]) {
				return false
			}
		}
	}

	return true
}
