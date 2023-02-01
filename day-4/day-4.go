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
	split := []string{}
	sum := 0
	ol := 0
	for fileScan.Scan() {
		item := fileScan.Text()
		// Splits strings based on ","
		split = strings.Split(item, ",")
		// fmt.Println(split)
		p1 := split[:1]
		p2 := split[1:]
		sp1 := splitter(p1)
		sp2 := splitter(p2)
		fmt.Println(sp1, sp2)
		fmt.Println("------------")
		fmt.Println(checker(sp1, sp2))
		fmt.Println("----")
		fmt.Println(overlap(sp1, sp2))
		fmt.Println("============")
		sum += checker(sp1, sp2)
		ol += overlap(sp1, sp2)
	}
	fmt.Printf("Found %d duplicates.\n", sum)
	fmt.Printf("Found %d overlaps.\n", ol)

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
func overlap(a, b []int) int {
	firstA := a[0]
	lastA := a[1]
	firstB := b[0]
	lastB := b[1]
	c := 0
	if firstA >= firstB && firstA <= lastB {
		c++
	} else if firstB >= firstA && firstB <= lastA {
		c++
	} else {
		c += 0
	}
	return c
}

// Finds full overlaps
func checker(a, b []int) int {
	firstA := a[0]
	lastA := a[1]
	firstB := b[0]
	lastB := b[1]
	c := 0
	if firstA >= firstB && lastA <= lastB {
		c++
	} else if firstB >= firstA && lastB <= lastA {
		c++
	} else {
		c += 0
	}
	return c
}
