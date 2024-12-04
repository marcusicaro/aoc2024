package main

import (
	"aoc2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	filecontent, err := utils.GetFileContent("day3.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	content := strings.Join(filecontent, "\n")

	pattern := `mul\(\d+,\d+\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(content, -1)

	patternNums := `\d+`
	reNums := regexp.MustCompile(patternNums)

	result := 0
	consider := true

	for _, match := range matches {
		if match == "do()" {
			consider = true
			continue
		} else if match == "don't()" {
			consider = false
			continue
		}
		if !consider {
			continue
		}
		matchesNums := reNums.FindAllString(match, -1)
		num1, err1 := strconv.Atoi(matchesNums[0])
		num2, err2 := strconv.Atoi(matchesNums[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting string to int")
			return
		}

		result += num1 * num2
	}

	fmt.Println(result)
}
