package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Functions for both

func readInput() (map[int][]int, [][]int) {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	reader := bufio.NewReader(f)
	i := 0

	allRules := make(map[int][]int)
	allUpdates := [][]int{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break //will break at the end of the file
		}
		rule := strings.SplitN(line, "|", 2)
		if len(rule) == 2 {
			//Real rule
			before, _ := strconv.Atoi(rule[0])
			after, _ := strconv.Atoi(rule[1][:len(rule[1])-1])
			allRules[before] = append(allRules[before], after)
		} else {
			update := strings.Split(strings.ReplaceAll(line, "\n", ""), ",")
			if len(update) > 1 {
				//Real update
				allUpdates = append(allUpdates, []int{})
				for _, char := range update {
					intChar, _ := strconv.Atoi(char)
					allUpdates[i] = append(allUpdates[i], intChar)
				}
				i++ //used only for the updates
			}
		}
	}
	defer f.Close()
	return allRules, allUpdates
}

func contains(slice []int, value int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
func checkUpdate(rules map[int][]int, update []int) bool {
	for index, numIndex := range update {
		// if index == 0 { //useless because update[:index] is empty for index = 0
		// 	continue
		// }
		for _, numToCheck := range update[:index] {
			if contains(rules[numIndex], numToCheck) {
				return false
			}
		}
	}
	return true
}

func addMiddles(updates [][]int, indexes []int) int {
	res := 0
	for _, index := range indexes {
		res += updates[index][len(updates[index])/2]
	}
	return res
}

// /Functions for part 1

func checkAllUpdates(rules map[int][]int, updates [][]int) []int {

	res := []int{}
	for index, update := range updates {
		if checkUpdate(rules, update) {
			res = append(res, index)
		}
	}
	return res
}

// /Functions for part 2

func checkAllNonUpdates(rules map[int][]int, updates [][]int) []int {

	res := []int{}
	for index, update := range updates {
		if !checkUpdate(rules, update) {
			res = append(res, index)
		}
	}
	return res
}

func updateOne(rules map[int][]int, toUpdate []int) {
	for cursor, numIndex := range toUpdate {
		for index, numToCheck := range toUpdate[:cursor] {
			if contains(rules[numIndex], numToCheck) {
				waitingValue := toUpdate[cursor]
				toUpdate[cursor] = toUpdate[index]
				toUpdate[index] = waitingValue
				return
			}
		}
	}
}

func updateAll(rules map[int][]int, updates [][]int, indexes []int) {
	for _, index := range indexes {
		count := 0
		for count <= len(updates[index])^2 && !checkUpdate(rules, updates[index]) {
			updateOne(rules, updates[index])
		}
	}
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
	allRules, allUpdates := readInput()
	fmt.Println("Number of rules :", len(allRules))
	middlesIndexes := checkAllUpdates(allRules, allUpdates)
	fmt.Println("Number of correct updates :", len(middlesIndexes), "number of updates : ", len(allUpdates))
	res := addMiddles(allUpdates, middlesIndexes)
	fmt.Println("Result :", res)

}

func part2() {
	allRules, allUpdates := readInput()
	middlesIndexes := checkAllNonUpdates(allRules, allUpdates)
	updateAll(allRules, allUpdates, middlesIndexes)
	res := addMiddles(allUpdates, middlesIndexes)
	fmt.Println("Number of corrected updates :", len(middlesIndexes), "number of updates : ", len(allUpdates))
	fmt.Println("Result :", res)
}
