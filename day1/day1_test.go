package day1

import (
	"aoc2024/utils"
	"log"
	"testing"
)

var input []string
var err error

func TestMain(m *testing.M) {
	input, err = utils.ReadInput("day1_example.txt")
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}
	m.Run()
}

func TestPart1(t *testing.T) {
	expected := 0
	result := part1(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 0
	result := part2(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}
