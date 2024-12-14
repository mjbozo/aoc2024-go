package day11

import (
	"aoc2024/utils"
	"log"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := utils.ReadInputRaw("example.txt")
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	var expected uint = 55312
	result := part1(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

func TestPart2(t *testing.T) {
	input, err := utils.ReadInputRaw("example.txt")
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	var expected uint = 22
	result := part2(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

