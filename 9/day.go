package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
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

func readInput() []int {
	//Let's define the map as a 3d matrix, [x][y][0] is the map as the first part ask it, [x][y][1:] are for the direction took on those paths
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	reader := bufio.NewReader(f)
	disk := []int{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break //will break at the end of the file
		}
		for _, char := range line {
			number, _ := strconv.Atoi(string(char))
			disk = append(disk, number)
		}
	}
	defer f.Close()
	return disk
}

func nextIDToMove(ids []int, previous int) int {
	for i := previous - 1; i > 0; i-- {
		if ids[i] != -1 {
			return i
		}
	}
	return -1
}

func checkSum(disk []int) int {
	sum := 0
	for i, id := range disk {
		if id == -1 {
			continue //Continue for the second part, break for the first
		}
		sum += i * id
	}
	return sum
}

//FUNCTION PART 1

func reOrderId(disk []int) {
	nextID := nextIDToMove(disk, len(disk))
	nextSpace := nextFreeSpace(disk, 0)
	for {
		if nextSpace > nextID {
			break
		}
		moveID(disk, nextID, nextSpace)
		nextSpace = nextFreeSpace(disk, nextSpace)
		nextID = nextIDToMove(disk, nextID)
	}
}
func createID(disk []int) []int {
	//ID = -1 <=> char = "."
	ids := []int{}
	for index, number := range disk {
		if index%2 == 1 {
			//Free space
			for i := 0; i < number; i++ {
				ids = append(ids, -1)
			}
		} else {
			//ID
			for i := 0; i < number; i++ {
				ids = append(ids, index/2)
			}
		}
	}
	return ids
}

func moveID(ids []int, index1 int, index2 int) {
	ids[index1], ids[index2] = ids[index2], ids[index1]
}

func nextFreeSpace(ids []int, previous int) int {
	for i := previous; i < len(ids); i++ {
		if ids[i] == -1 {
			return i
		}
	}
	return -1
}

//FUNCTION PART 2

func searchSpaceLength(disk []int, length int) int {
	if length == -1 {
		return -1
	}
	for index := range disk[:len(disk)-length+1] {
		found := false
		for _, id := range disk[index : index+length] {
			if id != -1 {
				found = false
				break
			}
			found = true
		}
		if found {
			return index
		}
	}
	return -1
}

func searchIDs(disk []int, previous int) (int, int) {
	//Return length and last index of the same ids
	nextID := nextIDToMove(disk, previous)
	if nextID == -1 {
		return -1, -1
	}
	length := 0
	for {
		if disk[nextID-length] != disk[nextID] {
			break
		}
		length++
		if nextID-length < 0 {
			return -1, -1
		}
	}
	return length, nextID
}

func reOrderIdBlock(disk []int) {
	length, nextID := searchIDs(disk, len(disk))
	nextSpace := searchSpaceLength(disk, length)
	for {
		if nextSpace > nextID {
			continue
		}
		if nextID == -1 {
			break
		}
		for i := 0; i < length; i++ {
			moveID(disk, nextID-i, nextSpace+i)
		}
		for {
			length, nextID = searchIDs(disk, nextID-length+1)
			nextSpace = searchSpaceLength(disk, length)
			if nextSpace != -1 && nextSpace < nextID {
				break
			}
			if nextID == -1 {
				break
			}
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
	disk := readInput()
	diskID := createID(disk)
	reOrderId(diskID)
	res := checkSum(diskID)
	fmt.Println("Result :", res)

}

func part2() {
	t1 := time.Now()
	disk := readInput()
	diskID := createID(disk)
	reOrderIdBlock(diskID)
	res := checkSum(diskID)
	t2 := time.Now()
	fmt.Println("Result :", res, "time :", t2.Sub(t1).Seconds())
}
