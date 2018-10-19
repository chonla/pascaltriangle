package main

import (
	"fmt"

	"github.com/chonla/pascal-triangle/pascaltriangle"
)

func main() {
	lines := 10

	printPascalRecursive(lines)
	printPascalLoop(lines)
}

func printPascalRecursive(lines int) {
	for i := 1; i <= lines; i++ {
		line := pascaltriangle.PascalLineMemo(i)
		for j, n := 0, len(line); j < n; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(line[j])
		}
		fmt.Println()
	}
}

func printPascalLoop(lines int) {
	arr := [][]int{}

	for line := 0; line < lines; line++ {
		arr = append(arr, []int{})
		for i := 0; i <= line; i++ {
			if i == 0 || i == line {
				arr[line] = append(arr[line], 1)
			} else {
				arr[line] = append(arr[line], arr[line-1][i-1]+arr[line-1][i])
			}
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(arr[line][i])
		}
		fmt.Println()
	}
}
