package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

const FILEPATH string = "puzzle_input.txt"
const CALIBRATION_METHOD = 2

func main() {
	f, err := os.Open(FILEPATH)
	if err != nil {
		log.Fatal("Unable to open input file", err)
	}
	defer f.Close()

	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := string(scanner.Text())
		result := 0

		switch CALIBRATION_METHOD {
		case 1:
			result = part1_method(line)
		case 2:
			result = part2_method(line)
		}
		total += result

	}
	fmt.Printf("Answer is: %d\n", total)
}

func part1_method(line string) int {
	length := len(line)
	result := 0

	for i := 0; i < length; i++ {
		if unicode.IsDigit(rune(line[i])) {
			startDigit, _ := strconv.Atoi(string(line[i]))
			result = startDigit * 10
			break
		}
	}

	for j := length - 1; j >= 0; j-- {
		if unicode.IsDigit(rune(line[j])) {
			endDigit, _ := strconv.Atoi(string(line[j]))
			result += endDigit
			break
		}
	}

	return result
}

func part2_method(line string) int {
	result := 0
	reg := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|[0-9])")
	regRev := regexp.MustCompile("(enin|thgie|neves|xis|evif|ruof|eerht|owt|eno|[0-9])")

	firstMatch := reg.Find([]byte(line))
	secondMatch := regRev.Find([]byte(reverse(line)))

	// Take first and last matches
	firstDigit := convertToDigit(string(firstMatch))
	secondDigit := convertToDigit(reverse(string(secondMatch)))

	result = firstDigit*10 + secondDigit

	return result
}

// Takes a string digit or spelled out number and returns an integer
// e.g "1" : 1
//
//	"one" : 1
func convertToDigit(input string) int {
	// Digit as string
	if unicode.IsDigit(rune(input[0])) {
		result, _ := strconv.Atoi(input)
		return result
	}

	// Digit as word
	number_names := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	result := number_names[input]
	return result
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
