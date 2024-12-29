package day24

import (
	"aoc2024/utils"
	"log"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := utils.ReadInputRaw("example.txt", 24)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := 2024
	result := part1(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

func TestPart2(t *testing.T) {
	input, err := utils.ReadInputRaw("example.txt", 24)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := "ffh,mjb,tgd,wpb,z02,z03,z05,z06,z07,z08,z10,z11"
	result := part2(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %s, got %s\n"), expected, result)
	}
}

