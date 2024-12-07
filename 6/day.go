package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var (
	UP        = 3
	RIGHT     = 4
	DOWN      = 5
	LEFT      = 6
	DIRECTION = 3
	XGUARD    = -1
	YGUARD    = -1
	OBSTACLE  = 1
	EMPTY     = 0
	VISITED   = 2
)

//Functions for both

func contains(slice []int, value int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func readInput() [][][]int {
	//Let's define the map as a 3d matrix, [x][y][0] is the map as the first part ask it, [x][y][1:] are for the direction took on those paths
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	reader := bufio.NewReader(f)
	numberLine := 0
	mapInput := [][][]int{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break //will break at the end of the file
		}
		mapInput = append(mapInput, [][]int{})
		for y, char := range line {
			switch string(char) {
			case ".":
				mapInput[numberLine] = append(mapInput[numberLine], []int{})
				mapInput[numberLine][y] = append(mapInput[numberLine][y], EMPTY)
			case "#":
				mapInput[numberLine] = append(mapInput[numberLine], []int{})
				mapInput[numberLine][y] = append(mapInput[numberLine][y], OBSTACLE)
			case "^":
				mapInput[numberLine] = append(mapInput[numberLine], []int{})
				mapInput[numberLine][y] = append(mapInput[numberLine][y], VISITED)
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
	DIRECTION = (DIRECTION+1-3)%4 + 3
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

func isInBorder(allMap [][][]int, newX int, newY int) bool {
	return newX < len(allMap) && newX >= 0 && newY < len(allMap[0]) && newY >= 0
}

// /Functions for part 1
func isInLoop(allMap [][][]int) bool {
	return contains(allMap[XGUARD][YGUARD][:len(allMap[XGUARD][YGUARD])-1], DIRECTION) //more than once the direction
}
func goForwardOne(allMap [][][]int, newX int, newY int) {
	allMap[newX][newY][0] = VISITED
	XGUARD = newX
	YGUARD = newY
	allMap[newX][newY] = append(allMap[newX][newY], DIRECTION) //So we remember the direction on this position
}

func goForward(allMap [][][]int) bool {
	//Return true if in a loop, false otherwise
	for {
		if isInLoop(allMap) {
			return true
		}
		toAddx, toAddy := goForwardWhatToAdd()
		if isInBorder(allMap, XGUARD+toAddx, YGUARD+toAddy) {
			newX := XGUARD + toAddx
			newY := YGUARD + toAddy
			if allMap[newX][newY][0] == OBSTACLE {
				turn()
				continue
			}
			goForwardOne(allMap, newX, newY)
		} else {
			return false
		}
	}
}

func countPath(allMap [][][]int) int {
	count := 0
	for _, line := range allMap {
		for _, column := range line {
			if column[0] == VISITED {
				count++
			}
		}
	}
	return count
}

// /Functions for part 2

func resetPos(initX int, initY int, initDirection int) {
	XGUARD = initX
	YGUARD = initY
	DIRECTION = initDirection
}

func copyMap(allMap [][][]int) [][][]int {
	copyAllMap := make([][][]int, len(allMap))
	for i := range allMap {
		copyAllMap[i] = make([][]int, len(allMap[i]))
		for j := range allMap[i] {
			copyAllMap[i][j] = append([]int{}, allMap[i][j]...) // Copie indépendante
		}
	}
	return copyAllMap
}

func tryAllObstacle(allMap [][][]int) [][]int {
	initX := XGUARD
	initY := YGUARD
	initDirection := DIRECTION
	allCoords := [][]int{}
	for x, line := range allMap {
		for y, column := range line {
			if column[0] != OBSTACLE {
				//Not already an obstacle

				copyAllMap := copyMap(allMap)

				//Becomes an obstacle
				copyAllMap[x][y][0] = OBSTACLE
				if goForward(copyAllMap) {
					allCoords = append(allCoords, []int{x, y})
				}
				resetPos(initX, initY, initDirection)
			}
		}
	}
	return allCoords
}

func countObstacles(allCoords [][]int) int {
	return len(allCoords)
}

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
	t1 := time.Now()
	allMap := readInput()
	allCoords := tryAllObstacle(allMap)
	res := countObstacles(allCoords)
	t2 := time.Now()
	fmt.Println("Result :", res, "time :", t2.Sub(t1).Seconds())
}
