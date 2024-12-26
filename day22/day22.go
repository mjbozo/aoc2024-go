package day22

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
	input, err := utils.ReadInput("day22/input.txt", 22)
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
	for _, line := range lines {
		secretNumber, _ := strconv.Atoi(line)

		for range 2000 {
			step1 := secretNumber * 64
			secretNumber = mix(secretNumber, step1)
			secretNumber = prune(secretNumber)

			step2 := int(math.Floor(float64(secretNumber) / 32.0))
			secretNumber = mix(secretNumber, step2)
			secretNumber = prune(secretNumber)

			step3 := secretNumber * 2048
			secretNumber = mix(secretNumber, step3)
			secretNumber = prune(secretNumber)
		}

		sum += secretNumber
	}

	return sum
}

func part2(lines []string) int {
	diffsToBananas := make(map[string]int)

	for _, line := range lines {
		secretNumber, _ := strconv.Atoi(line)
		prevLastDigit, _ := strconv.Atoi(line[len(line)-1:])
		diffs := make([]string, 0)
		lastDigits := make([]int, 0)

		for range 2000 {
			step1 := secretNumber * 64
			secretNumber = mix(secretNumber, step1)
			secretNumber = prune(secretNumber)

			step2 := int(math.Floor(float64(secretNumber) / 32.0))
			secretNumber = mix(secretNumber, step2)
			secretNumber = prune(secretNumber)

			step3 := secretNumber * 2048
			secretNumber = mix(secretNumber, step3)
			secretNumber = prune(secretNumber)

			secretNumberString := fmt.Sprintf("%d", secretNumber)
			secretNumberLastDigit, _ := strconv.Atoi(secretNumberString[len(secretNumberString)-1:])
			lastDigits = append(lastDigits, secretNumberLastDigit)

			diffs = append(diffs, fmt.Sprintf("%d", secretNumberLastDigit-prevLastDigit))

			prevLastDigit = secretNumberLastDigit
		}

		seen := make(utils.HashSet[string])

		for i := 0; i < len(diffs)-3; i++ {
			sequence := strings.Join(diffs[i:i+4], ",")
			if !seen.Contains(sequence) {
				seen.Insert(sequence)
				numBananas := lastDigits[i+3]
				diffsToBananas[sequence] += numBananas
			}
		}
	}

	bananas := 0
	for _, value := range diffsToBananas {
		if value > bananas {
			bananas = value
		}
	}

	return bananas
}

func mix(secretNumber int, mixer int) int {
	mix := mixer ^ secretNumber
	return mix
}

func prune(secretNumber int) int {
	return secretNumber % 16777216
}
