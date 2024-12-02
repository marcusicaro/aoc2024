package main

import (
	"aoc2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func countValidLines(strArr []string) (int, error) {
	count := 0
	for _, str := range strArr {
		line := strings.Fields(str)
		validLine, err := isLineValid(line)
		if err != nil {
			return 0, err
		}

		if validLine {
			count++
		}
	}
	return count, nil
}

func isLineValid(input []string) (bool, error) {
	var nums []int

	for _, strNum := range input {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			return false, err
		}
		nums = append(nums, num)
	}

	for i := range nums {
		numsCopy := utils.RemoveElement(nums, i)

		if validateLine(numsCopy) {
			return true, nil
		}
	}

	return false, nil
}

func validateLine(line []int) bool {
	var order string
	if line[0] < line[1] {
		order = "ascending"
	} else if line[0] > line[1] {
		order = "descending"
	} else {
		return false
	}

	for index := 1; index < len(line); index++ {
		if order == "ascending" {
			diff := line[index] - line[index-1]
			if diff >= 1 && diff <= 3 {
				continue
			} else {
				return false
			}
		} else {
			diff := line[index-1] - line[index]
			if diff >= 1 && diff <= 3 {
				continue
			} else {
				return false
			}
		}
	}

	return true
}

func main() {
	filecontent, err := utils.GetFileContent("day2.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	count, err := countValidLines(filecontent)
	if err != nil {
		fmt.Println("Error counting lines:", err)
		return
	}

	fmt.Println("Total valid lines:", count)

}
