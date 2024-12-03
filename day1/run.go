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

func sum(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func getLeftAndRight(file *os.File) ([]int, []int) {
	left := []int{}
	right := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())

		if len(nums) < 2 {
			log.Fatal("not enough numbers")
		}

		leftNum, pErr := strconv.Atoi(nums[0])

		if pErr != nil {
			log.Fatal("string is NaN")
		}

		rightNum, pErr := strconv.Atoi(nums[1])

		if pErr != nil {
			log.Fatal("string is NaN")
		}

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	if len(left) != len(right) {
		log.Fatal("lists are not equal in length")
	}

	return left, right
}

func getTotalDistance(left []int, right []int) int {
	slices.Sort(left)
	slices.Sort(right)
	distances := []int{}
	for index, number := range left {
		result := number - right[index]

		if result < 0 {
			result = result * -1
		}

		distances = append(distances, result)
	}

	return sum(distances)
}

func sumEquals(enumerable []int, value int) int {
	occurences := 0

	for _, it := range enumerable {
		if it == value {
			occurences++
		}
	}

	return occurences
}

func getSimilarityScore(left []int, right []int) int {
	similarityScore := 0
	for _, number := range left {
		similarityScore += number * sumEquals(right, number)
	}

	return similarityScore
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	left, right := getLeftAndRight(file)
	// fmt.Println(getTotalDistance(left, right))
	fmt.Println(getSimilarityScore(left, right))
}
