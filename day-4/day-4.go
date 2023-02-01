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
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileScan := bufio.NewScanner(readFile)
	sum := 0
	ol := 0
	for fileScan.Scan() {
		item := fileScan.Text()
		split := strings.Split(item, ",")
		p1 := split[:1]
		p2 := split[1:]
		sp1 := splitter(p1)
		sp2 := splitter(p2)
		sum += fullOverlap(sp1, sp2)
		ol += partialOverlap(sp1, sp2)
	}
	fmt.Printf("Found %d full overlaps.\n", sum)
	fmt.Printf("Found %d partial overlaps.\n", ol)
}

// Splits string arrays based on "-"
func splitter(split []string) []int {
	result := []int{}
	for _, r := range split {
		for _, t := range strings.Split(r, "-") {
			val, _ := strconv.Atoi(t)
			result = append(result, val)
		}
	}
	return result
}

// Finds if it overlaps at all
func partialOverlap(a, b []int) int {
	if (a[0] >= b[0] && a[0] <= b[1]) || (b[0] >= a[0] && b[0] <= a[1]) {
		return 1
	}
	return 0
}

// Finds full overlaps
func fullOverlap(a, b []int) int {
	if (a[0] >= b[0] && a[1] <= b[1]) || (b[0] >= a[0] && b[1] <= a[1]) {
		return 1
	}
	return 0
}
