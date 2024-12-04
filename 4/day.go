package main

import (
	"bufio"
	"fmt"
	"os"
)

//Functions for both

func readInput() [][]string {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	reader := bufio.NewReader(f)
	i := 0

	allLines := [][]string{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break //will break at the end of the file
		}
		allLines = append(allLines, []string{})
		for _, char := range line {
			allLines[i] = append(allLines[i], string(char))
		}
		i++

	}
	defer f.Close()
	return allLines
}

// /Functions for part 1
func searchUP(grid [][]string, i, j int) bool {
	return i >= 3 && grid[i][j] == "X" && grid[i-1][j] == "M" && grid[i-2][j] == "A" && grid[i-3][j] == "S"
}
func searchDOWN(grid [][]string, i, j int) bool {
	return i <= len(grid)-4 && grid[i][j] == "X" && grid[i+1][j] == "M" && grid[i+2][j] == "A" && grid[i+3][j] == "S"
}
func searchLEFT(grid [][]string, i, j int) bool {
	return j >= 3 && grid[i][j] == "X" && grid[i][j-1] == "M" && grid[i][j-2] == "A" && grid[i][j-3] == "S"
}
func searchRIGHT(grid [][]string, i, j int) bool {
	return j <= len(grid[0])-4 && grid[i][j] == "X" && grid[i][j+1] == "M" && grid[i][j+2] == "A" && grid[i][j+3] == "S"
}
func searchUPLEFT(grid [][]string, i, j int) bool {
	return i >= 3 && j >= 3 && grid[i-1][j-1] == "M" && grid[i-2][j-2] == "A" && grid[i-3][j-3] == "S"
}
func searchUPRIGHT(grid [][]string, i, j int) bool {
	return i >= 3 && j <= len(grid[0])-4 && grid[i-1][j+1] == "M" && grid[i-2][j+2] == "A" && grid[i-3][j+3] == "S"
}
func searchDOWNLEFT(grid [][]string, i, j int) bool {
	return i <= len(grid)-4 && j >= 3 && grid[i+1][j-1] == "M" && grid[i+2][j-2] == "A" && grid[i+3][j-3] == "S"
}
func searchDOWNRIGHT(grid [][]string, i, j int) bool {
	return i <= len(grid)-4 && j <= len(grid[0])-4 && grid[i+1][j+1] == "M" && grid[i+2][j+2] == "A" && grid[i+3][j+3] == "S"
}
func searchXMAS(grid [][]string) int {
	founds := 0
	for i, line := range grid {
		for j, char := range line {
			if char == "X" {
				for _, search := range []func([][]string, int, int) bool{searchUP, searchDOWN, searchLEFT, searchRIGHT, searchUPLEFT, searchUPRIGHT, searchDOWNLEFT, searchDOWNRIGHT} {
					if search(grid, i, j) {
						founds++
					}
				}
			}
		}
	}
	return founds
}

// /Functions for part 2
func searchX_MAS_TOPLEFT(grid [][]string, i, j int) bool {
	//Check if the part of the cross starting from the top left corner is a MAS or SAM

	verifIndex := i >= 1 && j >= 1 && i < len(grid[0])-1 && j < len(grid)-1
	if verifIndex {
		return (grid[i-1][j-1] == "S" && grid[i+1][j+1] == "M") || (grid[i-1][j-1] == "M" && grid[i+1][j+1] == "S")
	}
	return false
}

func searchX_MAS_DOWNLEFT(grid [][]string, i, j int) bool {
	//Check if the part of the cross starting from the down left corner is a MAS or SAM

	verifIndex := i >= 1 && j >= 1 && i < len(grid[0])-1 && j < len(grid)-1
	if verifIndex {
		return (grid[i+1][j-1] == "S" && grid[i-1][j+1] == "M") || (grid[i+1][j-1] == "M" && grid[i-1][j+1] == "S")
	}
	return false
}

func searchX_MAS(grid [][]string) int {
	founds := 0
	for i, line := range grid {
		for j, char := range line {
			if j == 0 || j == len(line)-1 || i == 0 || i == len(grid)-1 {
				//can't be a MAS or SAM at the edge of the grid
				continue
			}
			if char == "A" && searchX_MAS_TOPLEFT(grid, i, j) && searchX_MAS_DOWNLEFT(grid, i, j) {
				founds++
			}
		}
	}
	return founds
}

func main() {
	args := os.Args
	if args[1] == "1" {
		part1()
	}
	if args[1] == "2" {
		part2()
	}
}

func part1() {
	allLines := readInput()
	res := searchXMAS(allLines)
	fmt.Println("Result :", res)

}

func part2() {
	allLines := readInput()
	res := searchX_MAS(allLines)
	fmt.Println("Result :", res)
}
