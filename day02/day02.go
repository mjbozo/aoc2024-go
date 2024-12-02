package day02

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

func Run() {
	input, err := utils.ReadInput("day02/day02_input.txt")
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
	safeCount := 0

	for _, line := range lines {
		nums := strings.Split(line, " ")
		ints := utils.Map(nums, func(x string) int {
			number, _ := strconv.Atoi(x)
			return number
		})

		if isSafe(ints) {
			safeCount++
		}
	}

	return safeCount
}

func part2(lines []string) int {
	safeCount := 0

	for _, line := range lines {
		nums := strings.Split(line, " ")
		ints := utils.Map(nums, func(x string) int {
			number, _ := strconv.Atoi(x)
			return number
		})

		if isSafe(ints) {
			safeCount++
		} else {
			// just brute force??
			for removeIndex := 0; removeIndex < len(ints); removeIndex++ {
				newInts := remove(ints, removeIndex)
				if isSafe(newInts) {
					safeCount++
					break
				}
			}
		}
	}

	return safeCount
}

func isSafe(nums []int) bool {
	var prevDiff int
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if i > 1 && diff*prevDiff <= 0 {
			return false
		}

		absDiff := math.Abs(float64(diff))
		if absDiff > 3 || absDiff == 0 {
			return false
		}

		prevDiff = diff
	}

	return true
}

func remove(s []int, i int) []int {
	removed := make([]int, 0)
	removed = append(removed, s[:i]...)
	removed = append(removed, s[i+1:]...)
	return removed
}

