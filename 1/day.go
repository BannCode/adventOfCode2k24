package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func addSortedInt(slice []int, value int) []int {
	index := sort.SearchInts(slice, value) //binary search

	slice = append(slice[:index], append([]int{value}, slice[index:]...)...) // the ... for adding each element of the slice 1 by 1

	return slice
}

func absForInts(value int) int {
	if value < 0 {
		return -value
	}
	return value
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
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	sliceLeft := []int{}
	sliceRight := []int{}
	reader := bufio.NewReader(f)
	i := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break //will break at the end of the file
		}
		res := strings.Fields(line) //Delete the spaces and split the remaining parts
		if len(res) == 2 {
			leftNumber, _ := strconv.Atoi(res[0]) // Convert in Integer
			rightNumber, _ := strconv.Atoi(res[1])
			sliceLeft = addSortedInt(sliceLeft, leftNumber)
			sliceRight = addSortedInt(sliceRight, rightNumber)
		}
		i++
	}
	defer f.Close()

	res := 0
	for i := 0; i < len(sliceLeft); i++ {
		res += absForInts(sliceLeft[i] - sliceRight[i])
	}

	fmt.Printf("Result : %d", res)

}

func part2() {

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	sliceLeft := []int{}
	rightList := make(map[int]int)
	reader := bufio.NewReader(f)
	i := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break //will break at the end of the file
		}
		res := strings.Fields(line) //Delete the spaces and split the remaining parts
		if len(res) == 2 {
			leftNumber, _ := strconv.Atoi(res[0]) // Convert in Integer
			rightNumber, _ := strconv.Atoi(res[1])
			sliceLeft = append(sliceLeft, leftNumber)
			rightList[rightNumber]++
		}
		i++
	}
	defer f.Close()

	res := 0
	for i := 0; i < len(sliceLeft); i++ {
		res += sliceLeft[i] * rightList[sliceLeft[i]]
	}

	fmt.Printf("Result : %d", res)

}
