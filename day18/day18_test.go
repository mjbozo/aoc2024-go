package day18

import (
	"aoc2024/utils"
	"log"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := utils.ReadInput("example.txt", 18)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := 22
	result := part1(input, 7, 7, 12)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

func TestPart2(t *testing.T) {
	input, err := utils.ReadInput("example.txt", 18)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := "6,1"
	result := part2(input, 7, 7, 12)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

