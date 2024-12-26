package day23

import (
	"aoc2024/utils"
	"log"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := utils.ReadInput("example.txt", 23)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := 7
	result := part1(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

func TestPart2(t *testing.T) {
	input, err := utils.ReadInput("example.txt", 23)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}
	expected := "co,de,ka,ta"
	result := part2(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %s, got %s\n"), expected, result)
	}

	// TestPls(input)
}
