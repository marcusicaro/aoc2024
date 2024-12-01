package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getFileContent() ([]string, error) {
	file, err := os.Open("day1.txt")
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var output []string

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output, nil
}

func splitStringArrayInLeftAndRightArrays(strArr []string) ([]int, []int, error) {
	left := []int{}
	right := []int{}

	for _, str := range strArr {
		input := strings.Fields(str)
		leftNum, err := strconv.Atoi(input[0])
		if err != nil {
			return []int{}, []int{}, err
		}
		rightNum, err := strconv.Atoi(input[1])
		if err != nil {
			return []int{}, []int{}, err
		}
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	return left, right, nil

}

func compareToGetTheDifference(num1, num2 int) int {
	return int(math.Abs(float64(num1 - num2)))
}

func handleNumberValues(leftArr, rightArr []int) int {
	alreadyAppearedNums := make(map[int]int)
	sum := 0

	for _, numLeftArr := range leftArr {
		_, exists := alreadyAppearedNums[numLeftArr]

		if !exists {
			count := 0
			for _, numRightArr := range rightArr {
				if numLeftArr == numRightArr {
					count++
				}
			}
			alreadyAppearedNums[numLeftArr] = count * numLeftArr
		}

		sum += alreadyAppearedNums[numLeftArr]
	}

	return sum
}

func main() {
	filecontent, err := getFileContent()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	leftArr, rightArr, err := splitStringArrayInLeftAndRightArrays(filecontent)
	if err != nil {
		fmt.Println("Error converting string array to int array:", err)
		return
	}

	sort.Ints(leftArr)
	sort.Ints(rightArr)

	sumTotalDifference := 0

	for i := 0; i < len(leftArr); i++ {
		sumTotalDifference += compareToGetTheDifference(leftArr[i], rightArr[i])
	}

	fmt.Println("Sum of the total difference between the two arrays is:", sumTotalDifference)

	sumTotalValues := handleNumberValues(leftArr, rightArr)

	fmt.Println("Sum of the total values is:", sumTotalValues)

}
