package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readIntoDict(file *os.File) [][]int {
	reports := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sequence := strings.Fields(scanner.Text())
		nums := []int{}

		for _, elem := range sequence {
			num, pErr := strconv.Atoi(elem)

			if pErr != nil {
				log.Fatal("string is NaN")
			}

			nums = append(nums, num)
		}

		reports = append(reports, nums)
	}

	return reports
}

func isDescendingOrder(arr []int) bool {
	return slices.IsSortedFunc(arr, func(a, b int) int {
		if a < b {
			return 1
		} else if a > b {
			return -1
		} else {
			return 0
		}
	})
}

func checkReport(report []int) bool {
	for index, number := range report {
		if index+1 >= len(report) {
			continue
		}

		nextNum := report[index+1]

		if (nextNum > number && (nextNum-number) > 3) ||
			(nextNum < number && (number-nextNum) > 3) ||
			(nextNum == number) {
			return false
		}
	}

	return true
}

func main() {
	file, err := os.Open("./reports.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dict := readIntoDict(file)
	safeReport := 0

	for _, report := range dict {
		originalReport := slices.Clone(report)
		ascending := slices.IsSorted(report)
		descending := isDescendingOrder(report)

		if !ascending && !descending {
			continue
		}

		unsafe := !checkReport(originalReport)

		if unsafe {
			fmt.Println(fmt.Errorf("report %s is unsafe", fmt.Sprint(originalReport)))
			continue
		}

		safeReport++
	}

	// str, err := fmt.Printf("%s valid reports", string(safeReport))

	// if err != nil {
	// 	log.Fatal("it broke")
	// }

	// fmt.Println(str)
	fmt.Println(safeReport)
}
