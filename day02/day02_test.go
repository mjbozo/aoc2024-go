package day02

import (
	"aoc2024/utils"
	"log"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := utils.ReadInput("day02_example1.txt")
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := 2
	result := part1(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

func TestPart2(t *testing.T) {
	input, err := utils.ReadInput("day02_example2.txt")
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := 4
	result := part2(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

