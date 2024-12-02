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
	input, err := utils.ReadInput("day02/input.txt")
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

		reversed := make([]int, len(ints))
		for i, n := range ints {
			reversedIndex := len(ints) - 1 - i
			reversed[reversedIndex] = n
		}

		// teeeechnically O(n) right??
		if isSafeEnough(ints) || isSafeEnough(reversed) {
			safeCount++
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

func isSafeEnough(nums []int) bool {
	var prevDiff int
	tolerable := true
	removedLast := false

	for i := 1; i < len(nums); i++ {
		prevIndex := i - 1
		if removedLast {
			prevIndex = i - 2
		}

		removedLast = false
		diff := nums[i] - nums[prevIndex]
		absDiff := utils.Abs(diff)

		if (prevIndex > 0 && diff*prevDiff <= 0) || absDiff > 3 || absDiff == 0 {
			if !tolerable {
				return false
			}

			tolerable = false
			removedLast = true
		} else {
			// ensure prevDiff will not be the removed element next iteration
			prevDiff = diff
		}
	}

	return true
}
