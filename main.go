package main

import "fmt"

func main() {
	grid := [][]int{[]int{1, 0}, []int{0, 1}}
	// row := []int{0, 0}

	// for i := 0; i < 2; i++ {
	// 	grid = append(grid, row)
	// }

	for _, r := range grid {
		fmt.Printf("%v\n", r)
	}
}

func isRowValid(grid [][]int, row int, num int) bool {
	for col := 0; col < 9; col++ {
		if grid[row][col] == num {
			return false
		}
	}
}
