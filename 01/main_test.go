package main

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	inputs := []struct {
		name     string
		list     []int
		val      int
		expected []int
	}{
		{
			name:     "insert at the end of list",
			list:     []int{1, 2, 3},
			val:      4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "insert into middle of list",
			list:     []int{1, 3},
			val:      2,
			expected: []int{1, 2, 3},
		},
		{
			name:     "insert into empty list",
			list:     []int{},
			val:      1,
			expected: []int{1},
		},
	}

	for _, test := range inputs {
		t.Run(test.name, func(t *testing.T) {
			actual := insert(test.val, test.list)
			if len(actual) != len(test.expected) {
				fmt.Printf("\texpected: %v\n\tactual: %v\n", test.expected, actual)
				t.FailNow()
			}
			for i, val := range actual {
				if val != test.expected[i] {
					fmt.Printf("%d != %d\n", val, test.expected[i])
					fmt.Printf("\texpected: %v\n\tactual: %v\n", test.expected, actual)
					t.Fail()

				}
			}
		})
	}
}

func TestCalcTotalDistance(t *testing.T) {
	inputs := []struct {
		name        string
		inputStruct input
		expected    int
	}{
		{
			name: "example",
			inputStruct: input{
				leftList:  []int{1, 2, 3, 3, 3, 4},
				rightList: []int{3, 3, 3, 4, 5, 9},
			},
			expected: 11,
		},
	}

	for _, test := range inputs {
		t.Run(test.name, func(t *testing.T) {
			actual := test.inputStruct.calcTotalDistance()
			if actual != test.expected {
				fmt.Printf("%d != %d\n", actual, test.expected)
				t.Fail()
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	input := []struct {
		name          string
		target        int
		inputStruct   input
		expectedLeft  int
		expectedRight int
	}{
		{
			name:   "single index and dne",
			target: 2,
			inputStruct: input{
				leftList:  []int{1, 2, 3, 3, 3, 4},
				rightList: []int{3, 3, 3, 4, 5, 9},
			},
			expectedLeft:  1,
			expectedRight: -1,
		},
		{
			name:   "exists in both",
			target: 4,
			inputStruct: input{
				leftList:  []int{1, 2, 3, 3, 3, 4},
				rightList: []int{3, 3, 3, 4, 5, 9},
			},
			expectedLeft:  5,
			expectedRight: 3,
		},
		{
			name:   "multiple exists",
			target: 3,
			inputStruct: input{
				leftList:  []int{1, 2, 3, 3, 3, 4},
				rightList: []int{3, 3, 3, 4, 5, 9},
			},
			expectedLeft:  3,
			expectedRight: 1,
		},
	}
	for _, test := range input {
		t.Run(test.name, func(t *testing.T) {
			actualLeftIndex := binarySearch(test.target, test.inputStruct.leftList)
			actualRightIndex := binarySearch(test.target, test.inputStruct.rightList)
			if actualLeftIndex != test.expectedLeft {
				fmt.Printf("expectedLeft: %d | actualLeft: %d\n", test.expectedLeft, actualLeftIndex)
				t.Fail()
			}
			if actualRightIndex != test.expectedRight {
				fmt.Printf("expectedRight: %d | actualRight: %d\n", test.expectedRight, actualRightIndex)
				t.Fail()
			}
		})
	}
}

func TestSimilarityScore(t *testing.T) {
	inputs := []struct {
		name        string
		inputStruct input
		expected    int
	}{
		{
			name: "example",
			inputStruct: input{
				leftList:  []int{1, 2, 3, 3, 3, 4},
				rightList: []int{3, 3, 3, 4, 5, 9},
			},
			expected: 31,
		},
	}
	for _, test := range inputs {
		t.Run(test.name, func(t *testing.T) {
			actual := test.inputStruct.findSimilarityScore()
			if actual != test.expected {
				fmt.Printf("expected: %d | actual: %d\n", test.expected, actual)
				t.Fail()
			}
		})
	}
}
