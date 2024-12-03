package main

import (
	"bufio"
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

func main() {
	file, err := os.Open("./reports.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dict := readIntoDict(file)

	for _, report := range dict {
		ascending := slices.IsSorted(report)
	}
}
