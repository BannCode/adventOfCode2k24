package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func absForInts(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func addFromMuls(line string, re *regexp.Regexp) int {
	matches := re.FindAllStringSubmatch(line, -1)
	res := 0
	for _, match := range matches {
		leftNum, _ := strconv.Atoi(match[1])
		rightNum, _ := strconv.Atoi(match[2])
		res += leftNum * rightNum
	}
	return res
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
	res := 0
	pattern := `mul\(([0-9]{1,3}),([0-9]{1,3})\)` // Format : mul(%d,%d)
	re := regexp.MustCompile(pattern)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break //will break at the end of the file
		}
		res += addFromMuls(line, re)
	}
	defer f.Close()

	fmt.Println("Result :", res)

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
	pattern := `mul\(([0-9]{1,3}),([0-9]{1,3})\)` // Format : mul(%d,%d)
	re := regexp.MustCompile(pattern)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break //will break at the end of the file
		}
		for indexDont, betweenDont := range strings.Split(line, "don't()") {
			splitAtFirstDo := strings.SplitN(betweenDont, "do()", 2)
			if indexDont == 0 {
				res += addFromMuls(splitAtFirstDo[0], re)
			}
			if len(splitAtFirstDo) > 1 {
				res += addFromMuls(splitAtFirstDo[1], re)
			}
		}
		i++
	}
	defer f.Close()

	fmt.Println("Result :", res)
}
