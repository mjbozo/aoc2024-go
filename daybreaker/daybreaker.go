package daybreaker

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func Create(day string) error {
	fmt.Printf("Daybreaking %s\n", day)

	// create new directory with argument name from base directory
	err := os.Mkdir(day, 0750)
	if err != nil && !errors.Is(err, fs.ErrExist) {
		return err
	}

	// create `dayX.go`, `dayX_test.go`, `dayX_input.go` and `dayX_example.txt` files in new directory
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
	fmt.Println(lines)
	return 0
}

func part2(lines []string) int {
	fmt.Println(lines)
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

var input []string
var err error

func TestMain(m *testing.M) {
	input, err = utils.ReadInput("%s_example.txt")
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}
	m.Run()
}

func TestPart1(t *testing.T) {
    expected := 0
	result := part1(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %%d, got %%d\n"), expected, result)
	}
}

func TestPart2(t *testing.T) {
    expected := 0
	result := part2(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %%d, got %%d\n"), expected, result)
	}
}`, day, day)), 0660)
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s_input.txt", day, day), []byte(""), 0660)
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s_example.txt", day, day), []byte(""), 0660)
	if err != nil {
		return err
	}

	return nil
}
