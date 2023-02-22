package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Crate = string
type Stack = []Crate
type Stacks = []Stack

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileScan := bufio.NewScanner(readFile)

	// Create an empty array for raw stack lines
	// in the for loop append each raw stack line to the array
	// if we hit an empty line, then we break
	var rawLines []string
	for fileScan.Scan() {
		request := fileScan.Text()
		if request == "" {
			break
		}
		rawLines = append(rawLines, request)
	}
	stacks := makeStacks(rawLines[:len(rawLines)-1])
	fmt.Printf("Input %s\n", stacks)

	// pass the array (except the last item) to makeStacks func
	for fileScan.Scan() {
		request := fileScan.Text()
		moveStuff2(request, stacks)
	}

	fmt.Printf("Final:  %s\n", stacks)
	for _, stack := range stacks {
		fmt.Printf("%s", stack[len(stack)-1])
	}
}

// function moves single crate at a time, inverts the order
func moveStuff(moveLine string, stacks Stacks) {
	split := strings.Split(moveLine, " ")
	howMany, _ := strconv.Atoi(split[1])
	from, _ := strconv.Atoi(split[3])
	to, _ := strconv.Atoi(split[5])
	from = from - 1
	to = to - 1
	for i := howMany; i > 0; i-- {
		fromLength := len(stacks[from]) - 1
		move := stacks[from][fromLength]
		stacks[to] = append(stacks[to], move)
		stacks[from] = stacks[from][:fromLength]
	}
}

// function moves multiple crates at once, maintains the order
func moveStuff2(moveLine string, stacks Stacks) {
	split := strings.Split(moveLine, " ")
	howMany, _ := strconv.Atoi(split[1])
	from, _ := strconv.Atoi(split[3])
	to, _ := strconv.Atoi(split[5])
	from = from - 1
	to = to - 1
	fromLength := len(stacks[from]) - howMany
	move := stacks[from][fromLength:]
	stacks[to] = append(stacks[to], move...)
	stacks[from] = stacks[from][:fromLength]
}

// Create a stacks array that will be populated with crates
// use make to create an array of certain length
func makeStacks(lines []string) Stacks {
	stackSize := (len(lines[0]) + 1) / 4
	stacks := make(Stacks, stackSize)
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		for j := 1; j < len(line); j += 4 {
			crate := line[j : j+1]
			if crate != " " {
				stacks[j/4] = append(stacks[j/4], crate)
			}
		}
	}
	return stacks
}
