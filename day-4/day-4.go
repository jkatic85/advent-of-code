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
	readFile, err := os.Open("sample-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileScan := bufio.NewScanner(readFile)
	split := []string{}
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
		fmt.Println("++++++++++++")
	}

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

func checker(a, b []int) bool {
	return true
}
