package day01

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
)

func Run() {
	input, err := utils.ReadInput("day01/input.txt")
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

func part1(lines []string) int {
	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for _, pair := range lines {
		nums := strings.Split(pair, "   ")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	slices.SortFunc(list1, func(a, b int) int { return a - b })
	slices.SortFunc(list2, func(a, b int) int { return a - b })

	totalDiff := 0
	for i := range len(list1) {
		totalDiff += int(math.Abs(float64(list1[i]) - float64(list2[i])))
	}

	return totalDiff
}

func part2(lines []string) int {
	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for _, pair := range lines {
		nums := strings.Split(pair, "   ")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	freq := make(map[int]int)
	for _, n := range list2 {
		if val, ok := freq[n]; !ok {
			freq[n] = 1
		} else {
			freq[n] = val + 1
		}
	}

	similarityScore := 0
	for _, n := range list1 {
		if val, ok := freq[n]; ok {
			similarityScore += n * val
		}
	}

	return similarityScore
}
