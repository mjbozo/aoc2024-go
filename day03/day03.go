package day03

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

func Run() {
	input, err := utils.ReadInput("day03/input.txt")
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
	mulRegex := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	digitRegex := regexp.MustCompile(`[0-9]+`)

	for _, line := range lines {
		matches := mulRegex.FindAll([]byte(line), -1)

		for _, match := range matches {
			digits := digitRegex.FindAll(match, -1)
			if len(digits) != 2 {
				panic(fmt.Sprintf("Got more than 2 digits in mul operation: %s", match))
			}

			num1, _ := strconv.Atoi(string(digits[0]))
			num2, _ := strconv.Atoi(string(digits[1]))
			sum += num1 * num2
		}
	}

	return sum
}

func part2(lines []string) int {
	sum := 0
	enabled := true
	mulRegex := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	digitRegex := regexp.MustCompile(`[0-9]+`)

	for _, line := range lines {
		matches := mulRegex.FindAll([]byte(line), -1)

		for _, match := range matches {
			switch string(match) {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				if enabled {
					digits := digitRegex.FindAll(match, -1)
					if len(digits) != 2 {
						panic(fmt.Sprintf("Got more than 2 digits in mul operation: %s", match))
					}

					num1, _ := strconv.Atoi(string(digits[0]))
					num2, _ := strconv.Atoi(string(digits[1]))
					sum += num1 * num2
				}
			}
		}
	}

	return sum
}

