package main

import (
	"aoc2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func handleFileContext(file string) (string, string) {
	splitFile := strings.Split(file, "\n\n")
	return splitFile[0], splitFile[1]
}

func handleRules(rules string) map[int][]int {
	stringRules := strings.Split(rules, "\n")
	rulesMap := make(map[int][]int)

	for _, rule := range stringRules {
		splitRule := strings.Split(rule, "|")
		key, err := strconv.Atoi(splitRule[0])
		if err != nil {
			panic(err)
		}
		value, err := strconv.Atoi(splitRule[1])
		if err != nil {
			panic(err)
		}
		if !utils.Contains(rulesMap[key], value) {
			rulesMap[key] = append(rulesMap[key], value)
		}

	}

	return rulesMap
}

func convertNumbersStringIntoNumberSequece(numbers string) [][]int {
	numbersSequenceArray := strings.Split(numbers, "\n")
	numbersSequence := [][]int{}

	for _, numberSequence := range numbersSequenceArray {
		numbers := strings.Split(numberSequence, ",")
		numbersInt := []int{}
		for _, number := range numbers {
			numberInt, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			numbersInt = append(numbersInt, numberInt)
		}
		numbersSequence = append(numbersSequence, numbersInt)
	}

	return numbersSequence
}

func validateNumbersWithRules(rules map[int][]int, numbers []int) bool {
	currentNumberMandatoryPreviousNumbers := []int{}
	for numsArrayKey, number := range numbers {

		for key, values := range rules {
			if utils.Contains(values, number) {
				currentNumberMandatoryPreviousNumbers = append(currentNumberMandatoryPreviousNumbers, key)
			}
		}

		fromNumberBeyond := numbers[numsArrayKey:]
		for _, num := range currentNumberMandatoryPreviousNumbers {
			if utils.Contains(fromNumberBeyond, num) {
				return false
			}
		}

	}

	return true
}

func solveInvalidSequences(rules map[int][]int, numbers *[]int) {
	for range *numbers {
		for numsArrayKey, number := range *numbers {
			if numsArrayKey+1 < len(*numbers) && utils.Contains(rules[(*numbers)[numsArrayKey+1]], number) {
				(*numbers)[numsArrayKey+1], (*numbers)[numsArrayKey] = (*numbers)[numsArrayKey], (*numbers)[numsArrayKey+1]
			}

		}
	}

}

func main() {
	file, err := utils.GetFileContentInSingleString("day5.txt")
	if err != nil {
		panic(err)
	}

	rules, numbers := handleFileContext(file)

	rulesMap := handleRules(rules)
	numbersSequence := convertNumbersStringIntoNumberSequece(numbers)
	validSequences := [][]int{}
	invalidSequences := [][]int{}

	for _, numbers := range numbersSequence {
		if validateNumbersWithRules(rulesMap, numbers) {
			validSequences = append(validSequences, numbers)
		} else {
			invalidSequences = append(invalidSequences, numbers)
		}
	}

	result := 0
	for _, seq := range validSequences {
		middleIndex := len(seq) / 2
		result += seq[middleIndex]
	}
	// fmt.Print(result)

	for _, seq := range invalidSequences {
		solveInvalidSequences(rulesMap, &seq)
	}

	resultPt2 := 0
	for _, seq := range invalidSequences {
		middleIndex := len(seq) / 2
		resultPt2 += seq[middleIndex]
	}
	fmt.Print(resultPt2)

}
