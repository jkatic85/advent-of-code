package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	caloriesSums := []int{}
	caloriesSum := 0

	for fileScanner.Scan() {
		calorie, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			caloriesSums = append(caloriesSums, caloriesSum)
			caloriesSum = 0
		} else {
			caloriesSum += calorie
		}
	}
	caloriesSums = append(caloriesSums, caloriesSum)

	sort.Sort(sort.Reverse(sort.IntSlice(caloriesSums)))
	fmt.Println(caloriesSums)

	sum := 0
	// for i, v := range caloriesSums {
	// 	if i == 3 {
	// 		break
	// 	}
	// 	sum += v
	// }
	for _, v := range caloriesSums[:3] {
		sum += v
	}
	fmt.Println(sum)
}
