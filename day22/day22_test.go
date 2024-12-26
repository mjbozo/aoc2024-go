package day22

import (
	"aoc2024/utils"
	"log"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := utils.ReadInput("example.txt", 22)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := 37327623
	result := part1(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

func TestPart2(t *testing.T) {
	input, err := utils.ReadInput("example2.txt", 22)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := 23
	result := part2(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}
