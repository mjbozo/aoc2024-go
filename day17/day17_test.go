package day17

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

	expected := "4,6,3,5,6,3,5,2,1,0"
	result := part1(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %s, got %s\n"), expected, result)
	}
}

func TestPart2(t *testing.T) {
	input, err := utils.ReadInputRaw("example2.txt")
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := 366332 // output for MY program only
	result := part2(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

