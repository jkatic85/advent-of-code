package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	sum1 := 0
	sum2 := 0

	for fileScan.Scan() {
		item := fileScan.Text()
		p1 := item[:len(item)/2]
		p2 := item[len(item)/2:]
		fmt.Println(p1 + " " + p2)
		rep := 0
		rep1 := 0
		for _, r1 := range p1 {
			fmt.Printf("Checking: %c\n", r1)
			// strings.Contains
			isTrue := strings.Contains(p2, string(r1))
			if isTrue && rep1 == 0 {
				fmt.Printf("strings.Contains found a duplicate: %c! \n", r1)
				sum += priority(r1)
				rep1++
			}
			for _, r2 := range p2 {
				if r1 == r2 && rep == 0 {
					fmt.Printf("Second for loop found a duplicate: %c! \n", r1)
					sum1 += priority(r2)
					sum2 += priorityChar(r2)
					rep++
					break
				}
			}
		}
	}
	fmt.Println(sum, sum1, sum2)
}

//  Array approach:
// 		Range over a string adding a index +1 value
const str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func priority(r rune) int {
	return strings.Index(str, string(r)) + 1
}

// Numeric approach:
//		Get utf-8 char code for each letter and convert it to a number between 1 - 52
func priorityChar(r rune) int {
	if r < rune(97) {
		return int((r - 38))
	} else {
		return int((r - 96))
	}
}

// to-do create func that finds duplicate item in rucksack
// it returns the value of the item
// it does not continue iterating after finding the item
