package daybreaker

import (
	"aoc2024/utils"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

type DaybreakError struct {
	msg string
}

func (e *DaybreakError) Error() string {
	return e.msg
}

func Create(day string) error {
	dayNum := strings.TrimPrefix(day, "day")
	if len(dayNum) == 1 {
		// then i forgot to prepend day 1-9 with a zero
		day = fmt.Sprintf("day0%s", dayNum)
	}

	// validate directory does not already exist
	info, err := os.Stat(day)
	if err != nil && errors.Is(err, fs.ErrExist) {
		return err
	}

	if info != nil {
		return &DaybreakError{fmt.Sprintf("%s already exists", day)}
	}

	// create new directory with argument name from base directory
	err = os.Mkdir(day, 0750)
	if err != nil && !errors.Is(err, fs.ErrExist) {
		return err
	}

	// create `dayX.go`, `dayX_test.go`, `dayX_input.go`, `dayX_example1.txt` and `dayX_example2.txt` files in new directory
	err = os.WriteFile(fmt.Sprintf("%s/%s.go", day, day), []byte(fmt.Sprintf(`package %s

import (
	"aoc2024/utils"
	"fmt"
	"log"
)

func Run() {
	input, err := utils.ReadInput("%s/%s_input.txt")
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	fmt.Printf("Part 1: %%d\n", part1(input))
	fmt.Printf("Part 2: %%d\n", part2(input))
}

func part1(lines []string) int {
	return 0
}

func part2(lines []string) int {
	return 0
}`, day, day, day)), 0660)

	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s_test.go", day, day), []byte(fmt.Sprintf(`package %s

import (
	"aoc2024/utils"
	"log"
	"testing"
)

func TestPart1(t *testing.T) {
    input, err := utils.ReadInput("%s_example1.txt")
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

    expected := 0
	result := part1(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %%d, got %%d\n"), expected, result)
	}
}

func TestPart2(t *testing.T) {
    input, err := utils.ReadInput("%s_example2.txt")
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

    expected := 0
	result := part2(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %%d, got %%d\n"), expected, result)
	}
}`, day, day, day)), 0660)
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s_input.txt", day, day), []byte(""), 0660)
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s_example1.txt", day, day), []byte(""), 0660)
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s_example2.txt", day, day), []byte(""), 0660)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", utils.Green(fmt.Sprintf("\tSuccessfully created %s files", day)))
	return nil
}
