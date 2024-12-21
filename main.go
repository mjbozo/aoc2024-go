package main

import (
	"aoc2024/day01"
	"aoc2024/day02"
	"aoc2024/day03"
	"aoc2024/day04"
	"aoc2024/day05"
	"aoc2024/day06"
	"aoc2024/day07"
	"aoc2024/day08"
	"aoc2024/day09"
	"aoc2024/day10"
	"aoc2024/day11"
	"aoc2024/day12"
	"aoc2024/day13"
	"aoc2024/day14"
	"aoc2024/day15"
	"aoc2024/day16"
	"aoc2024/day17"
	"aoc2024/day18"
	"aoc2024/day19"
	"aoc2024/day20"
	/*
		"aoc2024/day21"
		"aoc2024/day22"
		"aoc2024/day23"
		"aoc2024/day24"
		"aoc2024/day25"
	*/
	"aoc2024/daybreaker"
	"aoc2024/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatalln(utils.Red("[ ERR ] Not enough arguments. Provide the day you want to run"))
	}

	if args[1] == "-c" {
		if len(args) < 3 {
			log.Fatalln(utils.Red("[ ERR ] Not enough arguments. Provide the day you want to create"))
		}

		err := daybreaker.Create(args[2])
		if err != nil {
			log.Fatalln(utils.Red(fmt.Sprintf("[ ERR ] Daybreaker failed: %s", err.Error())))
		}

		return
	}

	day := args[1]

	if day == "all" {
		runAll()
		return
	}

	dayNum := strings.TrimPrefix(day, "day")
	if len(dayNum) == 1 {
		// then i forgot to prepend day 1-9 with a zero
		day = fmt.Sprintf("day0%s", dayNum)
	}

	switch day {
	case "day01":
		day01.Run()
	case "day02":
		day02.Run()
	case "day03":
		day03.Run()
	case "day04":
		day04.Run()
	case "day05":
		day05.Run()
	case "day06":
		day06.Run()
	case "day07":
		day07.Run()
	case "day08":
		day08.Run()
	case "day09":
		day09.Run()
	case "day10":
		day10.Run()
	case "day11":
		day11.Run()
	case "day12":
		day12.Run()
	case "day13":
		day13.Run()
	case "day14":
		day14.Run()
	case "day15":
		day15.Run()
	case "day16":
		day16.Run()
	case "day17":
		day17.Run()
	case "day18":
		day18.Run()
	case "day19":
		day19.Run()
	case "day20":
		day20.Run()
		/*
			case "day21":
				day21.Run()
			case "day22":
				day22.Run()
			case "day23":
				day23.Run()
			case "day24":
				day24.Run()
			case "day25":
				day25.Run()
		*/
	default:
		fmt.Printf("%s not completed yet\n", day)
	}
}

func runAll() {
	fmt.Println("DAY 01")
	day01.Run()
	fmt.Println("\nDAY 02")
	day02.Run()
	fmt.Println("\nDAY 03")
	day03.Run()
	fmt.Println("\nDAY 04")
	day04.Run()
	fmt.Println("\nDAY 05")
	day05.Run()
	fmt.Println("\nDAY 06")
	day06.Run()
	fmt.Println("\nDAY 07")
	day07.Run()
	fmt.Println("\nDAY 08")
	day08.Run()
	fmt.Println("\nDAY 09")
	day09.Run()
	fmt.Println("\nDAY 10")
	day10.Run()
	fmt.Println("\nDAY 11")
	day11.Run()
	fmt.Println("\nDAY 12")
	day12.Run()
	fmt.Println("\nDAY 13")
	day13.Run()
	fmt.Println("\nDAY 14")
	day14.Run()
	fmt.Println("\nDAY 15")
	day15.Run()
	fmt.Println("\nDAY 16")
	day16.Run()
	fmt.Println("\nDAY 17")
	day17.Run()
	fmt.Println("\nDAY 18")
	day18.Run()
	fmt.Println("\nDAY 19")
	day19.Run()
	/*
	   fmt.Println("\nDAY 20")
	   day20.Run()
	   fmt.Println("\nDAY 21")
	   day21.Run()
	   fmt.Println("\nDAY 22")
	   day22.Run()
	   fmt.Println("\nDAY 23")
	   day23.Run()
	   fmt.Println("\nDAY 24")
	   day24.Run()
	   fmt.Println("\nDAY 25")
	   day25.Run()
	*/
}
