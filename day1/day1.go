package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sortDescending(list []int) []int {
	sort.Slice(list, func(i, j int) bool {
		return list[i] > list[j]
	})
	for _, v := range list {
		fmt.Println(v)
	}
	return list
}

func topThree(list []int) int {
	var sorted = sortDescending(list)
	return sorted[0] + sorted[1] + sorted[2]
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	var weights []int
	count := 0
	for _, eachline := range txtlines {
		fmt.Println(count)
		if eachline == "" {
			weights = append(weights, count)
			count = 0
		} else {
			var num int
			_, err := fmt.Sscan(eachline, &num)
			check(err)
			count += num
		}
	}
	fmt.Println("MAX INDEX:", topThree(weights))
}
