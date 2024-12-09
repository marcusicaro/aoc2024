package main

import (
	"aoc2024/utils"
)

func findXmasWord(filecontent []string) (int, int) {
	count := 0
	countPt2 := 0

	// for i, line := range filecontent {
	// 	start := 0
	// 	for {
	// 		index := strings.Index(line[start:], "X")
	// 		absoluteIndex := start + index
	// 		if index == -1 {
	// 			break
	// 		}

	// 		if absoluteIndex+3 < len(line) && filecontent[i][absoluteIndex+1] == 'M' && filecontent[i][absoluteIndex+2] == 'A' && filecontent[i][absoluteIndex+3] == 'S' {
	// 			count++
	// 		}
	// 		if absoluteIndex-3 >= 0 && filecontent[i][absoluteIndex-1] == 'M' && filecontent[i][absoluteIndex-2] == 'A' && filecontent[i][absoluteIndex-3] == 'S' {
	// 			count++
	// 		}
	// 		if i+3 < len(filecontent) && filecontent[i+1][absoluteIndex] == 'M' && filecontent[i+2][absoluteIndex] == 'A' && filecontent[i+3][absoluteIndex] == 'S' {
	// 			count++
	// 		}
	// 		if i-3 >= 0 && filecontent[i-1][absoluteIndex] == 'M' && filecontent[i-2][absoluteIndex] == 'A' && filecontent[i-3][absoluteIndex] == 'S' {
	// 			count++
	// 		}
	// 		if i+3 < len(filecontent) && absoluteIndex+3 < len(line) && filecontent[i+1][absoluteIndex+1] == 'M' && filecontent[i+2][absoluteIndex+2] == 'A' && filecontent[i+3][absoluteIndex+3] == 'S' {
	// 			count++
	// 		}
	// 		if i-3 >= 0 && absoluteIndex-3 >= 0 && filecontent[i-1][absoluteIndex-1] == 'M' && filecontent[i-2][absoluteIndex-2] == 'A' && filecontent[i-3][absoluteIndex-3] == 'S' {
	// 			count++
	// 		}
	// 		if i+3 < len(filecontent) && absoluteIndex-3 >= 0 && filecontent[i+1][absoluteIndex-1] == 'M' && filecontent[i+2][absoluteIndex-2] == 'A' && filecontent[i+3][absoluteIndex-3] == 'S' {
	// 			count++
	// 		}
	// 		if i-3 >= 0 && absoluteIndex+3 < len(line) && filecontent[i-1][absoluteIndex+1] == 'M' && filecontent[i-2][absoluteIndex+2] == 'A' && filecontent[i-3][absoluteIndex+3] == 'S' {
	// 			count++
	// 		}

	// 		start += index + 1
	// 	}
	// }

	for i, line := range filecontent {
		if i+2 <= len(filecontent)-1 {
			for j := 0; j < len(line)-3; j++ {
				if filecontent[i][j] == 'M' && filecontent[i+1][j+1] == 'A' && filecontent[i+2][j+2] == 'S' && filecontent[i][j+2] == 'M' && filecontent[i+2][j] == 'S' {
					countPt2++
				}

				if filecontent[i][j] == 'S' && filecontent[i+1][j+1] == 'A' && filecontent[i+2][j+2] == 'M' && filecontent[i][j+2] == 'S' && filecontent[i+2][j] == 'M' {
					countPt2++
				}

				if filecontent[i][j] == 'M' && filecontent[i+1][j+1] == 'A' && filecontent[i+2][j+2] == 'S' && filecontent[i][j+2] == 'S' && filecontent[i+2][j] == 'M' {
					countPt2++
				}

				if filecontent[i][j] == 'S' && filecontent[i+1][j+1] == 'A' && filecontent[i+2][j+2] == 'M' && filecontent[i][j+2] == 'M' && filecontent[i+2][j] == 'S' {
					countPt2++
				}
			}
		}

	}

	return count, countPt2
}

func main() {
	filecontent, _ := utils.GetFileContent("day4.txt")

	_, pt2 := findXmasWord(filecontent)

	println(pt2)

}
