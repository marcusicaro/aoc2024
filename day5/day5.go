package main

import (
	"aoc2024/utils"
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
		rulesMap[key] = append(rulesMap[key], value)

		for _, v := range rulesMap[key] {
			if values, ok := rulesMap[v]; ok {
				rulesMap[key] = append(rulesMap[key], values...)
			}
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

		currentArray := numbers[numsArrayKey:]

		for _, currentNumberMandatoryPreviousNumber := range currentNumberMandatoryPreviousNumbers {
			if !utils.Contains(currentArray, currentNumberMandatoryPreviousNumber) {
				return false
			}
		}
	}

	return true
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

	for _, numbers := range numbersSequence {
		if validateNumbersWithRules(rulesMap, numbers) {
			validSequences = append(validSequences, numbers)
		}
	}

	println(len(validSequences))
}
