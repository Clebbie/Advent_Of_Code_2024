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
	fmt.Println(lists.calcTotalDistance())
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

	return &input{left, right}, nil
}
