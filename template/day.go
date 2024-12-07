package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var (
	PIPE    = 2
	MULT    = 1
	ADD     = 0
	TOOHIGH = 2
	EQUAL   = 1
	TOOLOW  = 0
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

func readInput() [][]int {
	//Let's define the map as a 3d matrix, [x][y][0] is the map as the first part ask it, [x][y][1:] are for the direction took on those paths
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	reader := bufio.NewReader(f)
	numberLine := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break //will break at the end of the file
		}
		fmt.Println(line)
		numberLine++
	}
	defer f.Close()
	return make([][]int, 0)
}

//FUNCTION PART 1

//FUNCTION PART 2

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
	fmt.Println("Result :")

}

func part2() {
	t1 := time.Now()
	t2 := time.Now()
	fmt.Println("Result :", "time :", t2.Sub(t1).Seconds())
}
