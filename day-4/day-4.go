package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
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
		isp1 := intArray(sp1)
		isp2 := intArray(sp2)
		// fmt.Println(sp1, sp2)
		fmt.Println("------------")
		fmt.Println(intArray(sp1), intArray(sp2))
		fmt.Println("------------")
		fmt.Println("Counter 1:", counter(isp1))
		fmt.Println("Counter 2:", counter(isp2))
		fmt.Println(reflect.DeepEqual(counter(isp1), counter(isp2)))
		fmt.Println(checker(isp1, isp2))
		fmt.Println("++++++++++++")
	}

}

// Splits string arrays based on "-"
func splitter(split []string) []string {
	result := []string{}
	for _, r := range split {
		result = strings.Split(r, "-")
	}
	return result
}

// Converts slice of string into slice of int
func intArray(l []string) []int {
	var t = []int{}
	for _, i := range l {
		j, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal()
		}
		t = append(t, j)
	}
	return t
}

// Makes array of sections
func counter(x []int) []int {
	count := []int{}
	i := x[0]
	j := x[1]
	for i <= j {
		count = append(count, i)
		i++
	}
	return count
}

func checker(a, b []int) bool {
	rep := 0
	if len(b) >= len(a) {
		for _, v := range a {
			s := int(v)
			for _, w := range b {
				t := int(w)
				if s == t {
					rep++
				}
			}
		}
	} else if len(a) > len(b) {
		for _, v := range b {
			s := int(v)
			for _, w := range a {
				t := int(w)
				if s == t {
					rep++
				}
			}
		}
	}
	if (len(a) >= len(b) && rep == len(a)) || (len(b) >= len(a) && rep == len(b)) {
		return true
	} else {
		return false
	}
}
