package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	UP        = 0
	RIGHT     = 1
	DOWN      = 2
	LEFT      = 3
	DIRECTION = 0
	XGUARD    = -1
	YGUARD    = -1
)

//Functions for both

func readInput() [][]int {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	reader := bufio.NewReader(f)
	numberLine := 0
	mapInput := [][]int{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break //will break at the end of the file
		}
		mapInput = append(mapInput, []int{})
		for y, char := range line {
			switch string(char) {
			case ".":
				mapInput[numberLine] = append(mapInput[numberLine], 0)
			case "#":
				mapInput[numberLine] = append(mapInput[numberLine], 1)
			case "^":
				mapInput[numberLine] = append(mapInput[numberLine], 2)
				XGUARD = numberLine
				YGUARD = y
			}
		}
		numberLine++
	}
	defer f.Close()
	return mapInput
}

func turn() {
	DIRECTION = (DIRECTION + 1) % 4
}

func goForwardWhatToAdd() (int, int) {
	switch DIRECTION {
	case UP:
		return -1, 0
	case RIGHT:
		return 0, 1
	case DOWN:
		return 1, 0
	case LEFT:
		return 0, -1
	}
	return 0, 0
}

func isInBorder(allMap [][]int, newX int, newY int) bool {
	return newX < len(allMap) && newX >= 0 && newY < len(allMap[0]) && newY >= 0
}

// /Functions for part 1

func goForwardOne(allMap [][]int, newX int, newY int) {
	allMap[newX][newY] = 2
	XGUARD = newX
	YGUARD = newY
}

func goForward(allMap [][]int) {
	for {
		toAddx, toAddy := goForwardWhatToAdd()
		if isInBorder(allMap, XGUARD+toAddx, YGUARD+toAddy) {
			newX := XGUARD + toAddx
			newY := YGUARD + toAddy
			if allMap[newX][newY] == 1 {
				turn()
				continue
			}
			goForwardOne(allMap, newX, newY)
		} else {
			break
		}
	}
}

func countPath(allMap [][]int) int {
	count := 0
	for _, line := range allMap {
		for _, column := range line {
			if column == 2 {
				count++
			}
		}
	}
	return count
}

// /Functions for part 2

//MAIN

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
	allMap := readInput()
	goForward(allMap)
	res := countPath(allMap)
	fmt.Println("Result :", res)

}

func part2() {

	fmt.Println("Result :")
}
