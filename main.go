package main

import (
	"aoc2024/day1"
	"aoc2024/daybreaker"
	"aoc2024/utils"
	"fmt"
	"log"
	"os"
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
	switch day {
	case "day1":
		day1.Run()
	default:
		fmt.Printf("%s not completed yet\n", day)
	}
}
