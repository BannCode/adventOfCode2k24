package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	allEquation := [][]int{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break //will break at the end of the file
		}
		allEquation = append(allEquation, []int{})
		equation := strings.SplitN(line, ":", 2)
		total, _ := strconv.Atoi(equation[0])
		allEquation[numberLine] = append(allEquation[numberLine], total)

		//The format is <total: 1stNbr 2ndNbr ...> so by splitting by spaces the first is empty
		//And the last character is an \n so we take it off
		for _, number := range strings.Split(equation[1][1:len(equation[1])-1], " ") {
			testValue, _ := strconv.Atoi(number)
			allEquation[numberLine] = append(allEquation[numberLine], testValue)
		}
		numberLine++
	}
	defer f.Close()
	return allEquation
}

// /Functions for part 1

func operand(value1 int, value2 int, operand int) int {
	switch operand {
	case ADD:
		return value1 + value2
	case MULT:
		return value1 * value2
	case PIPE:
		res, _ := strconv.Atoi(strconv.Itoa(value1) + strconv.Itoa(value2))
		return res
	}
	return 0
}

func createOperation(equation []int) []int {
	//The first slot is for the total, the rest are for the operands
	//there is len(equation)-1 number to add/multiply, so there is len(equation)-2 operands
	res := make([]int, len(equation)-1)
	for index, _ := range res {
		res[index] = ADD
	}
	return res
}

func updateTotalOperation(equation []int, operations []int) int {
	total := equation[1]
	for i, op := range operations[1:] {
		//even when starting operations at 1, i starts at 0 so we search for equation[i+2]
		total = operand(total, equation[i+2], op)
		operations[0] = total
		if total > equation[0] {
			return TOOHIGH
		}
	}
	if total == equation[0] {
		return EQUAL
	}
	return TOOLOW
}

func modifOperation(operations []int, index int, nbOp int) {
	operations[index] = (operations[index] + 1) % nbOp
}

func searchCombination(equation []int, operations []int, index int, nbOp int) int {
	if index == len(operations) {
		return updateTotalOperation(equation, operations)
	}
	if updateTotalOperation(equation, operations) != EQUAL {
		//choose the best option
		if searchCombination(equation, operations, index+1, nbOp) == EQUAL {
			return EQUAL
		}
		if nbOp == 3 {
			modifOperation(operations, index, nbOp)
			if searchCombination(equation, operations, index+1, nbOp) == EQUAL {
				return EQUAL
			}
		}
		modifOperation(operations, index, nbOp)
		return searchCombination(equation, operations, index+1, nbOp)
	}
	return EQUAL
}

func findCombination(equation []int, nbOp int) []int {
	operations := createOperation(equation)
	updateTotalOperation(equation, operations)

	if searchCombination(equation, operations, 1, nbOp) == EQUAL {
		return operations
	}
	// else {
	// 	modifOperation(operations, 1)
	// 	searchCombination(equation, operations, 1)
	// }
	return make([]int, 0)
}

func findAllCombination(allEquation [][]int, nbOp int) [][]int {
	allOperation := [][]int{}
	for _, equation := range allEquation {
		operation := findCombination(equation, nbOp)
		allOperation = append(allOperation, operation)
	}
	return allOperation
}

func countCombination(allOperation [][]int) int {
	res := 0
	for _, operation := range allOperation {
		if len(operation) > 1 {
			res += operation[0]
		}
	}
	return res
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
	allEquation := readInput()
	allOperation := findAllCombination(allEquation, 2)
	res := countCombination(allOperation)
	fmt.Println("Result :", res)

}

func part2() {
	t1 := time.Now()
	allEquation := readInput()
	allOperation := findAllCombination(allEquation, 3)
	res := countCombination(allOperation)
	t2 := time.Now()
	fmt.Println("Result :", res, "time :", t2.Sub(t1).Seconds())
}
