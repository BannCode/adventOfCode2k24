package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func absForInts(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func isSafe(layer []string) bool {
	toAdd := true
	isIncreas := 0
	for k := 1; k < len(layer); k++ {
		previous, _ := strconv.Atoi(layer[k-1])
		next, _ := strconv.Atoi(layer[k])
		diff := (next - previous)
		absDiff := absForInts(diff)
		toAdd = absDiff > 0 && absDiff <= 3 && (diff*isIncreas > 0 || isIncreas == 0)
		if !toAdd {
			break
		}
		if isIncreas == 0 {
			isIncreas = diff / absDiff // isIncreas = 1 if increasing, -1 otherwise
		}
	}
	return toAdd
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
	reader := bufio.NewReader(f)
	i := 0
	res := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break //will break at the end of the file
		}
		layer := strings.Fields(line) //Delete the spaces and split the remaining parts
		if isSafe(layer) {
			res++
			// fmt.Println(layer)
		}
		i++
	}
	defer f.Close()

	fmt.Println("Result :", res)

}

func popForSlice(layer []string, index int) []string {
	if index >= len(layer) {
		fmt.Println("Problem here : ", layer)
		return layer
	}
	newLayer := make([]string, len(layer))
	copy(newLayer, layer)
	return append(newLayer[:index], newLayer[index+1:]...)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(f)
	i := 0
	res := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break //will break at the end of the file
		}
		layer := strings.Fields(line) //Delete the spaces and split the remaining parts
		if isSafe(layer) {
			res++
			// fmt.Println(layer)
		} else {
			for k := 0; k < len(layer); k++ {
				withoutK := popForSlice(layer, k)
				if isSafe(withoutK) {
					res++
					fmt.Println("Example of new success : ", layer)
					break
				}
			}
		}
		i++
	}
	defer f.Close()

	fmt.Println("Result :", res)

}
