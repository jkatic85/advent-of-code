package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	readFile, err := os.Open("sample-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileScan := bufio.NewScanner(readFile)
	// sum := 0

	// for fileScan.Scan() {
	// 	item := fileScan.Text()
	// 	p1 := item[:len(item)/2]
	// 	p2 := item[len(item)/2:]
	// 	fmt.Println(p1 + " " + p2)
	// 	rep := 0
	// 	for _, r := range p1 {
	// 		c1 := string(r)
	// 		fmt.Printf("Checking: %s\n", c1)
	// 		for _, s := range p2 {
	// 			c2 := string(s)
	// 			if c1 == c2 && rep == 0 {
	// 				fmt.Printf("Found a duplicate: %s! \n", c1)
	// 				sum += prioritySum(c2)
	// 				rep++
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Println(sum)
	group := []string{}
	for fileScan.Scan() {
		item := fileScan.Text()
		items := []string{item}
		var j int
		for i := 0; i < len(items); i += 3 {
			j += 3
			group = append(group, items[i])
			if j > len(items) {
				j = len(items)
			}
		}
	}
	fmt.Println(group)

}

func prioritySum(l string) int {
	x := 0
	switch l {
	case "a":
		x = 1
	case "b":
		x = 2
	case "c":
		x = 3
	case "d":
		x = 4
	case "e":
		x = 5
	case "f":
		x = 6
	case "g":
		x = 7
	case "h":
		x = 8
	case "i":
		x = 9
	case "j":
		x = 10
	case "k":
		x = 11
	case "l":
		x = 12
	case "m":
		x = 13
	case "n":
		x = 14
	case "o":
		x = 15
	case "p":
		x = 16
	case "q":
		x = 17
	case "r":
		x = 18
	case "s":
		x = 19
	case "t":
		x = 20
	case "u":
		x = 21
	case "v":
		x = 22
	case "w":
		x = 23
	case "x":
		x = 24
	case "y":
		x = 25
	case "z":
		x = 26
	case "A":
		x = 27
	case "B":
		x = 28
	case "C":
		x = 29
	case "D":
		x = 30
	case "E":
		x = 31
	case "F":
		x = 32
	case "G":
		x = 33
	case "H":
		x = 34
	case "I":
		x = 35
	case "J":
		x = 36
	case "K":
		x = 37
	case "L":
		x = 38
	case "M":
		x = 39
	case "N":
		x = 40
	case "O":
		x = 41
	case "P":
		x = 42
	case "Q":
		x = 43
	case "R":
		x = 44
	case "S":
		x = 45
	case "T":
		x = 46
	case "U":
		x = 47
	case "V":
		x = 48
	case "W":
		x = 49
	case "X":
		x = 50
	case "Y":
		x = 51
	case "Z":
		x = 52
	}
	return x
}

// func prioritySum(string) int{
// 	value1 := [range(abcdefghijklmnopqrstuvwxyz): range(1,26)]
// // I'm dumb...
// 	value2 = dict(zip("ABCDEFGHIJKLMNOPQRSTUVWXYZ", range(27,52)))
// 	sum := value1 + value2
// 	return sum
// }
