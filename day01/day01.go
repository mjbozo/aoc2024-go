package day01

import (
	"aoc2024/utils"
	"fmt"
	"log"
)

func Run() {
	input, err := utils.ReadInput("day01/day01_input.txt")
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(lines []string) int {
	return 0
}

func part2(lines []string) int {
	return 0
}
