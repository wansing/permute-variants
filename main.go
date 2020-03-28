package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var parts [][]string // part -> variants

func main() {

	var partFilenames = os.Args[1:]

	// read input files

	for _, partFilename := range partFilenames {

		contents, err := ioutil.ReadFile(partFilename)
		if err != nil {
			panic(err)
		}

		var variants = strings.Split(strings.TrimSpace(strings.ReplaceAll(string(contents), "\r\n", "\n")), "\n")
		parts = append(parts, variants)
	}

	// 1. all permutations of parts: Heap's algorithm
	// 2. all variants of each part: recursive

	var partIndices = make([]int, len(parts)) // just [0, 1, 2, ...]
	for i := range parts {
		partIndices[i] = i
	}

	heap(partIndices, func(partOrder []int) {
		var variantState = make([]int, len(parts))
		variants(partOrder, &variantState, 0)
	})
}

// variants recursively determines all combinations of variants, given an order of parts
//
// numberOfFixedElements refers to variantState
func variants(partOrder []int, variantState *[]int, numberOfFixedElements int) {

	// break condition
	if numberOfFixedElements == len(parts) {
		for _, i := range partOrder {
			fmt.Print(parts[i][(*variantState)[i]])
		}
		fmt.Println()
		return
	}

	// recursion
	for i := range parts[numberOfFixedElements] {
		(*variantState)[numberOfFixedElements] = i
		variants(partOrder, variantState, numberOfFixedElements+1)
	}
}

// Heap's algorithm, non-recursive, from https://en.wikipedia.org/wiki/Heap%27s_algorithm
func heap(A []int, output func([]int)) {
	var n = len(A)
	var c = make([]int, n)
	output(A)
	var i = 0
	var temp int
	for i < n {
		if c[i] < i {
			if i % 2 == 0 {
				temp = A[0]
				A[0] = A[i]
				A[i] = temp
			} else {
				temp = A[c[i]]
				A[c[i]] = A[i]
				A[i] = temp
			}
			output(A)
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
}
