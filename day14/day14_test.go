package day14

import (
	"aoc2024/utils"
	"log"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := utils.ReadInput("example.txt", 14)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := 12
	result := part1(input, 11, 7)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

func TestPart2(t *testing.T) {
	input, err := utils.ReadInput("example.txt", 14)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := 0
	result := part2(input, 11, 7)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}
