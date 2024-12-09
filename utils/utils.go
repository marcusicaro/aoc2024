package utils

import (
	"bufio"
	"os"
	"strings"
)

func GetFileContent(filename string) ([]string, error) {
	file, err := os.Open(filename)
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

func GetFileContentInSingleString(filename string) (string, error) {

	file, err := GetFileContent(filename)
	if err != nil {
		return "", err
	}

	output := strings.Join(file, "\n")

	return output, err
}

func RemoveElement(slice []int, index int) []int {
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)
	copySlice = append(copySlice[:index], copySlice[index+1:]...)
	return copySlice
}

func Contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
