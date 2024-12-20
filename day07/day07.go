package day07

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func Run() {
	input, err := utils.ReadInput("day07/input.txt", 7)
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
	sum := 0

	c := make(chan int, len(lines))
	for _, line := range lines {
		go func(line string) {
			parts := strings.Split(line, ": ")
			testValue, _ := strconv.Atoi(parts[0])
			calibrationNums := utils.Map(strings.Split(parts[1], " "), func(s string) int {
				num, _ := strconv.Atoi(s)
				return num
			})

			if testCalibration(testValue, calibrationNums) {
				c <- testValue
			} else {
				c <- 0
			}
		}(line)
	}

	for range len(lines) {
		sum += <-c
	}

	return sum
}

func part2(lines []string) int {
	sum := 0

	c := make(chan int, len(lines))
	for _, line := range lines {
		go func(line string) {
			parts := strings.Split(line, ": ")
			testValue, _ := strconv.Atoi(parts[0])
			calibrationNums := utils.Map(strings.Split(parts[1], " "), func(s string) int {
				num, _ := strconv.Atoi(s)
				return num
			})

			if testCalibrationWithConcat(testValue, calibrationNums) {
				c <- testValue
			} else {
				c <- 0
			}
		}(line)
	}

	for range len(lines) {
		sum += <-c
	}

	return sum
}

func testCalibration(testValue int, calibrationNums []int) bool {
	first := calibrationNums[0]
	if first > testValue {
		return false
	}

	second := calibrationNums[1]

	addition := first + second
	multiply := first * second

	if len(calibrationNums) == 2 {
		if addition == testValue || multiply == testValue {
			return true
		} else {
			return false
		}
	}

	addSlice := []int{addition}
	addSlice = append(addSlice, calibrationNums[2:]...)

	multiplySlice := []int{multiply}
	multiplySlice = append(multiplySlice, calibrationNums[2:]...)

	return testCalibration(testValue, addSlice) || testCalibration(testValue, multiplySlice)
}

func testCalibrationWithConcat(testValue int, calibrationNums []int) bool {
	first := calibrationNums[0]
	if first > testValue {
		return false
	}

	second := calibrationNums[1]

	addition := first + second
	multiply := first * second
	concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, second))

	if len(calibrationNums) == 2 {
		if addition == testValue || multiply == testValue || concat == testValue {
			return true
		} else {
			return false
		}
	}

	addSlice := []int{addition}
	addSlice = append(addSlice, calibrationNums[2:]...)

	multiplySlice := []int{multiply}
	multiplySlice = append(multiplySlice, calibrationNums[2:]...)

	concatSlice := []int{concat}
	concatSlice = append(concatSlice, calibrationNums[2:]...)

	return testCalibrationWithConcat(testValue, addSlice) ||
		testCalibrationWithConcat(testValue, multiplySlice) ||
		testCalibrationWithConcat(testValue, concatSlice)
}
