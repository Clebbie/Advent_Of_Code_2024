package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// construct input and sorted lists
	lists, err := readInput("./input/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	// print answer
	fmt.Printf("totoalDistance: %d\n", lists.calcTotalDistance())
	fmt.Printf("similarityScore: %d\n", lists.findSimilarityScore())
}

// insert takes a values and inserts insert into the array returning a new
// array after completion. Note that this array is unstably sorted!
func insert(val int, list []int) []int {
	out := []int{}
	inserted := false
	for _, v := range list {
		if !inserted && v >= val {
			inserted = true
			out = append(out, val)
		}
		out = append(out, v)
	}
	if !inserted {
		out = append(out, val)
	}
	return out
}

// input is a struct representing the input to the puzzle.
// in this case, it wraps two slices and has the method to
// calcTotalDistance
type input struct {
	leftList  []int
	rightList []int
}

// calcTotalDistance is what determines the answer to the puzzle.
// NOTICE: this adds together the absolute value of all the distances
func (i *input) calcTotalDistance() int {
	total := 0

	for j, val := range i.leftList {
		localSum := val - (i.rightList)[j]
		if localSum < 0 {

			fmt.Printf("%5d = %5d - %5d\n", (localSum * -1), val, i.rightList[j])
			total = total + (localSum * -1)
			continue
		}
		fmt.Printf("%5d = %5d - %5d\n", localSum, val, i.rightList[j])
		total = total + localSum
	}

	return total
}

func (in *input) findSimilarityScore() int {
	seenLocations := map[int]frequencyReport{}
	for i := 0; i < len(in.leftList); i++ {
		report := frequencyReport{}
		fmt.Printf("searching cache for %d\n", in.leftList[i])
		_, ok := seenLocations[in.leftList[i]]
		if !ok {
			fmt.Printf("not found in cache... scoring....\n")
			report.left = greedyGrab(in.leftList[i], in.leftList)
			report.right = greedyGrab(in.leftList[i], in.rightList)
			seenLocations[in.leftList[i]] = report
		}
	}

	similarityScore := 0
	for _, val := range in.leftList {
		report := seenLocations[val]
		fmt.Println(val * report.right)
		similarityScore = similarityScore + (val * report.right)
	}
	return similarityScore
}

// frequencyReport wraps the count of some number in the left list and right list
type frequencyReport struct {
	left  int
	right int
}

// binarySearch uses a binary search algo to return the first index found of the target
func binarySearch(target int, list []int) int {
	mid := len(list) / 2
	if target == list[mid] {
		return mid
	}
	if mid == 0 {
		return -1
	}
	if target < list[mid] {
		result := binarySearch(target, list[:mid])
		if result == -1 {
			return result
		}
		return result

	}
	result := binarySearch(target, list[mid:])
	if result == -1 {
		return result
	}
	return mid + result
}

// greedyGrab finds where to start grabbing the target number from the list
// then counts any occurences below or above the the starting index
func greedyGrab(target int, list []int) int {
	fmt.Printf("searching for %5d in %v\n", target, list)
	count := 0
	startIndex := binarySearch(target, list)
	if startIndex == -1 {
		fmt.Printf("not found!\n")
		return 0
	}
	// count the return index and search down
	for i := startIndex; i >= 0; i-- {
		if i < 0 {
			break
		}
		if list[i] == target {
			fmt.Print("found in left\n")
			count++
		}
	}
	// skip the returned index and search up
	for i := startIndex + 1; i < len(list); i++ {
		if target == list[i] {
			fmt.Println("found in right")
			count++
		}
	}
	return count
}

// readInput simply reads the input from a text file and transforms it into
// the input struct
// NOTICE: while creating the struct, it is sorting each list.
func readInput(pathToFile string) (*input, error) {
	f, err := os.Open(pathToFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	left := []int{}
	right := []int{}
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), "   ")
		leftVal, err := strconv.Atoi(vals[0])
		if err != nil {
			return nil, err
		}
		rightVal, err := strconv.Atoi(vals[1])
		if err != nil {
			return nil, err
		}
		left = insert(leftVal, left)
		right = insert(rightVal, right)
		continue
	}

	return &input{leftList: left, rightList: right}, nil
}
