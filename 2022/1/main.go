package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Unable to read file: %v\n", err)
		os.Exit(1)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	maxCalories := 0
	caloriesPerElf := 0
	for fileScanner.Scan() {
		lineEntry := fileScanner.Text()

		if lineEntry == "" {
			// if line entry is blank, check max of caloriesPerElf & maxCalories, and reset caloriesPerElf
			if caloriesPerElf > maxCalories {
				maxCalories = caloriesPerElf
			}
			caloriesPerElf = 0
			continue
		}

		calories, err := strconv.Atoi(lineEntry)
		if err != nil {
			fmt.Printf("Unable to convert string to int: %v\n", err)
			os.Exit(1)
		}

		caloriesPerElf += calories
	}

	fmt.Printf("Max calories carried by an Elf: %d\n", maxCalories)
}

func partTwo() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Unable to read file: %v\n", err)
		os.Exit(1)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	totalCaloriesPerElf := []int{}
	caloriesPerElf := 0
	for fileScanner.Scan() {
		lineEntry := fileScanner.Text()

		if lineEntry == "" {
			totalCaloriesPerElf = append(totalCaloriesPerElf, caloriesPerElf)
			caloriesPerElf = 0
			continue
		}

		calories, err := strconv.Atoi(lineEntry)
		if err != nil {
			fmt.Printf("Unable to convert string to int: %v\n", err)
			os.Exit(1)
		}

		caloriesPerElf += calories
	}

	sort.Ints(totalCaloriesPerElf)

	elves := len(totalCaloriesPerElf)
	totalCaloriesForTopThree := totalCaloriesPerElf[elves-1] + totalCaloriesPerElf[elves-2] + totalCaloriesPerElf[elves-3]
	fmt.Printf("Total calories carried by top 3 Elves: %d\n", totalCaloriesForTopThree)
}
