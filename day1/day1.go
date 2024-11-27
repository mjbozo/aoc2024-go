package day1

import (
	"aoc2024/utils"
	"fmt"
	"log"
)

func Run() {
	input, err := utils.ReadInput("day1/day1_input.txt")
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(lines []string) int {
	fmt.Println(lines)
	return 0
}

func part2(lines []string) int {
	fmt.Println(lines)
	return 0
}
