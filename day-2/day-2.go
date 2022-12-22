package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	// score := 0
	fix := 0

	// for fileScanner.Scan() {
	// 	match := fileScanner.Text()
	// 	score += matchScore(match)
	// }

	// fmt.Println(score)

	for fileScanner.Scan() {
		match := fileScanner.Text()
		fix += fixedScore(match)
	}

	fmt.Println(fix)
}

func fixedScore(match string) int {
	x := 0
	switch match[len(match)-1:] {
	case "X":
		switch match[:1] {
		case "A":
			x = 3
		case "B":
			x = 1
		case "C":
			x = 2
		default:
			x = 0
		}
	case "Y":
		switch match[:1] {
		case "A":
			x = 4
		case "B":
			x = 5
		case "C":
			x = 6
		default:
			x = 0
		}
	case "Z":
		switch match[:1] {
		case "A":
			x = 8
		case "B":
			x = 9
		case "C":
			x = 7
		default:
			x = 0
		}
	}
	return x
}

func matchScore(match string) int {
	x := 0
	switch match {
	case "A X":
		x = 4
	case "A Y":
		x = 8
	case "A Z":
		x = 3
	case "B X":
		x = 1
	case "B Y":
		x = 5
	case "B Z":
		x = 9
	case "C X":
		x = 7
	case "C Y":
		x = 2
	case "C Z":
		x = 6
	default:
		x = 0
	}
	return x
}
