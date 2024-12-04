package main

import (
	"aoc2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const patternNums = `\d+`

var reNums = regexp.MustCompile(patternNums)

func main() {
	filecontent, err := utils.GetFileContent("day3.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	content := strings.Join(filecontent, "\n")

	pattern := `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(content, -1)

	resultPt1 := getPt1Result(matches)
	resultPt2 := getPt2Result(matches)

	fmt.Println("pt1: ")
	fmt.Println(resultPt1)
	fmt.Println("pt2: ")
	fmt.Println(resultPt2)
}

func converToNumbersAndMultiply(result *int, match string) {
	matchesNums := reNums.FindAllString(match, -1)
	num1, err1 := strconv.Atoi(matchesNums[0])
	num2, err2 := strconv.Atoi(matchesNums[1])

	if err1 != nil || err2 != nil {
		fmt.Println("Error converting string to int")
	}

	*result += num1 * num2
}

func getPt1Result(matches []string) int {
	resultPt1 := 0
	for _, match := range matches {
		if match == "do()" {
			continue
		} else if match == "don't()" {
			continue
		}
		converToNumbersAndMultiply(&resultPt1, match)
	}

	return resultPt1
}

func getPt2Result(matches []string) int {
	consider := true
	resultPt2 := 0
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
		converToNumbersAndMultiply(&resultPt2, match)
	}

	return resultPt2
}
